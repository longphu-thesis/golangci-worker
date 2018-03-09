package analytics

import (
	"context"
	"os"

	"github.com/golangci/golangci-worker/app/utils/runmode"
	"github.com/stvp/rollbar"
)

func trackError(ctx context.Context, err error, level string) {
	if !runmode.IsProduction() {
		panic(err)
	}

	trackingProps := getTrackingProps(ctx)
	f := &rollbar.Field{
		Name: "props",
		Data: trackingProps,
	}

	rollbar.Error(level, err, f)
}

func init() {
	rollbar.Token = os.Getenv("ROLLBAR_API_TOKEN")
	rollbar.Environment = "production" // defaults to "development"
}