package main

import (
	"errors"
	"log"
	"os"

	"github.com/go-resty/resty/v2"
	gin "github.com/helios/go-sdk/proxy-libs/heliosgin"
	http "github.com/helios/go-sdk/proxy-libs/helioshttp"
)

func pingHandler(ctx *gin.Context) {
	req := resty.New().R().SetHeader("Content-Type", "application/text")
	url := ctx.Query("url")
	if len(url) == 0 {
		url = os.Getenv("PING_URL")
		if len(url) == 0 {
			httpErrorBadRequest(errors.New("url is empty"), ctx)
			return
		}
	}
	log.Printf("Sending a ping to %s", url)
	resp, err := req.Get(url)
	if err != nil {
		httpErrorBadRequest(err, ctx)
		return
	}
	log.Println(resp.String())
	ctx.String(http.StatusOK, resp.String())
}
