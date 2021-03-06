package exec

import (
	"bytes"
	"fmt"
	"github.com/gofunct/common/io"
	"github.com/gofunct/gogen/context"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

// NewVersionCommand create a new cobra.Command to print the version information.
func NewVersionCommand(io io.IO, cfg context.Build) *cobra.Command {
	return &cobra.Command{
		Use:           "version",
		Short:         "Print the version information",
		Long:          "Print the version information",
		SilenceErrors: true,
		SilenceUsage:  true,
		Run: func(cmd *cobra.Command, _ []string) {
			buf := bytes.NewBufferString(cfg.AppName + " " + cfg.Version)
			buf.WriteString(" (")
			var meta []string
			for _, c := range []string{runtime.Version(), fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)} {
				if c != "" {
					meta = append(meta, c)
				}
			}
			buf.WriteString(strings.Join(meta, " "))
			buf.WriteString(")")
			fmt.Fprintln(io.Out(), buf.String())
		},
	}
}
