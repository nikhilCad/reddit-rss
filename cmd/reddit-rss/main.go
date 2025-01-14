package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/trashhalo/reddit-rss/pkg/client"
)

func main() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: os.Getenv("SENTRY_DSN"),
	})

	if err != nil {
		panic(err)
	}

	log.Println("starting reddit-rss")

	sentryHandler := sentryhttp.New(sentryhttp.Options{})

	http.HandleFunc("/info/ping", sentryHandler.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	}))

	http.HandleFunc("/", sentryHandler.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
		httpClient := http.DefaultClient
		client.RssHandler("https://reddit.com", time.Now, httpClient, client.GetArticle, w, r)
	}))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
