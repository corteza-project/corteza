package cli

import (
	"os"
	"sort"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

// LoadEnv loads env-variables (KEY=VAL) from provided files and paths (by searching for .env file in the given path)
//
// Please not that loaded values DO NOT OVERRIDE the existing environmental variables
func LoadEnv(pp ...string) error {
	// preparse the input and try to figure out if .env should be appended
	for i, p := range pp {
		if s, err := os.Stat(p); err != nil {
			return err
		} else if s.IsDir() {
			pp[i] = p + "/.env"
		}
	}

	return godotenv.Load(pp...)
}

// EnvCommand outputs all loaded env variables
func EnvCommand() *cobra.Command {
	return &cobra.Command{
		Use:  "env",
		Long: "Outputs list (sorted by key) of of all environmental variables. Can be used for diagnosis and debugging",
		Run: func(cmd *cobra.Command, args []string) {
			kv := os.Environ()
			sort.Strings(kv)
			for _, l := range kv {
				cmd.Println(l)
			}
		},
	}
}
