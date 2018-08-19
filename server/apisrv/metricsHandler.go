package apisrv

import (
	"expvar"
	"fmt"
	"net/http"

	"github.com/qbeon/webwire-messenger/server/apisrv/util"
)

type MetricsHandler struct{}

func NewMetricsHandler() *MetricsHandler {
	return &MetricsHandler{}
}

// ServeHTTP handles incoming metrics export requests
func (handler *MetricsHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	first := true
	report := func(key string, value interface{}) {
		if !first {
			fmt.Fprintf(w, ",\n")
		}
		first = false
		if str, ok := value.(string); ok {
			fmt.Fprintf(w, "%q: %q", key, str)
		} else {
			fmt.Fprintf(w, "%q: %v", key, value)
		}
	}

	fmt.Fprintf(w, "{\n")

	v, _ := expvar.Get("GraphQueryLatencyAvg").(*expvar.Int)
	v.Set(
		v.Value() / util.Max(
			expvar.Get("GraphQueries").(*expvar.Int).Value(), 1,
		),
	)

	expvar.Do(func(kv expvar.KeyValue) {
		v, ok := expvar.Get(kv.Key).(*expvar.Int)
		report(kv.Key, kv.Value)
		if ok {
			v.Set(0)
		}
	})
	fmt.Fprintf(w, "\n}\n")
}
