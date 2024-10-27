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

func TestBlockchainProviderBidListWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := morpheusmarketplace.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Blockchain.Providers.Bids.List(
		context.TODO(),
		"id",
		morpheusmarketplace.BlockchainProviderBidListParams{
			Limit:  morpheusmarketplace.F(int64(0)),
			Offset: morpheusmarketplace.F("offset"),
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
