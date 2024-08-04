package main

import (
	"tag"

	"github.com/ddkwork/unison"
	"github.com/ddkwork/unison/app"
)

func main() {
	app.Run("tag", func(w *unison.Window) {
		w.Content().AddChild(tag.New().Layout())
	})
}
