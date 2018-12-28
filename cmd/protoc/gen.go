// Copyright © 2018 Coleman Word <coleman.word@gofunct.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package protoc

import (
	"fmt"
	"go.uber.org/zap"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)


func init() {
	ProtocCmd.AddCommand(genCmd)
}

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate protobug files",
	Run: func(cmd *cobra.Command, args []string) {
		if err := filepath.Walk(protoDir, walk()); err != nil {
			logger.Fatal("failed to generate protobufs", zap.Error(err))
		}
	},
}

func walk() filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {

		// skip vendor protoDirectory
		if info.IsDir() && info.Name() == "vendor" {
			return filepath.SkipDir
		}
		// find all protobuf files
		if filepath.Ext(path) == ".proto" {
			// args
			args := []string{
				"-I=.",
				fmt.Sprintf("-I=%s", filepath.Join(goPath, "src")),
				fmt.Sprintf("-I=%s", filepath.Join(goPath, "src", "github.com", "gogo", "protobuf", "protobuf")),
				fmt.Sprintf("--proto_path=%s", filepath.Join(goPath, "src", "github.com")),
				"--gogofaster_out=Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types," +
					"Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types:.",
				path,
			}
			cmd := exec.Command("protoc", args...)
			err = cmd.Run()
			if err != nil {
				return err
			}
		}
		return nil
	}
}