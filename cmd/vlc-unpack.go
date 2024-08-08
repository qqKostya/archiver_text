package cmd

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"example.com/pet_proj/archiver_text/lib/vlc"
	"github.com/spf13/cobra"
)

var vlcUnpackCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Unpack file using variable-length code",
	Run:   pack,
}

const unpackedExtension = "txt"

func unpack(_ *cobra.Command, args []string) {
	if len(args) == 0 || args[0] == "" {
		handleErr(ErrEmptyPath)
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

	packed := vlc.Decode(data)

	err = os.WriteFile(unpackedFileName(filePath), []byte(packed), 0644)
	if err != nil {
		handleErr(err)
	}
}

func unpackedFileName(path string) string {
	fileName := filepath.Base(path)               // получаем из пути имя файла
	ext := filepath.Ext(fileName)                 // получаем на основе имени файла его расширение
	baseName := strings.TrimSuffix(fileName, ext) // получаем имя файла без расширения

	return baseName + "." + packedExtension // добавляем к имени файла наше расширение
}

func init() {
	unpackCmd.AddCommand(vlcUnpackCmd)
}
