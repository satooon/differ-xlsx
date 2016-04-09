package main

import (
	"os"

	"fmt"

	"encoding/json"

	"github.com/codegangsta/cli"
	"github.com/tealeg/xlsx"
)

func main() {
	app := NewApp()
	app.Run(os.Args)
}

func NewApp() *cli.App {
	app := cli.NewApp()
	app.Name = "differ-xlsx"
	app.Usage = "differ-xlsx"
	app.Author = "satooon"
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "verbose, vv",
			Usage: "verbose mode",
		},
	}
	app.Action = Action
	return app
}

func Action(ctx *cli.Context) {
	verbose := ctx.Bool("verbose")
	debugLog = debugLogStruct{verbose: verbose}
	args := ctx.Args()

	if len(args) < 2 {
		panic("please select 2 files")
	}

	srcFilePath := args[0]
	destFilePath := args[1]
	debugLog.Println("srcFilePath:", srcFilePath, "destFilePath:", destFilePath)

	// open src file
	srcXlFile, err := xlsx.OpenFile(srcFilePath)
	if err != nil {
		panic(fmt.Sprintf("%s open error. err:%v", srcFilePath, err))
	}
	srcXlsxFile := NewXlsxFile(srcXlFile, srcFilePath)
	//outputFistFile := srcXlsxFile.OutPut()
	//debugLog.Println("outputFistFile", outputFistFile)
	// open dest file
	destXlFile, err := xlsx.OpenFile(destFilePath)
	if err != nil {
		panic(fmt.Sprintf("%s open error. err:%v", destFilePath, err))
	}
	destXlsxFile := NewXlsxFile(destXlFile, destFilePath)
	//destXutputFile := destXlsxFile.OutPut()
	//debugLog.Println("destXutputFile", destXutputFile)

	// check diff
	output := map[string]string{}

	for _, sheet := range destXlsxFile.Sheets {
		if _, ok := srcXlsxFile.Sheet[sheet.Name]; !ok {
			output[sheet.Name] = "new sheet"
			continue
		}
		srcJson, err := json.Marshal(srcXlsxFile.OutPutSheet(sheet.Name))
		if err != nil {
			panic(fmt.Sprintf("%v json marshal error. %v", srcXlsxFile.FilePath, err))
		}
		destJson, err := json.Marshal(destXlsxFile.OutPutSheet(sheet.Name))
		if err != nil {
			panic(fmt.Sprintf("%v json marshal error. %v", destXlsxFile.FilePath, err))
		}
		if string(srcJson) != string(destJson) {
			output[sheet.Name] = "modify"
		}
	}

	for _, sheet := range srcXlsxFile.Sheets {
		if _, ok := destXlsxFile.Sheet[sheet.Name]; !ok {
			output[sheet.Name] = "delete sheet"
			continue
		}
		if _, ok := output[sheet.Name]; ok {
			continue
		}
		//srcJson, err := json.Marshal(srcXlsxFile.OutPutSheet(sheet.Name))
		//if err != nil {
		//	panic(fmt.Sprintf("%v json marshal error. %v", srcXlsxFile.FilePath, err))
		//}
		//destJson, err := json.Marshal(destXlsxFile.OutPutSheet(sheet.Name))
		//if err != nil {
		//	panic(fmt.Sprintf("%v json marshal error. %v", destXlsxFile.FilePath, err))
		//}
		//if string(srcJson) != string(destJson) {
		//	output[sheet.Name] = "modify"
		//}
	}

	// output
	for sheetName, msg := range output {
		fmt.Printf("%s : %s\n", sheetName, msg)
	}
}
