package main

import (
	"os"

	"github.com/drone-plugins/drone-plugin-lib/errors"
	"github.com/drone-plugins/drone-plugin-lib/urfave"
	"github.com/joho/godotenv"
	"github.com/orangedeng/drone-tools/drone-checksum/plugin"
	"github.com/urfave/cli/v2"
)

var Version = "dev"
var settings = &plugin.Settings{}

func main() {

	if _, err := os.Stat("/run/drone/env"); err == nil {
		_ = godotenv.Overload("/run/drone/env")
	}

	app := &cli.App{
		Name:    "drone-checksum",
		Usage:   "drone plugin only do checksum",
		Version: Version,
		Flags:   append(settingsFlags(settings), urfave.Flags()...),
		Action:  run(settings),
	}
	if err := app.Run(os.Args); err != nil {
		errors.HandleExit(err)
	}
}

func run(settings *plugin.Settings) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		urfave.LoggingFromContext(ctx)
		plugin := plugin.New(
			settings,
			urfave.PipelineFromContext(ctx),
			urfave.NetworkFromContext(ctx),
		)
		if err := plugin.Validate(); err != nil {
			if e, ok := err.(errors.ExitCoder); ok {
				return e
			}

			return errors.ExitMessagef("validation failed: %w", err)
		}

		if err := plugin.Execute(); err != nil {
			if e, ok := err.(errors.ExitCoder); ok {
				return e
			}

			return errors.ExitMessagef("execution failed: %w", err)
		}

		return nil
	}
}
func settingsFlags(settings *plugin.Settings) []cli.Flag {
	return []cli.Flag{
		&cli.StringSliceFlag{
			Name:        "files",
			Usage:       "list of files to do checksum",
			EnvVars:     []string{"PLUGIN_FILES", "GITHUB_RELEASE_FILES"},
			Destination: &settings.Files,
		},
		&cli.StringSliceFlag{
			Name:        "checksum",
			Usage:       "generate specific checksums",
			EnvVars:     []string{"PLUGIN_CHECKSUM", "GITHUB_RELEASE_CHECKSUM"},
			Destination: &settings.Checksum,
		},
		&cli.StringFlag{
			Name:        "checksum-file",
			Usage:       "name used for checksum file. \"CHECKSUM\" is replaced with the chosen method",
			EnvVars:     []string{"PLUGIN_CHECKSUM_FILE"},
			Value:       "CHECKSUMsum.txt",
			Destination: &settings.ChecksumFile,
		},
		&cli.BoolFlag{
			Name:        "checksum-flatten",
			Usage:       "include only the basename of the file in the checksum file",
			EnvVars:     []string{"PLUGIN_CHECKSUM_FLATTEN"},
			Destination: &settings.ChecksumFlatten,
		},
	}
}
