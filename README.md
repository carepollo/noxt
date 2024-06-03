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
project.
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
