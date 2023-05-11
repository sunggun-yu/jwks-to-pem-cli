/*
Copyright © 2023 Sunggun Yu <sunggun.dev@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/sunggun-yu/jwks-to-pem/internal/jwks"
	"github.com/sunggun-yu/jwks-to-pem/internal/utils"
)

// rootCmdFlagFile is variable for --file flag
var rootCmdFlagFile string

var rootCmd = &cobra.Command{
	Use:          "jwks-to-pem",
	Short:        "A simple cli tool to convert jwks to pem",
	Args:         cobra.NoArgs,
	SilenceUsage: true,
	Example: `  # convert jwks to pem from file
  jwks-to-pem -f [jwks-json-file-path]
  
	# convert jwks to pem with text in json file passed into stdin
  cat [private-key-file-path] | jwks-to-pem -f -
  
  # convert jwks to pem from kubernetes jwks
  kubectl get --raw /openid/v1/jwks | jwks-to-pem -f -`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var src []byte
		var err error
		if rootCmdFlagFile == "-" {
			// read private key from stdin if flag value is `-`
			src = []byte(utils.ReadInOrStdin(cmd))
		} else {
			src, err = os.ReadFile(rootCmdFlagFile)
			if err != nil {
				return err
			}
		}

		pems, err := jwks.Convert(cmd.Context(), src)
		if err != nil {
			return err
		}
		fmt.Print(strings.Join(pems[:], ","))
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// Set the version of cmd
func SetVersion(version string) {
	rootCmd.Version = version
}

func init() {
	rootCmd.Flags().StringVarP(&rootCmdFlagFile, "file", "f", "", "The jwks json file")
	rootCmd.MarkFlagRequired("file")
}
