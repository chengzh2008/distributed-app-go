package main

import (
	"context"
	"disapp/registry"
	"disapp/service"
	"fmt"
	stlog "log"
)

func main() {

	ctx, err := service.Start(context.Background(), registry.RegService, "localhost", "4001", registry.RegisterHandlers)
	if err != nil {
		stlog.Fatal(err)
	}

	<-ctx.Done()
	fmt.Println("Shutting down registry service")

}
