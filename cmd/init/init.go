package init

import (
	"fmt"
	"os"

	"github.com/desmos-labs/juno/v2/types/config"

	"github.com/spf13/cobra"
)

const (
	flagReplace = "replace"
)

// InitCmd returns the command that should be run in order to properly initialize BDJuno
func InitCmd(cfg *Config) *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initializes the configuration files",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Create the config path if not present
			if _, err := os.Stat(config.HomePath); os.IsNotExist(err) {
				err = os.MkdirAll(config.HomePath, os.ModePerm)
				if err != nil {
					return err
				}
			}

			replace, err := cmd.Flags().GetBool(flagReplace)
			if err != nil {
				return err
			}

			// Get the config file
			configFilePath := config.GetConfigFilePath()
			file, _ := os.Stat(configFilePath)

			// Check if the file exists and replace is false
			if file != nil && !replace {
				return fmt.Errorf(
					"configuration file already present at %s. If you wish to overwrite it, use the --%s flag",
					configFilePath, flagReplace)
			}

			// Get the config from the flags
			yamlCfg := cfg.GetConfigCreator()(cmd)
			return config.Write(yamlCfg, configFilePath)
		},
	}
}
