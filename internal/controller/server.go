package controller

import (
	"app/internal/domain"
	"app/internal/repository"
	"app/internal/service/button"
	"app/internal/service/user"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/valyala/fasthttp"
	"gorm.io/gorm"
)

func Serve() {
	repo, err := repository.New("test.db")
	if err != nil {
		panic(err)
	}

	buttonService := button.New(repo)
	userSevice := user.New(repo)

	slog.SetLogLoggerLevel(slog.LevelDebug)

	err = fasthttp.ListenAndServe(":8080", func(ctx *fasthttp.RequestCtx) {
		token := string(ctx.Request.Header.Cookie("baton-session"))
		user, err := userSevice.GetUser(ctx, token)
		if errors.Is(err, gorm.ErrRecordNotFound) { // FIXME: убрать после тестов
			user, err = userSevice.CreateUser(ctx)
			if err != nil {
				ctx.SetBodyString(err.Error())
				ctx.SetStatusCode(http.StatusInternalServerError)

				return
			}

			slog.Debug(
				"user new",
				slog.Any("user", user),
				slog.String("token", token),
			)

			user, err = userSevice.GetUser(ctx, user.Token)
			if err != nil {
				ctx.SetBodyString(err.Error())
				ctx.SetStatusCode(http.StatusInternalServerError)

				return
			}

			c := fasthttp.Cookie{}
			c.SetHTTPOnly(true)
			c.SetKey("baton-session")
			c.SetPath("/")
			c.SetValue(user.Token)
			ctx.Response.Header.SetCookie(&c)
		}
		if err != nil {
			ctx.SetBodyString(err.Error())
			ctx.SetStatusCode(http.StatusInternalServerError)

			return
		}

		slog.Debug(
			"user final",
			slog.Any("user", user),
			slog.String("token", token),
		)

		b, err := buttonService.PressButton(ctx, domain.User{
			ID: user.ID,
		})
		if err != nil {
			ctx.SetBodyString(err.Error())
			ctx.SetStatusCode(http.StatusInternalServerError)

			return
		}

		ctx.SetBodyString(fmt.Sprintln(b.Count))
		ctx.SetStatusCode(http.StatusOK)
	})
	if err != nil {
		panic(err)
	}
}
