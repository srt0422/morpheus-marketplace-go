// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package morpheusmarketplace_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/srt0422/morpheus-marketplace-go"
	"github.com/srt0422/morpheus-marketplace-go/internal/testutil"
	"github.com/srt0422/morpheus-marketplace-go/option"
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
	_, err := client.BlockchainSessions.New(context.TODO(), morpheusmarketplace.BlockchainSessionNewParams{
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
	_, err := client.BlockchainSessions.Get(context.TODO(), "id")
	if err != nil {
		var apierr *morpheusmarketplace.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBlockchainSessionBudget(t *testing.T) {
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
	_, err := client.BlockchainSessions.Budget(context.TODO())
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
	err := client.BlockchainSessions.Close(context.TODO(), "id")
	if err != nil {
		var apierr *morpheusmarketplace.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBlockchainSessionProviderWithOptionalParams(t *testing.T) {
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
	_, err := client.BlockchainSessions.Provider(context.TODO(), morpheusmarketplace.BlockchainSessionProviderParams{
		Provider: morpheusmarketplace.F("4592d8f8d7b001e72cb26a73e4fa1806a51ac79d"),
		Limit:    morpheusmarketplace.F(int64(10)),
		Offset:   morpheusmarketplace.F(int64(0)),
	})
	if err != nil {
		var apierr *morpheusmarketplace.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBlockchainSessionUserWithOptionalParams(t *testing.T) {
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
	_, err := client.BlockchainSessions.User(context.TODO(), morpheusmarketplace.BlockchainSessionUserParams{
		User:   morpheusmarketplace.F("4592d8f8d7b001e72cb26a73e4fa1806a51ac79d"),
		Limit:  morpheusmarketplace.F(int64(10)),
		Offset: morpheusmarketplace.F(int64(0)),
	})
	if err != nil {
		var apierr *morpheusmarketplace.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
