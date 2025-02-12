package server_test

import (
	"context"
	"fmt"
	"testing"

	"grpc_client_streaming_sample/server"

	"github.com/k1LoW/runn"
)

func Test_scenario(t *testing.T) {
	t.Run("", func(t *testing.T) {
		ctx := context.Background()
		opts := []runn.Option{
			runn.T(t),
			runn.Scopes(runn.ScopeAllowReadParent),
		}
		s, err := server.New()
		if err != nil {
			t.Fatal(err)
		}
		if err := s.Serve(ctx); err != nil {
			t.Fatal(err)
		}
		o, err := runn.Load("../testdata/scenarios/**/*.yml", opts...)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("data:%+v\n", o)
		if err := o.RunN(ctx); err != nil {
			t.Fatal(err)
		}
	})
}
