package cmd

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	"example.com/pet_proj/archiver_text/lib/compression"
	"example.com/pet_proj/archiver_text/lib/compression/vlc"
	"github.com/spf13/cobra"
)

var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Pack file",
	Run:   pack,
}

const packedExtension = "vlc"

var ErrEmptyPath = errors.New("path to file is not specified")

func pack(cmd *cobra.Command, args []string) {
	var encoder compression.Encoder
	if len(args) == 0 || args[0] == "" {
		handleErr(ErrEmptyPath)
	}

	method := cmd.Flag("method").Value.String()

	switch method {
	case "vlc":
		encoder = vlc.New()
	default:
		cmd.PrintErr("unknown method")
	}

	filePath := args[0]

	r, err := os.Open(filePath)
	if err != nil {
		handleErr(err)
	}

	defer r.Close()

	data, err := io.ReadAll(r)
	if err != nil {
		handleErr(err)
	}

	packed := encoder.Encode(string(data))

	err = os.WriteFile(packedFileName(filePath), packed, 0644)
	if err != nil {
		handleErr(err)
	}

}

func packedFileName(path string) string {
	fileName := filepath.Base(path)               // получаем из пути имя файла
	ext := filepath.Ext(fileName)                 // получаем на основе имени файла его расширение
	baseName := strings.TrimSuffix(fileName, ext) // получаем имя файла без расширения

	return baseName + "." + packedExtension // добавляем к имени файла наше расширение
}

func init() {
	rootCmd.AddCommand(packCmd)

	packCmd.Flags().StringP("method", "m", "", "compression method: vlc")

	if err := packCmd.MarkFlagRequired("method"); err != nil {
		panic(err)
	}
}
