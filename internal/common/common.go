package common

const BASE_TEN = 10
const UINT_16_SIZE = 16

const URL_PART_STATIC = "/static/"

const STATIC_FILES_PATH_ENV_VAR_NAME = "TESTSERVER_STATICFILES_PATH"
const STATIC_FILES_DEFAULT_PATH = "./web/static"

const SERVER_IP_ENV_VAR_NAME = "TESTSERVER_IP"
const SERVER_DEFAULT_IP = "127.0.0.1"

const SERVER_PORT_ENV_VAR_NAME = "TESTSERVER_PORT"
const SERVER_DEFAULT_PORT = "8000"

type CmdlineArgs struct {
	EnvFilePath string
}

type RuntimeEnv struct {
	WorkingDirectory     string
	StaticFilesLocalPath string
	ServerIP             string
	ServerPort           string
}
