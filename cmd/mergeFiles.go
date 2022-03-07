/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// mergeFilesCmd represents the mergeFiles command
var mergeFilesCmd = &cobra.Command{
	Use:   "mergeFiles",
	Short: "Merges all js files in the ./src/ directory",
	Long: `Merges all JavaScript (js) files in the ./src/ directory`,
	Run: func(cmd *cobra.Command, args []string) {
		mergedFileExtension := "." + args[0]
		outputFileName := args[1] + mergedFileExtension
		var includes string
		var mergedData string

		filepath.Walk("src/", func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Fatalf(err.Error())
			}
			isImport := true
			fileName := info.Name()
			if strings.Contains(fileName, mergedFileExtension){
				dat, _ := os.ReadFile("src/" + fileName)
				lnScanner := bufio.NewScanner(strings.NewReader(string(dat)))
				for lnScanner.Scan(){
					if strings.Contains(lnScanner.Text(), "// IMPORT-END"){
						isImport = false
					}
					if isImport{
						if !strings.Contains(includes, lnScanner.Text()){
							includes += lnScanner.Text()
							includes += "\n"
						}
					}else{
						if !strings.Contains(lnScanner.Text(), "// IMPORT-END"){
							mergedData += lnScanner.Text()
							mergedData += "\n"
						}
					}
				}
			}
			return nil
		})

		file, _ := os.Create(outputFileName)
		file.Write([]byte(includes))
		file.Write([]byte("\n"))
		file.Write([]byte(mergedData))
		file.Close()
	},
}

func init() {
	rootCmd.AddCommand(mergeFilesCmd)
}
