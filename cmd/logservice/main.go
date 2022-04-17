package main

import (
	"context"
	"disapp/log"
	"disapp/service"
	"fmt"
	stlog "log"
)

func main() {
	log.Run("./app.log")

	host, port := "localhost", "4000"

	ctx, err := service.Start(context.Background(), "Log Servcie", host, port, log.RegisterHandlers)
	if err != nil {
		stlog.Fatal(err)
	}
	<-ctx.Done()
	fmt.Println("Shutting down log service")
}
