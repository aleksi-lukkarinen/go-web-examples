package main

import (
	"flag"
	"fmt"
	"go-web-examples/internal/common"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type cmdlineArgs struct {
	envFilePath string
}

type runtimeEnv struct {
	workingDirectory     string
	staticFilesLocalPath string
	serverPort           string
}

func main() {
	printWelcomeMessages()
	initLogging()

	args, err := parseCommandlineArguments()
	if err != nil {
		log.Fatal(err)
		return
	}

	env, err := initEnvironment(args)
	if err != nil {
		log.Fatal(err)
		return
	}

	printEnvironment(env)

	bootstrapServer(env)
}

func printWelcomeMessages() {
	fmt.Println("Test Server 2022 is starting...")
	fmt.Println("")
}

func initLogging() {

}

func parseCommandlineArguments() (cmdlineArgs, error) {
	args := cmdlineArgs{}

	envFilePathPtr := flag.String(
		"env",
		"",
		"A path to a .env file that sets environment variables")

	flag.Parse()

	args.envFilePath = ""
	if len(*envFilePathPtr) > 0 {
		args.envFilePath =
			filepath.FromSlash(strings.TrimSpace(*envFilePathPtr))
	}

	return args, nil
}

func initEnvironment(args cmdlineArgs) (*runtimeEnv, error) {
	if len(args.envFilePath) > 0 {
		fmt.Printf("Using environment file at: %s\n", args.envFilePath)

		err := godotenv.Load(args.envFilePath)
		if err != nil {
			return nil, err
		}
	}

	env := runtimeEnv{}

	path, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	env.workingDirectory = path

	strVal, defined := os.LookupEnv(common.SERVER_PORT_ENV_VAR_NAME)
	if !defined {
		strVal = common.SERVER_DEFAULT_PORT
	}
	intVal, err := strconv.ParseUint(
		strVal, common.BASE_TEN, common.UINT_16_SIZE)
	if err != nil {
		return nil, err
	}
	env.serverPort = strconv.FormatUint(intVal, common.BASE_TEN)

	strVal, defined = os.LookupEnv(common.STATIC_FILES_PATH_ENV_VAR_NAME)
	if !defined {
		strVal = common.STATIC_FILES_DEFAULT_PATH
	}
	env.staticFilesLocalPath = strings.TrimSpace(strVal)

	return &env, nil
}

func printEnvironment(e *runtimeEnv) {
	fmt.Printf("Server port: %s\n", e.serverPort)
	fmt.Printf("\n")
	fmt.Printf("Current working directory: %s\n", e.workingDirectory)
	fmt.Printf("Local path to static files: %s\n", e.staticFilesLocalPath)
}

func bootstrapServer(e *runtimeEnv) {
	router := mux.NewRouter()

	setupRoutes(router, e)
	startServer(router, e)
}

func setupRoutes(router *mux.Router, e *runtimeEnv) {
	router.HandleFunc("/book/{title}/page/{page}",
		func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			title := vars["title"]
			page := vars["page"]

			fmt.Fprintf(w, "Book: %s on page %s\n", title, page)
		},
	)

	staticUrl := "/static/"
	staticServer := http.FileServer(http.Dir(e.staticFilesLocalPath))
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

func startServer(router *mux.Router, e *runtimeEnv) {
	log.Fatal(http.ListenAndServe(":"+e.serverPort, router))
}
