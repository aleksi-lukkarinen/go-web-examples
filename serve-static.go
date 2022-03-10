package main

import (
	"fmt"
	"net/http"
	"time"
)

const STATIC_FILES_PATH = "static"

const defaultResponse = `<!DOCTYPE html>
<html>
	<body>
		<img src='` + STATIC_FILES_PATH + `/hoi-maailma.png' /><br />
		%s
	</body>
</html>`

func setupDynamicRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, defaultResponse, time.Now().String())
	})
}

func setupStaticServer() {
	const staticFilesLocalPath = STATIC_FILES_PATH + "/"
	staticServer := http.FileServer(http.Dir(staticFilesLocalPath))

	const staticPrefix = "/" + staticFilesLocalPath
	http.Handle(staticPrefix, http.StripPrefix(staticPrefix, staticServer))
}

func startServer() {
	http.ListenAndServe(":8000", nil)
}

func main() {
	setupDynamicRoutes()
	setupStaticServer()
	startServer()
}
