package utils

import (
	"bytes"
	"io"
	"strings"

	"github.com/spf13/cobra"
)

// ReadInOrStdin reads string from command ReadInOrStdin
func ReadInOrStdin(cmd *cobra.Command) string {
	var inputReader io.Reader = cmd.InOrStdin()
	buf := new(bytes.Buffer)
	buf.ReadFrom(inputReader)
	// NOTE: it keep waiting if there is no stdin. same like kubectl apply -f with no input
	// TODO: set timeout to avoid hang?
	return strings.TrimSpace(buf.String())
}
