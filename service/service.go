package service

import (
	"context"
	"disapp/registry"
	"fmt"
	"log"
	"net/http"
)

func Start(ctx context.Context, registration registry.Registration, host, port string, registerHandlersFunc func(), shouldRegister bool) (context.Context, error) {
	registerHandlersFunc()
	ctx = startService(ctx, registration.ServiceName, host, port, shouldRegister)
	if shouldRegister {
		err := registry.RegisterService(registration)
		if err != nil {
			return ctx, err
		}
	}

	return ctx, nil
}

func startService(ctx context.Context, serviceName registry.ServiceName, host, port string, shouldDeregister bool) context.Context {
	ctx, cancel := context.WithCancel(ctx)

	var srv http.Server
	srv.Addr = ":" + port

	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		fmt.Printf("%v started. Press any key to stop.\n", serviceName)
		var s string
		fmt.Scanln(&s)
		if shouldDeregister {
			err := registry.DerigesterService(fmt.Sprintf("http://%v:%v", host, port))
			if err != nil {
				log.Println(err)
			}
		}
		srv.Shutdown(ctx)
		cancel()
	}()

	return ctx
}
