package cmdline

import (
	"flag"
	"go-web-examples/internal/common"
	"path/filepath"
	"strings"
)

func Parse() (common.CmdlineArgs, error) {
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
