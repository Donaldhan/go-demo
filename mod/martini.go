package mod

import "github.com/go-martini/martini"

func MartiniDemo() {
	m := martini.Classic()
	// ... middleware and routing goes here
	m.Run()
	m.Get("/", func() (int, string) {
		return 418, "i'm a teapot" // HTTP 418 : "i'm a teapot"
	})
}
