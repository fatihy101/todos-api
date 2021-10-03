package main

import (
	"fmt"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

func TestProvider(t *testing.T) {
	port := 4000
	broker := "https://fatihy.pactflow.io"
	brokerToken := "HW1NX8ihy5dv5XUQzt3gpA"
	pactUrl := "https://fatihy.pactflow.io/pacts/provider/Todo%20API/consumer/todo-client/latest/master"
	// Create Pact connecting to local Daemon
	pact := &dsl.Pact{
		Provider:                 "Todo API",
		LogDir:                   "../logs",
		DisableToolValidityCheck: true,
		LogLevel:                 "DEBUG",
	}

	// Verify the Provider using the locally saved Pact Files
	_, err := pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL:            fmt.Sprintf("http://localhost:%d", port),
		PactURLs:                   []string{pactUrl},
		BrokerURL:                  broker,
		BrokerToken:                brokerToken,
		PublishVerificationResults: true,
		ProviderVersion:            "0.1.0",
		StateHandlers: types.StateHandlers{
			"get all todos": func() error {
				return nil
			},
		},
	})

	if err != nil {
		t.Fatal(err)
	}

}
