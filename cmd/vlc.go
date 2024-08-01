package cmd

import (
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

var vlcCmd = &cobra.Command{
	Use: "vlc",
	Short: "Pack file using variable-length code",
	// Run: func(cmd *cobra.Command, args []string),
}

func pack(_ *cobra.Command, args []string){
	filePath := args[0]
	
	r, err := os.Open(filePath)
	if err != nil {
		handleErr(err)
	}
	
	data, err := ioutil.ReadAll(r)
	if err != nil {
		handleErr(err)
	}
	
	// packed := Encode(data)
	packed := ""
	err = ioutil.WriteFile(packedFileName, []byte(packed), 0644)
	if err != nil {
		handleErr(err)
	}
}