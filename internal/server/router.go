package server

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"velo/internal/config"
)

type Server struct {
	cfg    config.Config
	router *router.Router
	stats  *Stats
}

func New(cfg config.Config) *Server {
	r := router.New()
	s := &Server{cfg: cfg, router: r, stats: NewStats()}

	go s.stats.startDiskMonitor(cfg.CacheDir)

	r.GET("/stats", s.HandleStats)
	r.GET("/{params:*}", s.HandleImage)
	return s
}

func (s *Server) ListenAndServe() error {
	return fasthttp.ListenAndServe(s.cfg.Addr, s.router.Handler)
}
