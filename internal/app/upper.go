package handel

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

func Upper(w http.ResponseWriter, r *http.Request) {
	s := r.URL.Query().Get("s")
	str := r.URL.Query().Get("string")
	switch {
	case s != "":
		fmt.Fprintln(w, strings.ToUpper(s))
	case str != "":
		fmt.Fprintln(w, strings.ToUpper(str))
	default:
		http.Error(w, "400 bad request", 400)
		logrus.Errorln("func upper argument is invalid")
	}
}
