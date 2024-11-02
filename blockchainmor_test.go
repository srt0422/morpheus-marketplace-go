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

func TestBlockchainMorSend(t *testing.T) {
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
	_, err := client.Blockchain.Mor.Send(context.TODO(), morpheusmarketplace.BlockchainMorSendParams{
		Amount: morpheusmarketplace.F("250"),
		To:     morpheusmarketplace.F("to"),
	})
	if err != nil {
		var apierr *morpheusmarketplace.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
