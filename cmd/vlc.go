package cmd

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	"example.com/pet_proj/archiver_text/lib/vlc"
	"github.com/spf13/cobra"
)

var vlcCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Pack file using variable-length code",
	Run:   pack,
}

const packedExtension = "vlc"

var ErrEmptyPath = errors.New("path to file is not specified")

func pack(_ *cobra.Command, args []string) {
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

	packed := vlc.Encode(string(data))

	err = os.WriteFile(packedFileName(filePath), []byte(packed), 0644)
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
	packCmd.AddCommand(vlcCmd)
}
