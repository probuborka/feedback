package http

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
)

type handler struct {
	feedback serviceFeedback
	// metric         metric
	log *logrus.Logger
}

func New(log *logrus.Logger, feedback serviceFeedback) *handler {
	return &handler{
		// recommendation: recommendation,
		// metric:         metric,
		log:      log,
		feedback: feedback,
	}
}

func (h handler) Init() http.Handler {
	r := http.NewServeMux()

	//swagger UI
	r.Handle("/swagger/", httpSwagger.WrapHandler)

	//metrics
	r.Handle("/metrics", promhttp.Handler())

	//feedback
	r.HandleFunc("POST /feedback", h.postFeedback)

	//middleware
	stack := []middleware{
		//h.recordMetrics,
		//h.logging,
	}

	return compileMiddleware(r, stack)
}
