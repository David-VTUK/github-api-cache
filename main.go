package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func main() {

	//resp, _ := http.Get("https://api.github.com/repos/rancher/rancher/releases/latest")
	//bodyBytes, _ := io.ReadAll(resp.Body)

	var etag string
	var lastModified string

	// this loop decrements "X-RateLimit-Remaining"
	for i := 0; i < 3; i++ {

		resp, _ := http.Get("https://api.github.com/repos/rancher/rancher/releases/latest")
		log.Info(resp.Header.Get("ETag"))
		log.Info(resp.Header.Get("Last-Modified"))
		log.Info(resp.Header.Get("X-RateLimit-Limit"))
		log.Info(resp.Header.Get("X-RateLimit-Remaining"))
		log.Info(resp.Header.Get("X-RateLimit-Used"))

		etag = resp.Header.Get("ETag")
		lastModified = resp.Header.Get("Last-Modified")

		log.Info(resp.StatusCode)
		time.Sleep(3 * time.Second)
	}

	// this loop doesn't decrement "X-RateLimit-Remaining", returns a http 304

	for i := 0; i < 3; i++ {

		client := http.Client{}
		req, _ := http.NewRequest("GET", "https://api.github.com/repos/rancher/rancher/releases/latest", nil)

		req.Header.Set("If-None-Match", etag)
		req.Header.Set("If-Modified-Since", lastModified)

		res, _ := client.Do(req)

		log.Info(res.Header.Get("ETag"))
		log.Info(res.Header.Get("Last-Modified"))
		log.Info(res.Header.Get("X-RateLimit-Limit"))
		log.Info(res.Header.Get("X-RateLimit-Remaining"))
		log.Info(res.Header.Get("X-RateLimit-Used"))
		log.Info(res.StatusCode)
		time.Sleep(3 * time.Second)
	}

}
