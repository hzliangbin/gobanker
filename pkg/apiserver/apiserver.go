package apiserver

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/hzliangbin/gobanker/pkg/apiserver/healthz"
	"github.com/hzliangbin/gobanker/pkg/blog"
)


type APIServer struct {
	*Config

	Server *http.Server
	SimpleServer *http.Server
}

func registerAPI(cfg *Config) http.Handler {
	router := gin.New()

	return router
}

func NewAPIServerWithOpts(cfg *Config) *APIServer {
	router := registerAPI(cfg)

	s := &APIServer{
		Config:       cfg,
		Server:       &http.Server{
			Addr:              fmt.Sprintf("%s:%d", cfg.BindAddr, cfg.InsecurePort),
			Handler:           router,
		},
	}

	if cfg.SecurePort != 0 {
		s.Server.Addr = fmt.Sprintf("%s:%d", s.Config.BindAddr, s.Config.SecurePort)
	}

	return withiSimpleServer(s)
}

func withiSimpleServer(s *APIServer) *APIServer {
	router := gin.New()
	router.Use(gin.Recovery())

	// The url pointing to API definition
	url := ginSwagger.URL("/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	// TODO
	router.GET("/healthz", healthz.Check)

	s.SimpleServer = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", s.Config.BindAddr, s.Config.GenericPort),
		Handler: router,
	}

	return s
}

func (s *APIServer) Initialize() error {
	return nil
}

func (s *APIServer) Run() {
	var err error

	if s.Config.SecurePort != 0 {
		err = s.Server.ListenAndServeTLS(s.Config.TlsCert, s.Config.TlsKey)
	} else {
		err = s.Server.ListenAndServe()
	}
	if err != nil {
		blog.Fatal("banker api-server failed to start: %s", err)
	}
}
