package cmd

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/ini.v1"
)

func checkErr(errorVal error){
	if(errorVal != nil){
		panic("Error: " + errorVal.Error())
	}
}

var decodeB64Cmd = &cobra.Command{
	Use:   "decodeB64",
	Short: "Decrypts line to Base64",
	Long: `Decodes selected line from Base 64 to Data`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, _ := ini.Load("settings.ini")
		decodeFilePath := cfg.Section("decodeFile").Key("defaultDecodeName").String() + "." +  cfg.Section("decodeFile").Key("defaultDecodeExtension").String()
		baseData := getFileData(args[0])
		decodedData, err := base64.StdEncoding.DecodeString(string(baseData))
		checkErr(err)
		if(cfg.Section("").Key("deleteLatestDecodeFile").String() == "true"){
			os.Remove(cfg.Section("other").Key("latestDecodeFile").String())
		}
		if(cfg.Section("").Key("writeDecodeFile").String() == "true"){
			file, _ := os.Create(decodeFilePath)
			file.Write([]byte(decodedData))
			file.Close()
			cfg.Section("other").Key("latestDecodeFile").SetValue(decodeFilePath)
			cfg.SaveTo("settings.ini")
		}
		fmt.Print(string(decodedData))
	},
}

func init() {
	rootCmd.AddCommand(decodeB64Cmd)
}
