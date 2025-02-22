package openfeature_test

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/open-feature/go-sdk/pkg/openfeature"
)

func ExampleNewClient() {
	client := openfeature.NewClient("example-client")
	fmt.Printf("Client Name: %s", client.Metadata().Name())
	// Output: Client Name: example-client
}

func ExampleClient_BooleanValue() {
	client := openfeature.NewClient("example-client")
	value, err := client.BooleanValue(
		context.Background(), "test-flag", true, openfeature.EvaluationContext{},
	)
	if err != nil {
		log.Fatal("error while getting boolean value : ", err)
	}

	fmt.Printf("test-flag value: %v", value)
	// Output: test-flag value: true
}

func ExampleClient_StringValue() {
	client := openfeature.NewClient("example-client")
	value, err := client.StringValue(
		context.Background(), "test-flag", "openfeature", openfeature.EvaluationContext{},
	)
	if err != nil {
		log.Fatal("error while getting string value : ", err)
	}

	fmt.Printf("test-flag value: %v", value)
	// Output: test-flag value: openfeature
}

func ExampleClient_FloatValue() {
	client := openfeature.NewClient("example-client")
	value, err := client.FloatValue(
		context.Background(), "test-flag", 0.55, openfeature.EvaluationContext{},
	)
	if err != nil {
		log.Fatalf("error while getting float value: %v", err)
	}

	fmt.Printf("test-flag value: %v", value)
	// Output: test-flag value: 0.55
}

func ExampleClient_IntValue() {
	client := openfeature.NewClient("example-client")
	value, err := client.IntValue(
		context.Background(), "test-flag", 3, openfeature.EvaluationContext{},
	)
	if err != nil {
		log.Fatalf("error while getting int value: %v", err)
	}

	fmt.Printf("test-flag value: %v", value)
	// Output: test-flag value: 3
}

func ExampleClient_ObjectValue() {
	client := openfeature.NewClient("example-client")
	value, err := client.ObjectValue(
		context.Background(), "test-flag", map[string]string{"foo": "bar"}, openfeature.EvaluationContext{},
	)
	if err != nil {
		log.Fatal("error while getting object value : ", err)
	}

	str, _ := json.Marshal(value)
	fmt.Printf("test-flag value: %v", string(str))
	// Output: test-flag value: {"foo":"bar"}
}
