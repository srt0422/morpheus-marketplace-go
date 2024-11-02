// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package morpheusmarketplace_test

import (
	"context"
	"os"
	"testing"

	"github.com/srt0422/morpheus-marketplace-go"
	"github.com/srt0422/morpheus-marketplace-go/internal/testutil"
	"github.com/srt0422/morpheus-marketplace-go/option"
)

func TestUsage(t *testing.T) {
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
	model, err := client.Blockchain.Models.New(context.TODO(), morpheusmarketplace.BlockchainModelNewParams{
		Fee:     morpheusmarketplace.F("0.01"),
		IpfsID:  morpheusmarketplace.F("QmX..."),
		ModelID: morpheusmarketplace.F("mod-67890"),
		Name:    morpheusmarketplace.F("Image Recognition Model"),
		Stake:   morpheusmarketplace.F("1000"),
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", model.ID)
}
