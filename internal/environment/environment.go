package environment

import (
	"fmt"
	"go-web-examples/internal/common"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

func Init(args common.CmdlineArgs) (*common.RuntimeEnv, error) {
	if len(args.EnvFilePath) > 0 {
		fmt.Printf("Using environment file at: %s\n", args.EnvFilePath)

		err := godotenv.Load(args.EnvFilePath)
		if err != nil {
			return nil, err
		}
	}

	env := common.RuntimeEnv{}

	path, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	env.WorkingDirectory = path

	strVal, defined := os.LookupEnv(common.SERVER_PORT_ENV_VAR_NAME)
	if !defined {
		strVal = common.SERVER_DEFAULT_PORT
	}
	intVal, err := strconv.ParseUint(
		strVal, common.BASE_TEN, common.UINT_16_SIZE)
	if err != nil {
		return nil, err
	}
	env.ServerPort = strconv.FormatUint(intVal, common.BASE_TEN)

	strVal, defined = os.LookupEnv(common.SERVER_IP_ENV_VAR_NAME)
	if !defined {
		strVal = common.SERVER_DEFAULT_IP
	}
	env.ServerIP = strings.TrimSpace(strVal)

	strVal, defined = os.LookupEnv(common.STATIC_FILES_PATH_ENV_VAR_NAME)
	if !defined {
		strVal = common.STATIC_FILES_DEFAULT_PATH
	}
	env.StaticFilesLocalPath = strings.TrimSpace(strVal)

	return &env, nil
}

func printEnvironment(e *common.RuntimeEnv) {
	fmt.Printf("\n")
	fmt.Printf("Current working directory: %s\n", e.WorkingDirectory)
	fmt.Printf("Local path to static files: %s\n", e.StaticFilesLocalPath)
	fmt.Printf("\n")
	fmt.Printf("Server IP: %s\n", e.ServerIP)
	fmt.Printf("Server port: %s\n", e.ServerPort)
}
