package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"main/checkchain"
	"net/http"
	"os"
	"os/signal"
	"time"
)

//var ProviderUrl string

func main() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})

	r := checkchain.NewHttpServer()
	srv := &http.Server{
		Addr: "0.0.0.0:8082",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}
	go func() {
		log.Info("Server listening on 8082 port ...")
		if err := srv.ListenAndServe(); err != nil {
			log.Error(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	var wait time.Duration
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatal("error shutdown")
	}

	log.Info("shutting down")
	checkchain.Client.Close()
	os.Exit(0)
}
