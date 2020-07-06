package main

import (
	"account/handler"
	"account/proto"
	tracer "account/trace"
	"github.com/micro/go-micro/v2"
	_ "github.com/micro/go-plugins/registry/kubernetes/v2"
	openTrace "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	"log"
)

func main() {
	serviceName := "micro.service.account"
	t, io, err := tracer.NewTracer(serviceName, "127.0.0.1:6831")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)
	service := micro.NewService(
		micro.Name(serviceName),
		micro.WrapHandler(openTrace.NewHandlerWrapper(opentracing.GlobalTracer())),
	)
	service.Init()
	if err := proto.RegisterAccountServiceHandler(service.Server(), new(handler.AccountService)); err != nil {
		log.Print(err.Error())
		return
	}
	if err := service.Run(); err != nil {
		log.Print(err.Error())
		return
	}
}
