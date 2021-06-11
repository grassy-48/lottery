package main

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	//	"github.com/labstack/echo/v4/middleware"
	"lottery/controllers/point"
	"lottery/controllers/raffle"
	"lottery/controllers/user"
)

type (
	PostUser struct {
		Name       string `form:"name" validate:"required,min=1,max=32"`
		CircleName string `form:"circleName" validate:"omitempty,min=1,max=32"`
		Mail       string `form:"mail" validate:"required,email,min=1,max=64"`
		CType      string `form:"type" validate:"omitempty,oneof=participant creator"`
		Place      string `form:"place" validate:"omitempty,oneof=onl evt"`
	}
	GetUser struct {
		UserID int    `query:"user_id" validate:"omitempty,gte=1"`
		Mail   string `query:"mail" validate:"omitempty,email"`
	}
	StorePoint struct {
		Code   string `form:"code" validate:"required,len=16"`
		UserID int    `param:"userID" validate:"required,gte=1"`
	}
	CheckRaffle struct {
		UserID int `param:"userID" validate:"required,gte=1"`
	}
	DrawRaffle struct {
		UserID int `param:"userID" validate:"required,gte=1"`
	}
	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.ERROR)
	//	e.Use(middleware.Logger())
	e.Validator = &CustomValidator{validator: validator.New()}

	prefix := "/api/v1"
	e.GET(prefix+"/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST(prefix+"/users", func(c echo.Context) (err error) {
		u := new(PostUser)
		if err = c.Bind(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err = c.Validate(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return user.Create(c)
	})

	e.GET(prefix+"/users", func(c echo.Context) (err error) {
		u := new(GetUser)
		if err = c.Bind(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err = c.Validate(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return user.Get(c)
	})

	e.POST(prefix+"/users/:userID/points", func(c echo.Context) (err error) {
		u := new(StorePoint)
		if err = c.Bind(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err = c.Validate(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return point.Store(c)
	})

	e.GET(prefix+"/users/:userID/draw", func(c echo.Context) (err error) {
		u := new(CheckRaffle)
		if err = c.Bind(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err = c.Validate(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return raffle.Check(c)
	})

	e.PUT(prefix+"/users/:userID/draw/evt", func(c echo.Context) (err error) {
		u := new(DrawRaffle)
		if err = c.Bind(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err = c.Validate(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return raffle.DrawEvt(c)
	})

	e.PUT(prefix+"/users/:userID/draw/onl", func(c echo.Context) (err error) {
		u := new(DrawRaffle)
		if err = c.Bind(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err = c.Validate(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return raffle.DrawOnl(c)
	})

	e.Logger.Fatal(e.Start(":12111"))
}
