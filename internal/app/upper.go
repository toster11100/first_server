package handel

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

func (h *Handel) Upper(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Upper")
	s := r.URL.Query().Get("s")
	if s == "" {
		s = r.URL.Query().Get("string")
	}
	if s == "" {
		http.Error(w, "400 bad request", 400)
		logrus.Errorln("func upper argument is invalid")
		return
	}
	fmt.Fprint(w, strings.ToUpper(s))
}

//{
//switch {
//case s != "":
//fmt.Fprintln(w, strings.ToUpper(s))
//case str != "":
//fmt.Fprintln(w, strings.ToUpper(str))
//default:
//http.Error(w, "400 bad request", 400)
//logrus.Errorln("func upper argument is invalid")
//}
//}
