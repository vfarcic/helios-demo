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
	sdk.Initialize("other-app", "65e7de34aaee5efb5be8", sdk.WithEnvironment("prod"))
	log.SetOutput(os.Stderr)

	// Server
	log.Println("Starting server...")
	router := gin.New()
	router.GET("/", rootHandler)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	router.Run(fmt.Sprintf(":%s", port))
}

func rootHandler(ctx *gin.Context) {
	if len(ctx.Query("fail")) > 0 {
		ctx.String(http.StatusBadRequest, "Something terrible happened")
		return
	}
	version := os.Getenv("VERSION")
	output := os.Getenv("MESSAGE")
	if len(output) == 0 {
		output = "This is a silly demo"
	}
	if len(version) > 0 {
		output = fmt.Sprintf("%s version %s", output, version)
	}
	if len(ctx.Query("html")) > 0 {
		output = fmt.Sprintf("<h1>%s</h1>", output)
	}
	output = fmt.Sprintf("%s\n", output)
	ctx.String(http.StatusOK, output)
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
