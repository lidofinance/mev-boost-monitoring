package metrics

import (
	"fmt"
	"regexp"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var re = regexp.MustCompile(`/-|\ /g`)

type Store struct {
	Prometheus *prometheus.Registry
	BuildInfo  prometheus.Counter
}

func New(promRegistry *prometheus.Registry, appName string, env string) *Store {
	store := &Store{
		Prometheus: promRegistry,
		BuildInfo: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_METRIC_BUILD_INFO", re.ReplaceAllString(appName, `_`)),
			Help: "Build information",
			ConstLabels: prometheus.Labels{
				"name": appName,
				"env":  env,
			},
		}),
	}

	_ = store.Prometheus.Register(store.BuildInfo)

	return store
}
