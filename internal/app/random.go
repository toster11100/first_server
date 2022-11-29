package handel

import (
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

func (h *Handel) Random(w http.ResponseWriter, r *http.Request) {
	var minrand, maxrand int64

	min := r.URL.Query().Get("min")
	max := r.URL.Query().Get("max")

	var err error
	if min == "" {
		minrand = 0
	} else {
		minrand, err = strconv.ParseInt(min, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	}
	if max == "" {
		maxrand = math.MaxInt64
	} else {
		maxrand, err = strconv.ParseInt(max, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	}

	if minrand < 0 || maxrand < 0 || minrand >= maxrand {
		http.Error(w, "400 bad request", 400)
		logrus.Errorln("func random argument is invalid")
		return
	}

	rand.Seed(time.Now().UnixNano())
	fmt.Fprintln(w, rand.Int63n(maxrand-minrand)+minrand)
}
