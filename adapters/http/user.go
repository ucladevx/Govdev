package http

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type CacheService interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte, expiration time.Duration) error
	Del(key string) error
}

type UserController struct {
	cacheService CacheService
}

type Credentials struct {
	Username string `json:"password"`
	Password string `json:"username"`
}

func NewUserController(cacheService CacheService) *UserController {
	return &UserController{
		cacheService: cacheService,
	}
}

func (uc *UserController) Mount(g *echo.Group) {
	g.POST("/signin", uc.Signin)
	g.GET("/welcome", uc.Welcome)
	g.POST("/refresh", uc.Refresh)
}

func (uc *UserController) Signin(c echo.Context) error {
	var creds Credentials

	if err := c.Bind(&creds); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func (uc *UserController) Welcome(c echo.Context) error {
	return nil
}

func (uc *UserController) Refresh(c echo.Context) error {
	return nil
}
