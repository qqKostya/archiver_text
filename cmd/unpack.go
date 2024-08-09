package cmd

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"example.com/pet_proj/archiver_text/lib/compression"
	"example.com/pet_proj/archiver_text/lib/compression/vlc"

	"github.com/spf13/cobra"
)

var unpackCmd = &cobra.Command{
	Use:   "unpack",
	Short: "Unpack file",
	Run:   unpack,
}

const unpackedExtension = "txt"

func unpack(cmd *cobra.Command, args []string) {
	var dencoder compression.Decoder

	if len(args) == 0 || args[0] == "" {
		handleErr(ErrEmptyPath)
	}

	method := cmd.Flag("method").Value.String()

	switch method {
	case "vlc":
		dencoder = vlc.New()
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

	packed := dencoder.Decode(data)

	err = os.WriteFile(unpackedFileName(filePath), []byte(packed), 0644)
	if err != nil {
		handleErr(err)
	}
}

func unpackedFileName(path string) string {
	fileName := filepath.Base(path)               // получаем из пути имя файла
	ext := filepath.Ext(fileName)                 // получаем на основе имени файла его расширение
	baseName := strings.TrimSuffix(fileName, ext) // получаем имя файла без расширения

	return baseName + "." + unpackedExtension // добавляем к имени файла наше расширение
}

func init() {
	rootCmd.AddCommand(unpackCmd)

	unpackCmd.Flags().StringP("method", "m", "", "decompression method: vlc")

	if err := unpackCmd.MarkFlagRequired("method"); err != nil {
		panic(err)
	}
}
