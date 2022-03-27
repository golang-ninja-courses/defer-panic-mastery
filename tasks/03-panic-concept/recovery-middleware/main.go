package main

import (
	"net/http"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.INFO)

	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
		LogLevel:  log.ERROR,
	}))

	e.GET("/recovered", func(c echo.Context) error {
		panic("no panic!")
	})

	e.GET("/", func(c echo.Context) error {
		var wg sync.WaitGroup
		wg.Add(3)

		for i := 0; i < 3; i++ {
			i := i
			go func() {
				defer wg.Done()

				select {
				case <-c.Request().Context().Done():
					return
				case <-time.After(time.Second):
				}

				if i == 1 {
					panic("internal logic error")
				}
				c.Logger().Infof("processed %d", i)
			}()
		}

		wg.Wait()
		return c.String(http.StatusOK, "DONE")
	})

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
