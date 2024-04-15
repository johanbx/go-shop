# Go web app template
go • gin • air • live-reload • picocssv2 • htmx • docker


## Quick start

Start developing with `make dev` and visit http://localhost:8080. Now whenever you are changing any
code it will recompile and reload the website. To see the logs run `make logs` and to shut down everything
run `make clean`.

### Production mode

To build for production and then run the final image, run `make prod` and visit http://localhost:8080

## Requirement

* [docker](https://www.docker.com/)


## Tools and techniques used

| Tool | Description | Website |
| -------- | ------- | ------- |
| Go | Programming language | https://go.dev/ |
| gin | Web framework for Go | https://github.com/gin-gonic/gin |
| Air | Live reload for Go apps. Filewatchers that recompile project and re-run it when file changes | https://github.com/cosmtrek/air |
| Live reload | Homebaked solution for live reloading client browser on recompilation from Air | ./cmd/livereload/main.go |
| Pico CSS | Effortless styling without classes | https://picocss.com/ |
| htmx | Modern web app functionality with javascript using html tags | https://htmx.org/ |
| docker | Your trusty container friend | https://www.docker.com/ |
