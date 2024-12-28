package api_design

import (
	"context"
	"testing"
)

func TestUserConfig(t *testing.T) {
	ctx := context.TODO()
	InitialiseConfigFromEnv(ctx, func(config *Options) {
		config.Name = "test"
		config.Region = "us-west-2"
	})
}
