package cmd

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/ini.v1"
)

func splitStartChars(splitString string, startFromChar int) string {
	tmpStr := splitString
	startIndex := startFromChar
	return tmpStr[startIndex:]
}

func getFileData(filePath string) []byte {
	fileData, _ := os.Open(splitStartChars(filePath, 2))
	fileReader := bufio.NewReader(fileData)
	fileAllContent, _ := ioutil.ReadAll(fileReader)
	return fileAllContent
}


var encodeB64Cmd = &cobra.Command{
	Use:   "encodeB64",
	Short: "Encodes to Base64",
	Long: `Encodes selected file to Base64.
	Supported file formats:
	- .TXT
	- .PNG`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, _ := ini.Load("settings.ini")
		encodeFilePath := cfg.Section("encodeFileParameters").Key("defaultEncodeName").String() + "." + cfg.Section("encodeFileParameters").Key("defaultEncodeExtension").String()
		filePath := args[0]
		fileContent := getFileData(filePath)
		encodedData := base64.StdEncoding.EncodeToString(fileContent)
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
	rootCmd.AddCommand(encodeB64Cmd)
}
