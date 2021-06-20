package main

import (
	"net/http"

	"github.com/go-playground/validator"
	//"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"lottery/controllers/code"
	"lottery/controllers/point"
	"lottery/controllers/raffle"
	"lottery/controllers/user"

	"github.com/labstack/echo/v4/middleware"
)

type (
	PostUser struct {
		Name       string `form:"name" validate:"omitempty,min=1,max=32"`
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
	CheckStorePoint struct {
		Code   string `query:"code" validate:"required,len=16"`
		UserID int    `param:"userID" validate:"required,gte=1"`
	}
	CheckRaffle struct {
		UserID int `param:"userID" validate:"required,gte=1"`
	}
	DrawRaffle struct {
		UserID int `param:"userID" validate:"required,gte=1"`
	}
	CheckCode struct {
		Code string `query:"code" validate:"required,len=16"`
	}
	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.ERROR)
	e.Use(middleware.Logger())
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
		if err = user.Create(c); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return err
	})

	e.GET(prefix+"/users", func(c echo.Context) (err error) {
		u := new(GetUser)
		if err = c.Bind(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err = c.Validate(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err = user.Get(c); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return err
	})

	e.GET(prefix+"/users/:userID/points", func(c echo.Context) (err error) {
		u := new(CheckStorePoint)
		if err = c.Bind(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err = c.Validate(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err = point.CheckStore(c); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return err
	})

	e.POST(prefix+"/users/:userID/points", func(c echo.Context) (err error) {
		u := new(StorePoint)
		if err = c.Bind(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err = c.Validate(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err = point.Store(c); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return err
	})

	e.GET(prefix+"/users/:userID/draw", func(c echo.Context) (err error) {
		u := new(CheckRaffle)
		if err = c.Bind(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err = c.Validate(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err = raffle.Check(c); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return err
	})

	e.PATCH(prefix+"/users/:userID/draw/evt", func(c echo.Context) (err error) {
		u := new(DrawRaffle)
		if err = c.Bind(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err = c.Validate(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err = raffle.DrawEvt(c); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return err
	})

	e.PATCH(prefix+"/users/:userID/draw/onl", func(c echo.Context) (err error) {
		u := new(DrawRaffle)
		if err = c.Bind(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err = c.Validate(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err = raffle.DrawOnl(c); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return err
	})

	e.GET(prefix+"/codes", func(c echo.Context) (err error) {
		u := new(CheckCode)
		if err = c.Bind(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err = c.Validate(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err = code.Check(c); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return err
	})

	e.Logger.Fatal(e.Start(":12111"))
}
