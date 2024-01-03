# Go + Templ + HTMX + Tailwind CSS Template

[![Deploy on Railway](https://railway.app/button.svg)](https://railway.app/template/_U7eCH?referralCode=BUOfF1)
## 🚀 Features
* [Go](https://golang.org/)
* [Echo](https://echo.labstack.com/) - Web framework
* [Templ](https://templ.guide/) - Templating engine
* [HTMX](https://htmx.org/) - Dynamic Sites without Javascript
* [Tailwind CSS](https://tailwindcss.com/) - CSS framework
* [Air](https://github.com/cosmtrek/air) - Live reload for Go apps
* [Taskfile](https://taskfile.dev/) - Task runner

## ✨ How to use

### Prerequisites
Install [Taskfile](https://taskfile.dev/) and [Air](https://github.com/cosmtrek/air#installation)
run `npm install`

### Development Mode

Development mode will watch for changes in the source code and restart the server automatically.
It will also watch for changes in the static files and recompile them automatically.
```shell
task dev
```

## 📦 Structure


```
static/: static files that are are served directly by the application server at /static/*
templates/: Reusable Templ components
views/: Views and routes for views
```

