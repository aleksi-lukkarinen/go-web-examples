package main

import (
	"flag"
	"fmt"
	"go-web-examples/internal/common"
	"go-web-examples/internal/environment"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	printWelcomeMessages()
	initLogging()

	args, err := parseCommandlineArguments()
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

func parseCommandlineArguments() (common.CmdlineArgs, error) {
	args := common.CmdlineArgs{}

	envFilePathPtr := flag.String(
		"env",
		"",
		"A path to a .env file that sets environment variables")

	flag.Parse()

	args.EnvFilePath = ""
	if len(*envFilePathPtr) > 0 {
		p, err := filepath.Abs(strings.TrimSpace(*envFilePathPtr))
		if err != nil {
			p = filepath.FromSlash(strings.TrimSpace(*envFilePathPtr))
		}

		args.EnvFilePath = p
	}

	return args, nil
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

	staticUrl := "/static/"
	staticServer := http.FileServer(http.Dir(e.StaticFilesLocalPath))
	handler := http.StripPrefix(staticUrl, staticServer)
	router.PathPrefix(staticUrl).Handler(handler)

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
