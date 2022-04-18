package main

import (
	"context"
	"disapp/grades"
	"disapp/registry"
	"disapp/service"
	"fmt"
	stlog "log"
)

func main() {

	host, port := "localhost", "5000"

	r := registry.Registration{
		ServiceName: registry.GradingService,
		ServiceURL:  fmt.Sprintf("http://%v:%v", host, port),
	}
	shouldRegister := true
	ctx, err := service.Start(context.Background(), r, host, port, grades.RegisterHandlers, shouldRegister)
	if err != nil {
		stlog.Fatal(err)
	}
	<-ctx.Done()
	fmt.Println("Shutting down grading service")
}
