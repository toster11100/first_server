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

func Random(w http.ResponseWriter, r *http.Request) {
	var minrand, maxrand int64

	min := r.URL.Query().Get("min")
	max := r.URL.Query().Get("max")

	if min == "" {
		minrand = 0
	} else {
		minrand, _ = strconv.ParseInt(min, 10, 64)
	}
	if max == "" {
		maxrand = math.MaxInt64
	} else {
		maxrand, _ = strconv.ParseInt(max, 10, 64)
	}

	if minrand < 0 || maxrand < 0 || minrand >= maxrand {
		http.Error(w, "400 bad request", 400)
		logrus.Errorln("func random argument is invalid")
		return
	}

	rand.Seed(time.Now().UnixNano())
	fmt.Fprintln(w, rand.Int63n(maxrand-minrand)+minrand)
}
