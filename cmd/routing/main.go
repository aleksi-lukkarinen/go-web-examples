package main

import (
	"fmt"
	"go-web-examples/internal/cmdline"
	"go-web-examples/internal/common"
	"go-web-examples/internal/environment"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	printWelcomeMessages()
	initLogging()

	args, err := cmdline.Parse()
	if err != nil {
		log.Fatal(err)
		return
	}

	env, err := environment.Parse(args)
	if err != nil {
		log.Fatal(err)
		return
	}

	bootstrapServer(env)
}

func printWelcomeMessages() {
	fmt.Println("Test Server 2022 is starting...")
	fmt.Println("")
}

func initLogging() {
	log.SetPrefix("Test server: ")
}

func bootstrapServer(e *common.RuntimeEnv) {
	router := mux.NewRouter()

	setupRoutes(router, e)
	startServer(router, e)
}

func setupRoutes(router *mux.Router, e *common.RuntimeEnv) {
	router.HandleFunc("/book/{title}/page/{page}",
		func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			title := vars["title"]
			page := vars["page"]

			fmt.Fprintf(w, "Book: %s on page %s\n", title, page)
		},
	)

	staticServer := http.FileServer(http.Dir(e.StaticFilesLocalPath))
	handler := http.StripPrefix(common.URL_PART_STATIC, staticServer)
	router.PathPrefix(common.URL_PART_STATIC).Handler(handler)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		const defaultResponse = `<!DOCTYPE html>
	   <html>
	   	<body>
	   		<img src="static/hoimaailma.png" /><br />
	   		%s
	   	</body>
	   </html>`

		fmt.Fprintf(w, defaultResponse, time.Now().String())
	})
}

func startServer(router *mux.Router, e *common.RuntimeEnv) {
	address := fmt.Sprintf("%s:%s", e.ServerIP, e.ServerPort)
	log.Fatal(http.ListenAndServe(address, router))
}
