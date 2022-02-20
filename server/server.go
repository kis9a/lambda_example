package server

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/kis9a/lambda-sls/config"
	"github.com/kis9a/lambda-sls/server/router"
)

type Server struct {
	Engine *gin.Engine
	Mode   string
	Port   string
}

func NewServer() *Server {
	var mode string
	// new config
	cfg := config.GetConfig()

	// new gin
	r := gin.New()

	// gin set mode
	if cfg.ENV != "debug" {
		mode = gin.ReleaseMode
		gin.SetMode(mode)
	}

	// init router
	router.InitRouter(r)

	return &Server{
		Engine: r,
		Mode:   mode,
		Port:   cfg.SERVER_PORT,
	}
}

func (s *Server) ListenAndServeHttp() error {
	server := http.Server{
		Addr:    ":" + s.Port,
		Handler: s.Engine,
	}
	return server.ListenAndServe()
}

func (s *Server) ListenAndServeGinProxy(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginadapter.New(s.Engine).ProxyWithContext(ctx, req)
}
