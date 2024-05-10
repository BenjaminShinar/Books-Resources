<!--
ignore these words in spell check for this file
// cSpell:ignore Trainium Inferntia parallelizable lemmatization
-->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

# Generative AI Learning Plan for Developers

[Generative AI Learning Plan for Developers](https://explore.skillbuilder.aws/learn/lp/2068/Generative%2520AI%2520Learning%2520Plan%2520for%2520Developers)

> This learning plan is designed to introduce generative AI to software developers interested in leveraging large language models without fine-tuning.\
> The digital training included in this learning plan will provide an overview of generative AI, planning a generative AI project, getting started with Amazon Bedrock, the foundations of prompt engineering, and the architecture patterns to build generative AI applications using Amazon Bedrock and Langchain.

## Introduction to Generative AI - Art of the Possible

<details>
<summary>
Introduction and use-cases.
</summary>

> The Introduction to Generative AI - Art of the Possible course provides an introduction to generative AI, use cases, risks and benefits. With the help of a content generation example, we illustrate the art of the possible.\
> By the end of the course, learners should be able to describe the basics of generative AI, its risks and benefits. They should also be able to articulate how content generation can be used in their business.

### Introduction To Generative AI

<details>
<summary>
What is ML and Generative AI.
</summary>

#### Overview of ML

> Generative artificial intelligence (generative AI) is a branch of machine learning (ML). It is concerned with the development of algorithms that can create natural language text, images, code, audio, or videos based on user input.

we use machine learning on data sets to recognize patterns and then make predictions.

> A dataset is used to train a model. In this dataset, there are features and labels. The goal is to take the features as inputs and find a formula that predicts the labels, or outputs. The resulting ML algorithms can take new data, recognize patterns in the data, apply the formula, and make predictions about the data.

the field exits for a few decades, and many services are using machine learning for more than twenty years. amazon itself uses it for personalized recomendations, amazon prime, Alexa voice assistant and other services.

Generative AI is a subset of deep learning (itself a subset of machine learning), which removes the re-training and fine-tuning steps that usually require labeled data to train new models. it is based on a pre-trained foundational model (FM) which is a large language model (LLM) that was trained on internet-scale datasets.

> The large language models (LLM) have the ability to predict the next word in a sentence by taking into consideration the position and the context of a word in a sentence. LLMs use this ability to generate new content.

#### Basics of Generative AI

> Like all artificial intelligence, generative AI is powered by ML models. However, generative AI is powered by very large models that are pretrained on vast collections of data.

one usage of generative AI is <cloud>AWS CodeWhisperer</cloud>, a code generation tool. it is pitched as a programming assistance tool that can help programmers write better code without leaving the IDE to search for answers in online forums or in the documentation.

other use cases for generative AI are for customer experience (chatbots, personalizations), boosting productivity (smart content search, text summarization, insights from data) and for generating content (video, animations, images, text).

<cloud>AWS Alexa</cloud> uses different models of generative AI to create custom stories based on user input.

#### Generative AI use cases

Aws has Generative AI services:

1. <cloud>Bedrock</cloud>
2. <cloud>SageMaker</cloud>, <cloud>SageMaker JumpStart</cloud>
3. <cloud>Trainium</cloud>, <cloud>Inferntia</cloud>

<cloud>Trainium</cloud>, <cloud>Inferntia</cloud> are specialized chips for machine learning. they were designed to run machine learning training with high performance and low costs.\
<cloud>SageMaker</cloud> service provides the option of training a LLM (using the specialized chips) or use <cloud>SageMaker JumpStart</cloud> to re-train a pre-build model with new data.\
<cloud>AWS Bedrock</cloud> provides foundational models in a fully managed service.

Examples of generative AI use-cases:

1. HealthCare: empowering healthcare software, personalized medicine care, enhancing medical image and diagnosis.
2. Life Science - Molecular structures, predicting protein folding, generating designs.
3. Finance - Fraud detection, portfolio management.
4. Manufacturing - maintenance and workflow optimization.
5. Retail - price optimization, store layout optimization, product review summary.
6. Media and Entrainment - content generation.

</details>

### Importance Of Generative AI

<details>
<summary>
Basic Usecases for generative AI
</summary>

#### Generative AI in Practice

Demo of generative AI use in content summary from a detailed report, content generation, code generation and chatbot for customer service. seeing the prompt and the response.

#### Risks and Benefits

> With the accelerated adoption and increased reach of generative AI, social and legal risks are growing too. You should also consider operational risks because of a single point of failure or inconsistent outputs. You can establish AI principles to prevent harm, audit systems, gain trust, and meet regulatory requirements.
>
> Regulatory requirements – For example, content that potentially violates another individual's intellectual property is a regulatory concern.
>
> Social risks – For example, the possibility of unwanted content that might reflect negatively on your organization is a social risk.
>
> Privacy concerns – For example, the information shared with your model can include personal information and can potentially violate privacy laws.

benefits:

> - Personalize customer interactions.
> - Generate novel content.
> - Efficiently adapt pre-built models to business use cases.
> - Achieve productivity gains through automation.

</details>

</details>

## Planning a Generative AI Project

<details>
<summary>
How to approach creating a Generative AI Project
</summary>

### Technical Foundation and Terminology for Generative AI

<details>
<summary>
A Bit of the basics.
</summary>

#### Generative AI Fundamentals

Foundational Models are pre-built models that were trained on large datasets, and can be adapted to specific uses downstream.

the process begins with the data, which is unlabeled, and is the starting point of the model. the data isn't specific to any domain, so the resulting foundational model is generalized, and can then be adapted to specific tasks.

> The transformer architecture is a type of neural network that is efficient, easy to scale and parallelize, and can model interdependence between input and output data.

transformeres use GPUs to process data at scale. a transformer for text data is aware of the positioning of words in the sentences, and it can use the context to differentiate between ambiguous words.

#### Generative AI Transformer in Practice

the transfomer gets an input, which in our case is a sentence. the first step is tokenization and encoding, which breaks down the sentence into tokens (words, punctuations, phrases). the tokens are then embedded into a three-dimensional space which maps the relation between words. the smaller the distance between them, the more related they are.\
Once all the tokens are encoded, the model can create a response vector and decode it into a textual response.

> To reiterate an important point about transformer models: When compared to their predecessors, like recurrent neural nets, they are more parallelizable.\
> This is because they do not process words sequentially one at a time. Instead, transformer models process the entire input all at once during the learning cycle.\
> Due to this and the thousands of hours engineers spend fine-tuning and training foundation models, they’re able to provide reasonable, or reasonable-sounding, answers to almost any input you provide

#### Generative AI Context

> Context is a one-on-one session with the model. It does not persist when you start a new conversation, and there is an upper limit on the number of tokens that can be remembered in each context. This means that the initial information the model is using can be lost.

this comes up in chatbots, which retain the context for the conversation so they could answer follow-up questions and understand pronouns.

</details>

### Planning The Generative AI Project

<details>
<summary>
Steps in Planning a Generative AI Project.
</summary>

deciding between using (adapting) a pre-build model or whether a model should be fine-tuned.

steps:

1. define the scope
2. select a model
3. adapt a model
4. use the model

the scope of the project is the customers who will use it, the problem they encounter (pain points), and what they wish to solve. we also look at our organization and identify if we can provide the solution based on resources, effort, and challenges from regulation and governance policies. then we consider the effects of the solution on the customers, the organization and how will it affect the market as a whole.\
The impact of an AI solution can be short-term and long-term, some solutions require more time than others.

a pre-trained model is a good option for quick solutions that don't require customization. in contrast, when we fine-tune an existing model we can get more specialized results and we have wider flexability. this comes with additional expenses of computational power, time and technical expertise.

> - **Prompt engineering** is the process of designing and refining your prompts or inputs in order for the model to generate specific types of outputs that suit your needs. By making a few small changes to the language you use as the input, you can drastically change the quality of the output.
> - **Fine-tuning** is a continuation of pre-training that creates a new specialized model and requires high-quality, labeled data. When fine-tuning, you change the parameters in the model and create a new model specific to your solution.

as with every project, even when it's done, we still have to monitor it. we need to ask ourselves:

> - Have you managed all of the responsible AI concerns?
> - Do you have a plan for feedback from users?
> - How are you going to track performance of your FM over time?
> - How are you tracking changes to the pre-trained model so you can re-train your fine-tuned model?

</details>

### Evaluating The Use Of Generative Ai For Your Project

<details>
<summary>
Risks and Mitigation for Generative AI.
</summary>

> Now that you have learned about the benefits of generative AI, consider some of the risks and actions from a technology standpoint to help mitigate them.

Fairness - does our model include a bias? LLM can pick up on markers which are related to specific groups and use them to make predictions, so we might need to counter-act that and ensure fairness.

Privacy - does the training data include private information which might later leak as a response? this also include copyright and intellectual property concerns.

when we use AI to generate content, it might generate content which we can't accepts, such as harmful, offensive and inappropriate phrases. the generated data can also be wrong (**hallucinations**), as it is just a predictive response, and isn't not necessarily be grounded in reality. we can't trust the AI completely, and it must be verified with other sources.

</details>

</details>

## Amazon Bedrock Getting Started

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>

</details>

## Foundations of Prompt Engineering

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>

Prompts are input to a foundational model, this is what the model takes and how it chooses the response. if we modify the prompt, we will get a different response.\
Prompt engineering is how we create those prompts.

### Basics of Foundation Models

foundation models are larger than traditional ML models, and are also suited to a wider variety of tasks. they use neural networks to handle complex tasks and can do various things. Also, unlike ML models which use supervised (or semi-supervised, or unsupervised) learning, Founation models rely on self-supervised training algorithms.

the first phase of a model is the pre-training, which uses large datasets. it can either be entirely self-supervised or use reinforcement learning from human feedback (RLHF). the next phase is fine-tunning the model, which is either done with human input, or by supplying the model with domain specific, specialized datasets. finally, we interact with the model via _prompts_.

Text-To-Text models can use NLP (natural language processing) as text processing technique:

- tokenization
- stemming
- lemmatization
- stop word removal
- part-of-speech tagging
- named entity recognition
- speech recognition
- sentiment analysis

Recurrent neural network (RNN) user sequential data, and are good for some tasks such as speech recognition or machine translation. but are rather costly and slow, and they don't scale well. in contrast, _Transformers_ are very suited to parallel processing, and are much faster to train. transformer are able to encode the input data to the model, and then decode the response out.

Text-To-Image models use diffusion architecture: which is a two-step process of gradually adding noise to an image until only the noise remains (forward diffusion) and the model learns to predict this noise. at the second step, a noisy image is "de-noised".

Large Language models (LLM) are a subset of foundation models. they are trained on large datasets and employ the transformer model. they rely on three layers.

> 1. embedding layer - The embedding layer converts input text to vector representations called embeddings. This layer can capture complex relationships between the embeddings, so the model can understand the context of the input text.
> 2. feedforward layer - The feedforward layer consists of several connected layers that transform the embeddings into more weighted versions of themselves. Essentially, this layer continues to contextualize the language and helps the model better understand the input text's intent.
> 3. attention mechanism - With the attention mechanism, the model can focus on the most relevant parts of the input text. This mechanism, a central part of the transformer model, helps the model achieve the most accurate output results.

### Fundamentals of Prompt Engineering

> Prompt engineering is an emerging field that focuses on developing, designing, and optimizing prompts to enhance the output of LLMs for your needs. It gives you a way to guide the model's behavior to the outcomes you want to achieve.

modifying the prompt is a way to change the model output without the expensive steps of fine-tunning and adding more data. prompts consist of:
bullet

> - Instructions: This is a task for the large language model to do. It provides a task description or instruction for how the model should perform.
> - Context: This is external information to guide the model.
> - Input data: This is the input for which you want a response.
> - Output indicator: This is the output type or format.

when creating a prompt, we should follow the best practices.

- be clear and concise - use natural langague, avoid isolated keywords.
- include context (if needed) - enhance the input data with the relevant context.
- use directives for the appropriate response type - specify how you would like the response to be formatted.
- consider the output in the prompt - mention the output at the end of the prompt.
- start prompts with an interrogation - phrase the input as a question.
- provide an example response
- break up complex tasks - either in the same prompt, or across several ones.
- experiment and be creative

### Basic Prompt Techniques

> **Zero-shot prompting** - is a prompting technique where a user presents a task to an LLM without giving the model further examples. Here, the user expects the model to perform the task without a prior understanding, or shot, of the task. Modern LLMs demonstrate remarkable zero-shot performance.

for zero-shot prompting, larger LLMs usually have better results. instruction tunning can greatly increase the quality.

> **Few-shot prompting** is a prompting technique where you give the model contextual information about the requested tasks. In this technique, you provide examples of both the task and the output you want. Providing this context, or a few shots, in the prompt conditions the model to follow the task guidance closely.

(this is basically providing a larger prompt with examples)

> **Chain-of-thought** (CoT) prompting breaks down complex reasoning tasks through intermediary reasoning steps. You can use both zero-shot and few-shot prompting techniques with CoT prompts.\
> Chain-of-thought prompts are specific to a problem type. You can use the phrase "Think step by step" to invoke CoT reasoning from your machine learning model.

(asking the model to work in steps)

### Advanced Prompt Techniques

> **Self-consistency** is a prompting technique that is similar to chain-of-thought prompting. However, instead of taking the obvious step-by-step, or greedy path, self-consistency prompts the model to sample a variety of reasoning paths. Then, the model aggregates the final answer based on multiple data points from the various paths.

(providing example of how to do the correct analysis)

> **Tree of thoughts** (ToT) is another technique that builds on the CoT prompting technique. CoT prompting samples thoughts sequentially, but ToT prompting follows a tree-branching technique. With the ToT technique, the LLM can learn in a nuanced way, considering multiple paths instead of one sequential path.
>
> **Retrieval Augmented Generation** (RAG) is a prompting technique that supplies domain-relevant data as context to produce responses based on that data and the prompt. This technique is similar to fine-tuning. However, rather than having to fine-tune an FM with a small set of labeled examples, RAG retrieves a small set of relevant documents from a large corpus and uses that to provide context to answer the question.\
> RAG will not change the weights of the foundation model whereas fine-tuning will change model weights.
>
> **Automatic Reasoning and Tool-use** (ART) - ART is a prompting technique that builds on the chain-of-thought process.

### Model-Specific Prompt Techniques

> - **Amazon Titan FMs** – Amazon Titan Foundation Models (FMs) are pretrained on large datasets, making them powerful, general-purpose models. Use them as is or customize them with your own data for a particular task without annotating large volumes of data.
> - **Anthropic Claude** – Claude is an AI chatbot built by Anthropic, which you can access through chat or API in a developer console. Claude can process conversation, text, summarization, search, creative writing, coding, question answering, and more. Claude is designed to respond conversationally and can modify character, style, and conduct to best suit output needs.
> - **AI21 Jurassic-2** – Jurassic-2 is trained specifically to process instructions-only prompts with no examples, or zero-shot prompts. Using only instructions in the prompt can be the most natural way to interact with large language models.

these three models are part of the <cloud>AWS Bedrock</cloud> service.

#### Prompets Parameters

> When interacting with LLMs through API or directly, you can configure prompt _parameters_ to get customized results. Generally, you should only adjust one parameter at a time, and results can vary depending on the LLM.
>
> **Determinism parameters** - Choosing lower values for each parameter provides factual results, and choosing higher values provides diverse and creative results. The following parameters control determinism:
>
> - _Temperature_ controls randomness. Lower values focus on probable tokens, and higher values add randomness and diversity. Use lower values for factual responses and higher values for creative responses.
> - _Top_p_ adjusts determinism with "nucleus sampling." Lower values give exact answers, while higher values give diverse responses. This value controls the diversity of the model's responses.
> - _Top_k_ is the number of the highest-probability vocabulary tokens to keep for top- k-filtering. Similar to the Top_p parameter, Top_k defines the cutoff where the model no longer selects the words.
>
> Token count parameters include the following:
>
> - _MinTokens_ is the minimum number of tokens to generate for each response.
> - _MaxTokenCount_ is the maximum number of tokens to generate before stopping.
>
> _StopSequences_ is a list of strings that will cause the model to stop generating.
>
> _numResults_ is the number of responses to generate for a given prompt.
>
> These penalties are only available in Jurassic. Penalties parameters include the following:
>
> - _FrequencyPenalty_ is a penalty applied to tokens that are frequently generated.
> - _PresencePenalty_ is a penalty applied to tokens that are already present in the prompt.
> - _CountPenalty_ is a penalty applied to tokens based on their frequency in the generated responses.

#### Amazon Titan Large Prompt Guidance

- specify output length
- provide simple, clear and complete instructions
- provide default output when necessary
- use separator characters for API calls
- personalize responses

### Anthropic Claude Prompt Guidance

- add tags in the prompts - "Human" and "Assistant" tags, since the model was fine-tuned using human feedback.
- include detailed description
- limit the response by pre-filling
- use XML tags
- specify output length
- set clear expectations
- break up complex tasks

#### AI21 Labs Jurassic-2 Prompt Guidance

- Specify output length
- avoid ambiguity
- include additional context or instructions - use the term _instruction_ in the prompt.
- avoid negative formulations
- switch the order of instructions for long documents

### Addressing Prompt Misuses

### Mitigating Bias

</details>

## TakeAways

<!-- <details> -->
<summary>
Stuff worth remembering
</summary>

- LLM - Large Language Model
- FM - Foundational Model
- RAG - Retrieval Augmented Generation
- RLHF - reinforcement learning from human feedback
- NLP - Natural language processing
- RNN - Recurrent neural network

</details>
