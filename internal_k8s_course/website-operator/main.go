// main.go
// A minimal Kubernetes operator that watches "Website" custom resources
// and reconciles them by creating/updating Deployments and Services.
//
// This demonstrates the core operator pattern:
// 1. Watch for changes to a custom resource (Website)
// 2. On each change, reconcile: ensure the real state matches desired state
// 3. Use OwnerReferences so child resources get garbage-collected on delete
package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/dynamicinformer"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/utils/ptr"
)

// WebsiteSpec matches our CRD's spec fields
type WebsiteSpec struct {
	Image    string `json:"image"`
	Replicas int32  `json:"replicas"`
}

// The GroupVersionResource tells the dynamic client which API to watch
var websiteGVR = schema.GroupVersionResource{
	Group:    "training.example.com",
	Version:  "v1",
	Resource: "websites",
}

func main() {
	// Try in-cluster config first (when running as a Pod),
	// fall back to ~/.kube/config (when running locally)
	config, err := rest.InClusterConfig()
	if err != nil {
		config, err = clientcmd.BuildConfigFromFlags("",
			clientcmd.RecommendedHomeFile)
		if err != nil {
			log.Fatalf("Error building kubeconfig: %v", err)
		}
	}

	// Typed client - for creating Deployments, Services, etc.
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating clientset: %v", err)
	}

	// Dynamic client - for watching our custom Website resources
	dynClient, err := dynamic.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating dynamic client: %v", err)
	}

	// Create an informer that watches Website resources and caches them.
	// The 30s resync period means it will re-list all resources every 30s
	// to catch any missed events.
	factory := dynamicinformer.NewDynamicSharedInformerFactory(
		dynClient, 30*time.Second)

	informer := factory.ForResource(websiteGVR).Informer()

	// Register event handlers - these are called when a Website is
	// created, updated, or deleted
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			u := obj.(*unstructured.Unstructured)
			log.Printf("Website ADDED: %s/%s", u.GetNamespace(), u.GetName())
			reconcile(clientset, dynClient, u)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			u := newObj.(*unstructured.Unstructured)
			log.Printf("Website UPDATED: %s/%s", u.GetNamespace(), u.GetName())
			reconcile(clientset, dynClient, u)
		},
		DeleteFunc: func(obj interface{}) {
			u := obj.(*unstructured.Unstructured)
			log.Printf("Website DELETED: %s/%s", u.GetNamespace(), u.GetName())
			// No manual cleanup needed - OwnerReferences handle garbage collection
		},
	})

	ctx := context.Background()
	factory.Start(ctx.Done())
	factory.WaitForCacheSync(ctx.Done())

	log.Println("Website operator started. Watching for Website resources...")
	<-ctx.Done()
}

// reconcile is the heart of the operator. It ensures a Deployment + Service
// exist for the given Website resource, creating or updating as needed.
func reconcile(
	clientset kubernetes.Interface,
	dynClient dynamic.Interface,
	website *unstructured.Unstructured,
) {
	name := website.GetName()
	namespace := website.GetNamespace()

	// Parse the spec from the unstructured object
	specRaw, found, _ := unstructured.NestedMap(website.Object, "spec")
	if !found {
		log.Printf("No spec found for %s", name)
		return
	}

	specJSON, _ := json.Marshal(specRaw)
	var spec WebsiteSpec
	json.Unmarshal(specJSON, &spec)

	log.Printf("Reconciling Website %s: image=%s replicas=%d",
		name, spec.Image, spec.Replicas)

	// --- Ensure Deployment exists ---
	deploymentsClient := clientset.AppsV1().Deployments(namespace)

	deploy := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			// OwnerReferences make K8s automatically delete this Deployment
			// when the parent Website resource is deleted
			OwnerReferences: []metav1.OwnerReference{{
				APIVersion: "training.example.com/v1",
				Kind:       "Website",
				Name:       website.GetName(),
				UID:        website.GetUID(),
				Controller: ptr.To(true),
			}},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: ptr.To(spec.Replicas),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{"website": name},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{"website": name},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Name:  "web",
						Image: spec.Image,
						Ports: []corev1.ContainerPort{{
							ContainerPort: 80,
						}},
					}},
				},
			},
		},
	}

	ctx := context.Background()
	existing, err := deploymentsClient.Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		// Deployment doesn't exist - create it
		_, err = deploymentsClient.Create(ctx, deploy, metav1.CreateOptions{})
		if err != nil {
			log.Printf("Failed to create Deployment: %v", err)
			return
		}
		log.Printf("Created Deployment %s", name)
	} else {
		// Deployment exists - update it to match desired state
		existing.Spec.Replicas = ptr.To(spec.Replicas)
		existing.Spec.Template.Spec.Containers[0].Image = spec.Image
		_, err = deploymentsClient.Update(ctx, existing, metav1.UpdateOptions{})
		if err != nil {
			log.Printf("Failed to update Deployment: %v", err)
			return
		}
		log.Printf("Updated Deployment %s", name)
	}

	// --- Ensure Service exists ---
	servicesClient := clientset.CoreV1().Services(namespace)
	svc := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			// Same OwnerReferences as the Deployment - ensures this Service
			// is also garbage collected when the Website is deleted
			OwnerReferences: []metav1.OwnerReference{{
				APIVersion: "training.example.com/v1",
				Kind:       "Website",
				Name:       website.GetName(),
				UID:        website.GetUID(),
				Controller: ptr.To(true),
			}},
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{"website": name},
			Ports: []corev1.ServicePort{{
				Port:       80,
				TargetPort: intstr.FromInt(80),
			}},
		},
	}

	_, err = servicesClient.Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		_, err = servicesClient.Create(ctx, svc, metav1.CreateOptions{})
		if err != nil {
			log.Printf("Failed to create Service: %v", err)
			return
		}
		log.Printf("Created Service %s", name)
	}

	// --- Update status subresource ---
	status := map[string]interface{}{
		"availableReplicas": int64(spec.Replicas),
		"ready":             true,
	}
	unstructured.SetNestedField(website.Object, status, "status")
	_, err = dynClient.Resource(websiteGVR).Namespace(namespace).
		UpdateStatus(ctx, website, metav1.UpdateOptions{})
	if err != nil {
		log.Printf("Failed to update status: %v", err)
	}
}
