package controller

import (
	"app/internal/repository"
	"app/internal/service/button"
	"app/internal/service/user"
	"context"
	"log/slog"
	"net/http"

	"github.com/valyala/fasthttp"

	_ "embed"
)

//go:embed index.html
var index_html_body []byte

//go:embed logo.png
var logo_body []byte

type Controller struct {
	addr  string
	debug bool

	logger *slog.Logger

	buttonService *button.Service
	userSevice    *user.Service

	handleIndex fasthttp.RequestHandler
}

func New(logger *slog.Logger, addr string, debug bool, dbType, dbDNS string) (*Controller, error) {
	repo, err := repository.New(logger, dbType, dbDNS)
	if err != nil {
		return nil, err
	}

	buttonService := button.New(repo)
	userSevice := user.New(repo)

	return &Controller{
		addr:  addr,
		debug: debug,

		logger: logger,

		buttonService: buttonService,
		userSevice:    userSevice,
	}, nil
}

func (c Controller) Serve(ctx context.Context) error {
	c.handleIndex = (&fasthttp.FS{
		Root:        "internal/controller",
		PathRewrite: func(ctx *fasthttp.RequestCtx) []byte { return []byte("/index.html") },
		SkipCache:   true,
	}).NewRequestHandler()

	server := &fasthttp.Server{
		Handler: c.handler,
	}

	go func() {
		<-ctx.Done()
		err := server.Shutdown()
		if err != nil {
			c.logger.Error("shutdown http", slog.Any("error", err))
		}
	}()

	err := server.ListenAndServe(c.addr)
	if err != nil {
		return err
	}

	return nil
}

func (c *Controller) handler(ctx *fasthttp.RequestCtx) {
	p := string(ctx.Path())
	ctx.SetContentType("application/json")

	if c.debug {
		c.logger.Debug("http request", "path", p, "method", string(ctx.Request.Header.Method()))
	}

	switch {
	case p == "/" && ctx.IsGet():
		ctx.SetStatusCode(http.StatusOK)
		ctx.SetContentType("text/html")
		if c.debug {
			c.handleIndex(ctx)
		} else {
			ctx.SetBody(index_html_body)
		}
	case p == "/logo.png" && ctx.IsGet():
		ctx.SetStatusCode(http.StatusOK)
		ctx.SetContentType("image/png")
		ctx.SetBody(logo_body)
	case p == "/api/user" && ctx.IsGet():
		c.GetUser(ctx)
	case p == "/api/user" && ctx.IsPost():
		c.CreateUser(ctx)
	case p == "/api/button" && ctx.IsGet():
		c.Buttons(ctx)
	case p == "/api/button" && ctx.IsPost():
		c.PressButton(ctx)
	default:
		ctx.SetStatusCode(http.StatusNoContent)
	}
}
