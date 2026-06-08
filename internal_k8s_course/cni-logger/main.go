// main.go
//
// A minimal CNI plugin that logs all CNI calls to a file.
// It doesn't do real networking — it just records what kubelet sends,
// so you can see the CNI contract in action.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

const logFile = "/var/log/cni-logger.log"

// logToFile writes a timestamped message to our log file
func logToFile(format string, args ...interface{}) {
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer f.Close()
	logger := log.New(f, "", 0)
	timestamp := time.Now().Format("2006-01-02T15:04:05.000Z")
	logger.Printf("[%s] %s", timestamp, fmt.Sprintf(format, args...))
}

func main() {
	command := os.Getenv("CNI_COMMAND")
	containerID := os.Getenv("CNI_CONTAINERID")
	netns := os.Getenv("CNI_NETNS")
	ifName := os.Getenv("CNI_IFNAME")
	cniPath := os.Getenv("CNI_PATH")
	cniArgs := os.Getenv("CNI_ARGS")

	// Read the network config JSON from stdin
	stdinData, _ := os.ReadFile("/dev/stdin")

	// Pretty-print the config for the log
	var prettyConfig json.RawMessage
	json.Unmarshal(stdinData, &prettyConfig)
	configStr, _ := json.MarshalIndent(prettyConfig, "  ", "  ")

	// Log everything we received
	logToFile("========== CNI %s ==========", command)
	logToFile("  CNI_CONTAINERID = %s", containerID)
	logToFile("  CNI_NETNS       = %s", netns)
	logToFile("  CNI_IFNAME      = %s", ifName)
	logToFile("  CNI_PATH        = %s", cniPath)
	logToFile("  CNI_ARGS        = %s", cniArgs)
	logToFile("  stdin config:\n  %s", string(configStr))

	switch command {
	case "ADD":
		// Return a minimal valid CNI result.
		// We use a documentation-range IP (198.51.100.x) since this interface
		// isn't used for real traffic — Multus handles primary networking.
		result := fmt.Sprintf(`{
  "cniVersion": "1.0.0",
  "ips": [{"address": "198.51.100.1/24"}],
  "interfaces": [{"name": "%s", "sandbox": "%s"}]
}`, ifName, netns)
		logToFile("  returning dummy result for ADD")
		logToFile("========== END %s ==========\n", command)
		fmt.Print(result)

	case "DEL":
		logToFile("  nothing to clean up (logging-only plugin)")
		logToFile("========== END %s ==========\n", command)

	case "CHECK":
		logToFile("  CHECK: reporting all OK")
		logToFile("========== END %s ==========\n", command)

	case "VERSION":
		logToFile("========== END %s ==========\n", command)
		fmt.Print(`{"cniVersion":"1.0.0","supportedVersions":["0.3.0","0.3.1","0.4.0","1.0.0"]}`)

	default:
		logToFile("  unknown command: %s", command)
		logToFile("========== END %s ==========\n", command)
		fmt.Fprintf(os.Stderr, "unknown CNI command: %s", command)
		os.Exit(1)
	}
}
