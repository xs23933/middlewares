package logger

import "github.com/xs23933/core"

type Config struct{}

func Logger() core.HandlerFunc {
	return core.HandlerFunc(func(c *core.Ctx) {
		c.SendString("im logger")
		c.Next()
		c.SendString("im logger done")
	})
}
