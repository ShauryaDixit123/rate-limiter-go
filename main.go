package main

import (
	"asgn/limiter"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/rate-limited-route", rateHandledHandler)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", limiter.LoadEnvVariable("API_PORT")), limiter.RateLimitByIPMiddleware(mux)); err != nil {
		log.Fatalf("SERVER_START_FAILED, ERR: %s", err.Error())
	}

}

func rateHandledHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("************INSIDE HANDLER**************")
}
