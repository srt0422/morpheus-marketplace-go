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

func TestProxySessionInitiate(t *testing.T) {
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
	_, err := client.Proxy.Sessions.Initiate(context.TODO(), morpheusmarketplace.ProxySessionInitiateParams{
		ModelID:         morpheusmarketplace.F("model_12345"),
		SessionDuration: morpheusmarketplace.F("3600"),
	})
	if err != nil {
		var apierr *morpheusmarketplace.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProxySessionProviderClaim(t *testing.T) {
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
	_, err := client.Proxy.Sessions.ProviderClaim(
		context.TODO(),
		"id",
		morpheusmarketplace.ProxySessionProviderClaimParams{
			Claim: morpheusmarketplace.F("claim_abc123"),
		},
	)
	if err != nil {
		var apierr *morpheusmarketplace.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
