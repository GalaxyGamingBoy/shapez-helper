package cmd

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/ini.v1"
)

var encodeB64LineCmd = &cobra.Command{
	Use:   "encodeB64Line",
	Short: "Encodes to Base64",
	Long: `Encodes selected line to Base64`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, _ := ini.Load("settings.ini")
		encodeFilePath := cfg.Section("encodeFileParameters").Key("defaultEncodeName").String() + "." + cfg.Section("encodeFileParameters").Key("defaultEncodeExtension").String()
		encodedData := base64.StdEncoding.EncodeToString([]byte(args[0]))
		if(cfg.Section("encodeData").Key("writeEncodeFile").String() == "true"){
			os.Remove(cfg.Section("other").Key("latestEncodeFile").String())
			file, _ := os.Create(encodeFilePath)
			file.Write([]byte(encodedData))
			file.Close()
			cfg.Section("other").Key("latestEncodeFile").SetValue(encodeFilePath)
			cfg.SaveTo("settings.ini")
		}
		if(cfg.Section("encodeData").Key("printEncodeData").String() == "true"){
			fmt.Println(encodedData)
		}
	},
}

func init() {
	rootCmd.AddCommand(encodeB64LineCmd)
}
