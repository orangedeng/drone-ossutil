package plugin

import (
	"fmt"
	"path/filepath"

	"github.com/drone-plugins/drone-plugin-lib/drone"
	"github.com/urfave/cli/v2"
)

// Settings for the plugin.
type Settings struct {
	Files           cli.StringSlice
	Checksum        cli.StringSlice
	ChecksumFile    string
	ChecksumFlatten bool

	uploads []string
}

type Plugin struct {
	settings *Settings
	pipeline drone.Pipeline
	network  drone.Network
}

func New(settings *Settings, pipeline drone.Pipeline, network drone.Network) drone.Plugin {
	return &Plugin{
		settings: settings,
		pipeline: pipeline,
		network:  network,
	}
}

func (p *Plugin) Validate() error {
	return nil
}

func (p *Plugin) Execute() error {
	var err error
	files := p.settings.Files.Value()
	for _, glob := range files {
		globed, err := filepath.Glob(glob)

		if err != nil {
			return fmt.Errorf("failed to glob %s: %w", glob, err)
		}

		if globed != nil {
			p.settings.uploads = append(p.settings.uploads, globed...)
		}
	}

	if len(files) > 0 && len(p.settings.uploads) < 1 {
		return fmt.Errorf("failed to find any file to release")
	}

	checksum := p.settings.Checksum.Value()
	if len(checksum) > 0 {
		p.settings.uploads, err = writeChecksums(p.settings.uploads, checksum, p.settings.ChecksumFile, p.settings.ChecksumFlatten)

		if err != nil {
			return fmt.Errorf("failed to write checksums: %w", err)
		}
	}

	return nil
}
