package main

import (
	"context"
	"disapp/log"
	"disapp/registry"
	"disapp/service"
	"fmt"
	stlog "log"
)

func main() {
	log.Run("./app.log")

	host, port := "localhost", "4000"

	r := registry.Registration{
		ServiceName: registry.LogService,
		ServiceURL:  fmt.Sprintf("http://%v:%v", host, port),
	}
	shouldRegister := true
	ctx, err := service.Start(context.Background(), r, host, port, log.RegisterHandlers, shouldRegister)
	if err != nil {
		stlog.Fatal(err)
	}
	<-ctx.Done()
	fmt.Println("Shutting down log service")
}
