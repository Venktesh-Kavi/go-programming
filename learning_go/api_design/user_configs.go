package api_design

import (
	"context"
	"fmt"
	"os"
)

type Options struct {
	Name          string
	Region        string
	internalField string
}

type Config interface{}

type configs []Config

// fn used to customise the configuration on demand apart from deriving from the context.
func InitialiseConfigFromEnv(ctx context.Context, fn func(config *Options)) {
	var o Options
	fn(&o) // pass it as a pointer to be modified by the function.

	resolveInternalField(&o)
	cc := configs{o}
	fmt.Println(cc)
}

func resolveInternalField(o *Options) {
	o.internalField = os.Getenv("region")
}
