package http

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/ucladevx/govdev/util/remember"
)

const (
	time15m = 900 * time.Second
	time7d  = 604800 * time.Second
	time24h = 86400 * time.Second
	b1      = 1000000000
)

type CacheService interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte, expiration time.Duration) error
	Del(key string) error
}

type UserController struct {
	cacheService CacheService
}

// Credentials struct helps serialize login credentials
type Credentials struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

// TODO: make Sessions fixed size, easily convert from struct to []byte
type Session struct {
	Username      string
	RememberToken string
	RefreshToken  string
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

	// TODO: password authentication

	// TODO: refresh token
	token, _ := remember.RememberToken()
	err := uc.cacheService.Set(token, []byte(creds.Username), 900000000000)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// set cookie
	cookie := &http.Cookie{
		Name:     "remember-token",
		Value:    token,
		HttpOnly: true,
	}
	c.SetCookie(cookie)
	return nil
}

func (uc *UserController) Welcome(c echo.Context) error {
	cookie, err := c.Cookie("remember-token")
	if err != nil {
		return err
	}

	username_bytes, err := uc.cacheService.Get(cookie.Value)
	if err != nil {
		return err
	}
	username := string(username_bytes)
	return c.String(http.StatusOK, username)
}

func (uc *UserController) Refresh(c echo.Context) error {
	return nil
}
