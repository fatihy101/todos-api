package main

import (
	"fmt"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

func TestProvider(t *testing.T) {
	go startServer(true)
	port := 4000
	broker := "https://fatihy.pactflow.io"
	brokerToken := "HW1NX8ihy5dv5XUQzt3gpA"
	pactUrl := "https://fatihy.pactflow.io/pacts/provider/Todo%20API/consumer/Todo%20Client/latest/master"
	// pactUrl := "./todo-client-todo-service.json"
	// Create Pact connecting to local Daemon
	pact := &dsl.Pact{
		Provider:                 "Todo API",
		DisableToolValidityCheck: true,
		LogLevel:                 "DEBUG",
	}

	// Verify the Provider using the locally saved Pact Files
	_, err := pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL:            fmt.Sprintf("http://localhost:%d", port),
		BrokerURL:                  broker,
		BrokerToken:                brokerToken,
		PactURLs:                   []string{pactUrl},
		PublishVerificationResults: true,
		ProviderVersion:            "2.0.0",
	})

	if err != nil {
		t.Fatal(err)
	}

}
