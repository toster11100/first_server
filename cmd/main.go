package main

import (
	"net/http"

	handel "main.go/internal/app"
	"main.go/internal/flags"

	"github.com/sirupsen/logrus"
)

func main() {
	flags := flags.New()
	db := 0
	h := handel.New(db)

	http.HandleFunc("/random", h.Random)
	http.HandleFunc("/upper", h.Upper)

	logrus.Info("Starting server")
	logrus.Fatal(http.ListenAndServe(*flags.Port, nil))
}
