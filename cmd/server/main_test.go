package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/k1LoW/runn"
)

func Test_main(t *testing.T) {
	go func() {
		main()
	}()
	t.Run("", func(t *testing.T) {
		ctx := context.Background()
		opts := []runn.Option{
			runn.T(t),
			runn.Scopes(runn.ScopeAllowReadParent),
		}
		o, err := runn.Load("testdata/scenarios/**/*.yml", opts...)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("data:%+v\n", o)
		if err := o.RunN(ctx); err != nil {
			t.Fatal(err)
		}
	})
}
