package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zawachte/caas-rest-api/pkg/caas"
)

func main() {
	r := gin.Default()

	srv, err := caas.NewCaasServer()
	if err != nil {
		panic(err)
	}
	//	caas.RegisterHandlers(r, srv)
	caas.RegisterHandlersWithOptions(r, srv, caas.GinServerOptions{BaseURL: "v1"})
	r.Run()
}
