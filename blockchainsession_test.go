// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package morpheusmarketplace_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/morpheus-marketplace-go"
	"github.com/stainless-sdks/morpheus-marketplace-go/internal/testutil"
	"github.com/stainless-sdks/morpheus-marketplace-go/option"
)

func TestBlockchainSessionNew(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := morpheusmarketplace.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Blockchain.Sessions.New(context.TODO(), morpheusmarketplace.BlockchainSessionNewParams{
		Approval:    morpheusmarketplace.F("approval_abc123"),
		ApprovalSig: morpheusmarketplace.F("signature_xyz789"),
		Stake:       morpheusmarketplace.F("500"),
	})
	if err != nil {
		var apierr *morpheusmarketplace.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBlockchainSessionGet(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := morpheusmarketplace.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Blockchain.Sessions.Get(context.TODO(), "id")
	if err != nil {
		var apierr *morpheusmarketplace.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBlockchainSessionClose(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := morpheusmarketplace.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Blockchain.Sessions.Close(context.TODO(), "id")
	if err != nil {
		var apierr *morpheusmarketplace.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
