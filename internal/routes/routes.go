package routes

import (
	"example.com/m/v2/internal/repositories/adapter"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	// ServerConfig
	// HealthHandler
	// ProductHandler
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

func (r *Router) SetRouters(repository adapter.Interface) *chi.Mux {
	r.setConfigRouters()
	r.RouterHealth(repository)
	r.RouterProduct(repository)

	return r.router
}

func (r *Router) setConfigRouters() {
	r.EnableCors()
	r.EnableLogger()
	r.EnableTimeout()
	r.EnableRecover()
	r.EnableRequestID()
	r.EnableRealIP()
}

func (r *Router) RouterHealth(repo adapter.Interface) *chi.Mux {
	handler := HealthHandler.NewRouter(repo)
	r.router.Route("/health", func(r chi.Router) {
		r.Post("/", handler.Post)
		r.Get("/", handler.Get)
		r.Put("/", handler.Put)
		r.Delete("/", handler.Delete)
		r.Options("/", handler.Options)
	})
}

func (r *Router) RouterProduct(repo adapter.Interface) *chi.Mux {
	handler := ProductHandler.NewRouter(repo)
	r.router.Route("/product", func(r chi.Router) {
		r.Post("/", handler.Post)
		r.Get("/", handler.Get)
		r.Put("/", handler.Put)
		r.Delete("/", handler.Delete)
		r.Options("/", handler.Options)
	})
	return chi.NewRouter()
}

func (r *Router) EnableTimeout() *Router {
	r.router.Use(middleware.Timeout(r.config.GetTimeout()))
	return r
}

func (r *Router) EnableCors() *Router {
	r.router.Use(r.config.Cors())
	return r
}

func (r *Router) EnableLogger() *Router {
	r.router.Use(middleware.Logger)
	return r
}

func (r *Router) EnableRecover() *Router {
	r.router.Use(middleware.Recoverer)
	return r
}

func (r *Router) EnableRequestID() *Router {
	r.router.Use(middleware.RequestID)
	return r
}

func (r *Router) EnableRealIP() *Router {
	r.router.Use(middleware.RealIP)
	return r
}
