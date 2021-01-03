package functions

import (
	"context"
	"log"

	"github.com/vmware/govmomi/vim25"
)

func Run(f func(context.Context, *vim25.Client) error) {

	var err error
	ctx := context.Background()
	c, err := NewClient(ctx)
	if err == nil {
		err = f(ctx, c)
	}
	if err != nil {
		log.Fatal(err)
	}
}
