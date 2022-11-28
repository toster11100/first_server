package main

import (
	"net/http"

	handel "main.go/internal/app"
	"main.go/internal/flags"

	"github.com/sirupsen/logrus"
)

func main() {
	flags := flags.New()

	http.HandleFunc("/random", handel.Random)
	http.HandleFunc("/upper", handel.Upper)

	logrus.Fatal(http.ListenAndServe(*flags.Port, nil))
}
