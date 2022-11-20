package myexcel

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/xuri/excelize/v2"
)

type ExcelInfo struct {
	Name   string
	Sheets []*SheetInfo
}

type SheetInfo struct {
	Name     string
	Types    []string
	Varnames []string
	Descs    []string
	Content  [][]string
}

func (this *ExcelInfo) Load(path, name string) error {
	f, err := excelize.OpenFile(path + "/" + name + ".xlsx")
	if err != nil {
		return err
	}
	sheets := f.GetSheetList()
	for _, sheet := range sheets {
		if strings.HasPrefix(sheet, "Sheet") {
			continue
		}
		if len([]byte(sheet)) != len([]rune(sheet)) {
			continue
		}
		sheetInfo := &SheetInfo{}
		sheetInfo.Name = sheet
		rows, err := f.GetRows(sheet)
		if err != nil {
			return err
		}
		if len(rows) < 4 {
			return errors.New("表结构不足4行:" + name + ".xlsz" + " sheet")
		}
		needExport := make(map[int]bool)
		for i, v := range rows[0] {
			if strings.Contains(v, "s") {
				needExport[i] = true
			}
		}
		for i, vname := range rows[1] {
			if !needExport[i] {
				continue
			}
			sheetInfo.Varnames = append(sheetInfo.Varnames, vname)
		}
		if len(sheetInfo.Varnames) != len(needExport) {
			return errors.New("字段名不能为空:" + name + ".xlsz" + " sheet")
		}
		for i, t := range rows[2] {
			if !needExport[i] {
				continue
			}
			sheetInfo.Types = append(sheetInfo.Types, t)
		}
		if len(sheetInfo.Types) != len(needExport) {
			return errors.New("类型不能为空:" + name + ".xlsz" + " sheet")
		}
		for i, desc := range rows[3] {
			if !needExport[i] {
				continue
			}
			sheetInfo.Descs = append(sheetInfo.Descs, desc)
		}
		for i := len(sheetInfo.Descs); i < len(needExport); i++ {
			sheetInfo.Descs = append(sheetInfo.Descs, "")
		}
		for r := 4; r < len(rows); r++ {
			row := rows[r]
			content := []string{}
			for i, cell := range row {
				if !needExport[i] {
					continue
				}
				content = append(content, cell)
			}
			for i := len(content); i < len(needExport); i++ {
				content = append(content, "")
			}
			sheetInfo.Content = append(sheetInfo.Content, content)
		}
		this.Sheets = append(this.Sheets, sheetInfo)
	}
	return nil
}

func (this *ExcelInfo) GenJson(path string) error {
	jsonStr, err := this.ToJson()
	if err != nil {
		return err
	}
	var out bytes.Buffer
	err = json.Indent(&out, []byte(jsonStr), "", "    ")
	filePath := path + "/" + this.Name + ".json"
	fmt.Println(filePath)
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if nil != err {
		return errors.New("open json file failed " + this.Name + ".json")
	}
	defer file.Close()
	file.WriteString(out.String())
	fmt.Println("gen json " + this.Name + ".xlsx success")
	return nil
}

func (this *ExcelInfo) ToJson() (string, error) {
	ret := "{"
	for i, s := range this.Sheets {
		str, err := s.ToJson()
		if err != nil {
			return "", err
		}
		if i != 0 {
			ret += ","
		}
		ret += str
	}
	ret += "}"
	return ret, nil
}

func (this *SheetInfo) ToJson() (string, error) {
	ret := "\"" + this.Name + "\":["

	for i, row := range this.Content {
		rowStr := "{"
		for j, cell := range row {
			cellStr := "\"" + this.Varnames[j] + "\":" + "\"" + cell + "\""
			if j != 0 {
				rowStr += ","
			}
			rowStr += cellStr
		}
		rowStr += "}"
		if i != 0 {
			ret += ","
		}
		ret += rowStr
	}
	ret += "]"
	return ret, nil
}
