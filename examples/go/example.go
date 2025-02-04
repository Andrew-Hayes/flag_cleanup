package main

import (
	"log"
	"os"

	harness "github.com/harness/ff-golang-server-sdk/client"
	"github.com/harness/ff-golang-server-sdk/evaluation"
)

var (
	sdkKey string = getEnvOrDefault("FF_API_KEY", "changeme")
	client *harness.CfClient
)

func main() {
	log.Println("Harness SDK Getting Started")

	// Create a feature flag client
	var err error
	client, err = harness.NewCfClient(sdkKey)
	if err != nil {
		log.Fatalf("could not connect to CF servers %s\n", err)
	}
	defer func() { client.Close() }()

	if isEnabled("STALE_FLAG") {
		log.Println("Run true code path")
	} else {
		log.Println("Run false code path")
	}

	if isEnabled("OTHER_FLAG") {
		log.Println("Run true code path")
	} else {
		log.Println("Run false code path")
	}

	if isEnabled("STALE_FLAG_2") {
		log.Println("Run true code path")
	} else {
		log.Println("Run false code path")
	}
}

func doSomething() {
	if isEnabled("STALE_FLAG") {
		log.Println("Run true code path")
	} else {
		log.Println("Run false code path")
	}
}

func isEnabled(flag string) bool {
	// Create a target (different targets can get different results based on rules)
	target := evaluation.Target{
		Identifier: "HT_1",
		Name:       "Harness_Target_1",
		Attributes: &map[string]interface{}{"email": "demo@harness.io"},
	}

	resultBool, err := client.BoolVariation(flag, &target, false)
	if err != nil {
		log.Printf("issue evaluating flag %s, serving default value of false. err: %s", flag, err)
	}
	return resultBool

}

func getEnvOrDefault(key, defaultStr string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultStr
	}
	return value
}
