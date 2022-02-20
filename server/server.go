package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/kis9a/lambda-todo/config"
	"go.uber.org/zap"
)

type Server struct {
	Engine *gin.Engine
	Mode   string
	Port   string
}

func NewServer(cfg *config.Config, logger *zap.Logger) *Server {
	var mode string
	var port string

	// new gin
	r := gin.New()

	// gin set mode
	if cfg.ENV != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}

	// gin server port
	if cfg.SERVER_PORT == "" {
		port = ":4000"
	}

	// init router
	initRouter(r)

	logger.Info(fmt.Sprintf("%v+", r.Routes()))

	// init middlewares
	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	r.Use(ginzap.RecoveryWithZap(logger, true))

	return &Server{
		Engine: r,
		Mode:   mode,
		Port:   port,
	}
}

func (h *Server) ListenAndServeHttp() error {
	s := http.Server{
		Addr:    h.Port,
		Handler: h.Engine,
	}
	return s.ListenAndServe()
}

func (h *Server) ListenAndServeGinProxy(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginadapter.New(h.Engine).ProxyWithContext(ctx, req)
}
