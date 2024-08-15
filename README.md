# Noxt

Noxt is a full-stack web framework that allows you to deliver reactive UIs from the backend without writing a single line of HTML, CSS and JS.

## Get started

install the cli
```bash
$ go install github.com/noxt/cli
```

create the app
```bash
$ noxt create my_awesome_app
```

the project starter generates a hello-world app for you to get started on developing.
```bash
$ noxt run # you have to be in the project's directory
```

## Project structure

it is recommended that you adhere to the following structure since it is how the CLI is expecting to read
the project.
```txt
root
├── assets
├── app
│   ├── my_component.go
│   ├── my_page.go
│   └── subfolder1
│       └── child_component.go
├── api
│   └── login_handler.go
└── main.go
```

## How to

Define a basic app.

```go
package main

import (
	"fmt"
	"libregit.dev/noxt/noxt"
	"libregit.dev/noxt/noxt/ui"
	"libregit.dev/noxt/noxt/router"
)

func MyComponent(request noxt.Request) ui.Component {
	// component definition here
}

func MyFunction(request noxt.Request) noxt.Response {
	return noxt.Response{
		Status:  http.StatusOK,
		Headers: noxt.Headers{},
		Body:    noxt.BodyJson(map[string]string{"foo": "bar"}),
	}
}

func main() {
	app := noxt.New()
	app.Render(router.Route{
		Path:      "",
		Component: MyComponent,
		RunBefore: []router.Middleware{},
		RunAfter:  []router.Middleware{},
		Children:  []router.Route{},
	})
	app.Route(noxt.Route{
		Path:     "api",
		Function: MyFunction,
	})
	if err := app.Run(8080); err != nil {
		fmt.Errorf("Could not start server: %v\n", err.Error())
	}
}
```
