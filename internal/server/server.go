package server

import (
	"app/internal/domain"
	"app/internal/repository"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/valyala/fasthttp"
	"gorm.io/gorm"
)

func Serve() {
	repo, err := repository.New("test.db")
	if err != nil {
		panic(err)
	}

	slog.SetLogLoggerLevel(slog.LevelDebug)

	err = fasthttp.ListenAndServe(":8080", func(ctx *fasthttp.RequestCtx) {
		y, m, d := time.Now().Date()

		b := domain.ButtonPressed{
			Year:  y,
			Month: int(m),
			Day:   d,
		}

		b2, err := repo.GetButton(ctx, b)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.SetBodyString(err.Error())
			ctx.SetStatusCode(http.StatusInternalServerError)

			return
		}

		slog.Debug(
			"b2 value",
			slog.Any("b", b),
			slog.Any("b2", b2),
		)

		b = domain.ButtonPressed{
			Year:  y,
			Month: int(m),
			Day:   d,
			Count: b2.Count + 1,
		}

		err = repo.ButtonUpdate(ctx, b)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
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
