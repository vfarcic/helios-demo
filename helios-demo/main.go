package main

import (
	"fmt"
	"log"
	"os"

	gin "github.com/helios/go-sdk/proxy-libs/heliosgin"
	http "github.com/helios/go-sdk/proxy-libs/helioshttp"
	"github.com/helios/go-sdk/sdk"
)

func main() {
	sdk.Initialize("helios-demo", "65e7de34aaee5efb5be8", sdk.WithEnvironment("prod"))
	log.SetOutput(os.Stderr)

	// Server
	log.Println("Starting server...")
	router := gin.New()
	router.POST("/video", videoPostHandler)
	router.GET("/videos", videosGetHandler)
	router.GET("/ping", pingHandler)
	router.GET("/", rootHandler)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	router.Run(fmt.Sprintf(":%s", port))
}

func httpErrorBadRequest(err error, ctx *gin.Context) {
	httpError(err, ctx, http.StatusBadRequest)
}

func httpErrorInternalServerError(err error, ctx *gin.Context) {
	httpError(err, ctx, http.StatusInternalServerError)
}

func httpError(err error, ctx *gin.Context, status int) {
	log.Println(err.Error())
	ctx.String(status, err.Error())
}
