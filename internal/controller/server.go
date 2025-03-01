package controller

import (
	"app/internal/repository"
	"app/internal/service/button"
	"app/internal/service/user"
	"log/slog"
	"net/http"
	"os"

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
}

func New(addr string, debug bool) (*Controller, error) {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}))

	repo, err := repository.New("test.db")
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

func (c Controller) Serve() error {
	handleIndex := (&fasthttp.FS{
		Root:        "internal/controller",
		PathRewrite: func(ctx *fasthttp.RequestCtx) []byte { return []byte("/index.html") },
		SkipCache:   true,
	}).NewRequestHandler()

	return fasthttp.ListenAndServe(c.addr, func(ctx *fasthttp.RequestCtx) {
		p := string(ctx.Path())
		ctx.SetContentType("application/json")

		switch {
		case p == "/" && ctx.IsGet():
			ctx.SetStatusCode(http.StatusOK)
			ctx.SetContentType("text/html")
			if c.debug {
				handleIndex(ctx)
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
	})
}
