package controller

import (
	"app/internal/repository"
	"app/internal/service/button"
	"app/internal/service/user"
	"log/slog"
	"net/http"
	"os"

	"github.com/valyala/fasthttp"
)

type Controller struct {
	buttonService *button.Service
	userSevice    *user.Service
	logger        *slog.Logger
}

func New() (*Controller, error) {
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
		buttonService: buttonService,
		userSevice:    userSevice,
		logger:        logger,
	}, nil
}

func (c Controller) Serve() error {
	return fasthttp.ListenAndServe(":8080", func(ctx *fasthttp.RequestCtx) {
		p := string(ctx.Path())
		ctx.SetContentType("application/json")

		switch {
		case p == "/" && ctx.IsGet():
			ctx.SetStatusCode(http.StatusOK)
			ctx.SetContentType("text/html")
			ctx.SendFile("internal/controller/index.html")
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
