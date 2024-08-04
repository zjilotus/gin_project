package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	//保证下面的优雅启停  放在一个携程下
	go func() {
		log.Printf("web server running in %s\n", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
}
