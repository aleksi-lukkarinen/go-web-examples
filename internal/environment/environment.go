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
	err := readEnvFile(args)
	if err != nil {
		return nil, err
	}

	env := common.RuntimeEnv{}

	fmt.Printf("\n")

	err = resolveWorkingDirectory(&env)
	if err != nil {
		return nil, err
	}

	resolveStaticFilesPath(&env)

	fmt.Printf("\n")

	resolveServerIP(&env)

	err = resolveServerPort(&env)
	if err != nil {
		return nil, err
	}

	return &env, nil
}

func readEnvFile(args common.CmdlineArgs) error {
	if len(args.EnvFilePath) > 0 {
		fmt.Printf("Using environment file at: %s\n", args.EnvFilePath)

		err := godotenv.Load(args.EnvFilePath)
		return err
	}

	return nil
}

func resolveWorkingDirectory(env *common.RuntimeEnv) error {
	path, err := os.Getwd()
	if err != nil {
		return err
	}

	env.WorkingDirectory = path
	fmt.Printf("Current working directory: %s\n", env.WorkingDirectory)

	return nil
}

func resolveStaticFilesPath(env *common.RuntimeEnv) {
	strVal, defined := os.LookupEnv(common.STATIC_FILES_PATH_ENV_VAR_NAME)
	if !defined {
		strVal = common.STATIC_FILES_DEFAULT_PATH
	}

	env.StaticFilesLocalPath = strings.TrimSpace(strVal)
	fmt.Printf("Local path to static files: %s\n", env.StaticFilesLocalPath)
}

func resolveServerIP(env *common.RuntimeEnv) {
	strVal, defined := os.LookupEnv(common.SERVER_IP_ENV_VAR_NAME)
	if !defined {
		strVal = common.SERVER_DEFAULT_IP
	}

	env.ServerIP = strings.TrimSpace(strVal)
	fmt.Printf("Server IP: %s\n", env.ServerIP)
}

func resolveServerPort(env *common.RuntimeEnv) error {
	strVal, defined := os.LookupEnv(common.SERVER_PORT_ENV_VAR_NAME)
	if !defined {
		strVal = common.SERVER_DEFAULT_PORT
	}

	intVal, err := strconv.ParseUint(strVal, common.BASE_TEN, common.UINT_16_SIZE)
	if err != nil {
		return err
	}

	env.ServerPort = strconv.FormatUint(intVal, common.BASE_TEN)
	fmt.Printf("Server port: %s\n", env.ServerPort)

	return nil
}
