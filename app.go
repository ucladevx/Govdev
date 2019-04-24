package govdev

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
	http_controllers "github.com/ucladevx/govdev/adapters/http"
	"github.com/ucladevx/govdev/services"
	"github.com/ucladevx/govdev/stores/postgresql"
	"github.com/ucladevx/govdev/stores/redis"
)

func Start(conf Config) {
	r := redis.NewConnection(
		conf.Cache.Host,
		conf.Cache.Port,
		conf.Cache.Password,
		conf.Cache.DBName,
	)
	defer r.Close()

	db := postgresql.NewConnection(
		conf.Database.Engine,
		conf.Database.User,
		conf.Database.Password,
		conf.Database.DBName,
		conf.Database.Port,
		conf.Database.Host,
	)
	defer db.Close()
	must(db.Ping())

	userStore := postgresql.NewUserStore(db)
	postgresql.CreateTables(
		userStore,
	)

	cacheStore := redis.NewRedisStore(r)

	cacheService := services.NewCacheService(cacheStore)

	userController := http_controllers.NewUserController(cacheService)

	app := initServer(conf)
	app.GET("/", hello)
	userController.Mount(app.Group("/user"))

	go func() {
		app.Logger.Fatal(app.Start(":" + conf.Port))
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, os.Kill)
	<-quit
	fmt.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := app.Shutdown(ctx); err != nil {
		app.Logger.Fatal(err)
	}
	fmt.Println("Shutdown govdev server. Goodbye.")
}

func initServer(conf Config) *echo.Echo {
	app := echo.New()
	app.HideBanner = true
	app.Debug = conf.Debug

	app.Use(middleware.Gzip())
	app.Use(middleware.Secure())
	app.Use(middleware.RemoveTrailingSlash())
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `${time_rfc3339_nano} ${method} {id":"${id}","remote_ip":"${remote_ip}",` +
			`"uri":"${uri}","status":${status},"latency":${latency},` +
			`"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
			`"bytes_out":${bytes_out}}` + "\n",
	}))

	return app
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// Function denotes that this error must not exist
func must(err error) {
	if err != nil {
		panic(err)
	}
}
