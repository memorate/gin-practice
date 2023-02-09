// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"gin-practice/apps/hello/access"
	"gin-practice/libs/router"
	"github.com/gin-gonic/gin"
)

// Injectors from wire.go:

func InitializeApplication() (application, error) {
	helloApp := access.NewHelloApp()
	http_serverApplication := application{
		helloApp: helloApp,
	}
	return http_serverApplication, nil
}

// wire.go:

type application struct {
	helloApp *access.HelloApp
}

func (app *application) Register(engine *gin.Engine) error {
	router.RegisterSchema("rest", app.helloApp)
	return router.RegisterUrlPatterns(engine)
}