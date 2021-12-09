package main

import (
	"log"
	"net/http"

	components "github.com/JakWai01/notes-special-topics-for-web-applications/web-shop/pkg/components"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func main() {

	app.Route("/", &components.MyComponent{})
	app.RunWhenOnBrowser()

	http.Handle("/", &app.Handler{
		Name:            "Hello",
		Description:     "An Hello World! example",
		LoadingLabel:    "Send IT with airdrip",
		Author:          "Jakob Waibel",
		ThemeColor:      "#151515",
		BackgroundColor: "#151515",
		Styles: []string{
			"https://unpkg.com/@patternfly/patternfly@4.135.2/patternfly.css",
			"https://unpkg.com/@patternfly/patternfly@4.135.2/patternfly-addons.css",
		},
		Title: "Airdrip",
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
