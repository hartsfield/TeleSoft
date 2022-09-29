package main

import (
	"context"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

var (
	templates = template.Must(template.New("main").ParseGlob("internal/*/*.tmpl"))
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/sendMail", sendMail)
	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	// Server configuration
	srv := &http.Server{
		// in production only ust SSL
		Addr:              ":9002",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       5 * time.Second,
	}

	ctx, cancelCtx := context.WithCancel(context.Background())

	go func() {
		err := srv.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server two closed\n")
		} else if err != nil {
			fmt.Printf("error listening for server two: %s\n", err)
		}
		cancelCtx()
	}()

	fmt.Println("Server started @ " + srv.Addr)
	<-ctx.Done()
}
