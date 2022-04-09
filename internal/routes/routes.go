package routes

import (
	"net/http"

	"github.com/go-chi/chi"
)

type Router struct {
	config *Config
	router *chi.Mux
}

func AppRouter() *Router {
	return &Router{
		config: getRouteConfig().SetTimeout(serviceConfig.GetConfig().Timeout),
		router: chi.NewRouter(),
	}
}

func (r *Router) SetRouters(handler http.Handler) *chi.Mux {

}

func (r *Router) setConfigRouters() *Config {
	return &Config{}
}

func RouterHealthCheck() *chi.Mux {
	return chi.NewRouter()
}

func RouterProduct() *chi.Mux {
	return chi.NewRouter()
}

func (r *Router) EnableTimeout(next http.Handler) http.Handler {
	return http.TimeoutHandler(next, r.config.GetTimeout(), "Request timed out")
}

func (r *Router) EnableCors(next http.Handler) http.Handler {
	return r.config.Cors(next)
}

func EnableRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"error": "Internal Server Error"}`))
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func EnableRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestId := r.Header.Get("X-Request-Id")
		if requestId == "" {
			requestId = r.Header.Get("X-Request-Id")
		}
		if requestId == "" {
			requestId = r.Header.Get("X-Request-Id")
		}
		w.Header().Set("X-Request-Id", requestId)
		next.ServeHTTP(w, r)
	})
}

func EnableRealIP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		realIP := r.Header.Get("X-Real-IP")
		if realIP == "" {
			realIP = r.Header.Get("X-Real-IP")
		}
		if realIP == "" {
			realIP = r.Header.Get("X-Real-IP")
		}
		w.Header().Set("X-Real-IP", realIP)
		next.ServeHTTP(w, r)
	})
}
