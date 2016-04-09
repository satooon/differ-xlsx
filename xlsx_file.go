package main

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

type xlsxFile struct {
	xlsx.File
	FilePath string
}

func NewXlsxFile(f *xlsx.File, FilePath string) *xlsxFile {
	return &xlsxFile{
		File:     *f,
		FilePath: FilePath,
	}
}

func (xlsx *xlsxFile) OutPutSheet(sheetName string) [][]string {
	output := [][]string{}
	sheet, ok := xlsx.Sheet[sheetName]
	if !ok {
		return output
	}
	for _, row := range sheet.Rows {
		if row == nil {
			output = append(output, []string{})
			continue
		}
		r := []string{}
		for _, cell := range row.Cells {
			str, err := cell.String()
			if err != nil {
				panic(fmt.Sprintf("%v cell to string error. err:%v", xlsx.FilePath, err))
			}
			r = append(r, str)
		}
		output = append(output, r)
	}
	return output
}
