package second

import (
	"github.com/kataras/iris"
	"log"
)

func main() {
	var app *iris.Application
	app = iris.Default()
	app.Get("/ping", Ping)

	if err := app.Run(iris.Addr(":8080")); err != nil {
		log.Fatalln(err)
	}

}

func Ping(ctx iris.Context) {
	var p Point
	err, z := p.initt(2, 2)
	if err == nil {
		log.Println(z)
	}
	_, _ = ctx.JSON(iris.Map{
		"message": "pong",
		"mes":     "pong",
	})
}

type Point struct {
	x int
	y int
}

func (p Point) initt(x int, y int) (error, int) {
	p.x = x
	p.y = y
	return nil, x + y
}
