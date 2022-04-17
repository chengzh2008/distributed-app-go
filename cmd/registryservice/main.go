package main

import (
	"context"
	"disapp/registry"
	"disapp/service"
	"fmt"
	stlog "log"
)

func main() {

	host, port := "localhost", "3000"
	r := registry.Registration{
		ServiceName: registry.RegService,
		ServiceURL:  fmt.Sprintf("http://%v:%v", host, port),
	}
	shouldRegister := false
	ctx, err := service.Start(context.Background(), r, host, port, registry.RegisterHandlers, shouldRegister)
	if err != nil {
		stlog.Fatal(err)
	}

	<-ctx.Done()
	fmt.Println("Shutting down registry service")

}
