package main

import (
	"context"
	"fmt"
	stlog "log"

	"github.com/MarouaneBouaricha/grades-portal/app/log"
	"github.com/MarouaneBouaricha/grades-portal/app/registry"
	"github.com/MarouaneBouaricha/grades-portal/app/service"
)

func main() {
	host, port := "localhost", "6000"
	servicesAddress := fmt.Sprintf("http://%v:%v", host, port)

	var r registry.Registration
	r.ServiceName = registry.LogService
	r.ServiceURL = servicesAddress

	ctx, err := service.Start(
		context.Background(),
		host,
		port,
		r,
		log.RegisterHandlers,
	)
	if err != nil {
		stlog.Fatal(err)
	}

	<-ctx.Done()
	fmt.Println("Shutting down grading service")

}
