package mod

import "github.com/Donaldhan/martini"

func MartiniDemo() {
	m := martini.Classic()
	// ... middleware.md and routing goes here
	m.Run()
	m.Get("/", func() (int, string) {
		return 418, "i'm a teapot" // HTTP 418 : "i'm a teapot"
	})
}
