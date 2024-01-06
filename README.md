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

### Generating Code
The server binary embeds the static file directory and serves anything contained within under `/static`. To build the server
binary you should first generate the `tailwind`, `templ`, and any other necessary generated code for your app:

```
task generate
```

### Building the binary:
Once all files have been generated, the binary can be built just like any other go binary:
```
go build -o build/server main.go
```
Or you can use the `build` task which will generate code before building the binary:
```
task build
```
### Containerized Builds

You can easily containerize the build using the provided [Dockerfile](Dockerfile). Make sure all code is generated before triggering
a build of the container, or use the `container` task which will generate any code first:

```
task container container_name=my-app container_tag=test
```
#### Railway

Deploying to Railway doesn't require any special configuration. You can just fork this repository and configure Railway to use
your forked repository. Railway will then automatically build and deploy your main branch on every change. You can click the
`Deploy to Railway` badge at the top of the README to streamline this process.

**Note**: This template was designed to deploy to Railway in mind but there is nothing prohibiting this template from working
elsewhere. The binary is just a standard go http server. The Dockerfile should work just about anywhere the runs containers.
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

