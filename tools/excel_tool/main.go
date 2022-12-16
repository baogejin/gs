package main

import (
	"fmt"
	"gs/define"
	"gs/tools/excel_tool/myexcel"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func main() {
	rootPath := os.Getenv(define.EnvName)
	if rootPath == "" {
		panic("gs环境变量没有设置，可以先运行bin下的set_env.bat")
	}
	excelPath := rootPath + "/data/excel"
	jsonPath := rootPath + "/data/json"
	codePath := rootPath + "/data/gencode"
	fileList, err := getExcelList(excelPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	//读取excel
	excels := loadExcels(excelPath, fileList)
	//生成json
	genJson(jsonPath, excels)
	//生成go代码
	genGoCode(codePath, excels)
	//生成ts代码
	genTsCode(codePath, excels)

}

func getExcelList(path string) ([]string, error) {
	ret := []string{}
	fileInfoList, err := ioutil.ReadDir(path)
	if err != nil {
		return ret, err
	}
	for _, fileInfo := range fileInfoList {
		name := fileInfo.Name()
		if strings.HasPrefix(name, "~") {
			continue
		}
		strSlc := strings.Split(name, ".")
		if len(strSlc) != 2 || strSlc[1] != "xlsx" {
			continue
		}
		if len([]byte(name)) != len([]rune(name)) {
			fmt.Println("文件名请不要用英文数字以外的字符:", name, "该文件已被忽略")
			continue
		}
		ret = append(ret, strSlc[0])
	}
	return ret, nil
}

func loadExcels(path string, files []string) []*myexcel.ExcelInfo {
	ret := []*myexcel.ExcelInfo{}
	for i, file := range files {
		excel := &myexcel.ExcelInfo{}
		excel.Name = file
		err := excel.Load(path, file)
		if err != nil {
			panic("load " + file + ".xlsx failed," + err.Error())
		}
		ret = append(ret, excel)
		fmt.Printf("[%d/%d]load %s.xlsx success\n", i+1, len(files), file)
	}
	return ret
}

func genJson(path string, excels []*myexcel.ExcelInfo) {
	for i, excel := range excels {
		if err := excel.GenJson(path); err != nil {
			panic("gen json failed:" + excel.Name + ".xlsx " + err.Error())
		}
		fmt.Printf("[%d/%d]gen json %s.xlsx success\n", i+1, len(excels), excel.Name)
	}
}

func genGoCode(path string, excels []*myexcel.ExcelInfo) {
	for i, excel := range excels {
		if excel.Name == "Global" {
			if err := excel.GenGlobalKey(path); err != nil {
				panic("gen global key failed " + err.Error())
			}
			fmt.Printf("[%d/%d]gen global key %s.xlsx success\n", i+1, len(excels), excel.Name)
			continue
		}
		if err := excel.GenCode(path); err != nil {
			panic("gen code failed:" + excel.Name + ".xlsx " + err.Error())
		}
		fmt.Printf("[%d/%d]gen code %s.xlsx success\n", i+1, len(excels), excel.Name)
	}
	cmd := exec.Command("gofmt", "-w", path)
	err := cmd.Run()
	if nil != err {
		fmt.Println(err)
	}
}

func genTsCode(path string, excels []*myexcel.ExcelInfo) {
	for i, excel := range excels {
		if excel.Name == "Global" {
			continue
		}
		if err := excel.GenTsCode(path); err != nil {
			panic("gen code failed:" + excel.Name + ".xlsx " + err.Error())
		}
		fmt.Printf("[%d/%d]gen code %s.xlsx success\n", i+1, len(excels), excel.Name)
	}
}
