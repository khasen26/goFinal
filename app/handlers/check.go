package handlers

import (
	"EfN20/goFirst/foundation/web"
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

type check struct {
	log *log.Logger
}

func (c check) readiness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if n := rand.Intn(100); n%2 == 0 {
		return web.NewRequestError(errors.New("trusted error"), http.StatusBadRequest)
		// panic("forcing panic")
	}

	//localhost:8080/readiness/10
	mp := web.Params(r)
	n := mp["id"]
	fmt.Println(n)

	status := struct {
		Status string
	}{
		Status: "OK",
	}
	return web.Respond(ctx, w, status, http.StatusOK)
}
