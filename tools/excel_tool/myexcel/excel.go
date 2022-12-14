package myexcel

import (
	"errors"
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
	Types    []*TypeInfo
	Varnames []string
	Descs    []string
	Content  [][]string
}

func (this *ExcelInfo) Load(path, name string, tag string) error {
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
			return errors.New("表结构不足4行:" + name + ".xlsx" + " sheet:" + sheet)
		}
		needExport := make(map[int]bool)
		for i, v := range rows[0] {
			if strings.Contains(v, tag) {
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
			return errors.New("字段名不能为空:" + name + ".xlsx" + " sheet:" + sheet)
		}
		for i, t := range rows[2] {
			if !needExport[i] {
				continue
			}
			typeInfo, err := getTypeInfoByStr(strings.ToLower(t))
			if err != nil {
				return err
			}
			typeInfo.FixType()
			sheetInfo.Types = append(sheetInfo.Types, typeInfo)
		}
		if len(sheetInfo.Types) != len(needExport) {
			return errors.New("类型不能为空:" + name + ".xlsx" + " sheet:" + sheet)
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
	filePath := path + "/" + this.Name + ".json"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if nil != err {
		return errors.New("open json file failed " + this.Name + ".json")
	}
	defer file.Close()
	file.WriteString(jsonStr)
	return nil
}

func (this *ExcelInfo) ToJson() (string, error) {
	ret := "{"
	for i, s := range this.Sheets {
		str, err := s.ToJson()
		if err != nil {
			return "", err
		}
		if str == "" {
			continue
		}
		if i != 0 {
			ret += ","
		}
		ret += "\n"
		ret += str
	}
	ret += "\n}"
	return ret, nil
}

func (this *SheetInfo) ToJson() (string, error) {
	if len(this.Varnames) == 0 {
		return "", nil
	}
	ret := "    \"" + this.Name + "\":["
	repeatCheck := make(map[string]bool)
	needCheck := false
	if this.Name == "Global" {
		needCheck = true
	}
	if this.Varnames[0] == "ID" {
		if this.Types[0].CType != CellTypeSimple || this.Types[0].ValueType1 != "int32" {
			return "", errors.New("ID type err in sheet " + this.Name)
		}
		needCheck = true
	}
	for i, row := range this.Content {
		rowStr := "{"
		for j, cell := range row {
			cellStr := "\"" + this.Varnames[j] + "\":"
			vStr, err := this.Types[j].ParseToJson(cell)
			if err != nil {
				return "", err
			}
			if needCheck && j == 0 {
				if repeatCheck[vStr] {
					return "", errors.New(this.Varnames[0] + " repeat in sheet " + this.Name)
				}
				repeatCheck[vStr] = true
			}
			cellStr += vStr
			if j != 0 {
				rowStr += ","
			}
			rowStr += cellStr
		}
		rowStr += "}"
		if i != 0 {
			ret += ","
		}
		ret += "\n        " + rowStr
	}
	ret += "\n    ]"
	return ret, nil
}

func (this *ExcelInfo) GenCode(path string) error {
	ret := ""
	ret += "package gencode\n\n"
	ret += "import (\n"
	ret += "	\"encoding/json\"\n"
	ret += "	\"gs/define\"\n"
	ret += "	\"gs/lib/mylog\"\n"
	ret += "	\"io/ioutil\"\n"
	ret += "	\"os\"\n"
	ret += "	\"sync\"\n"
	ret += ")\n\n"
	ret += "type " + this.Name + "Cfg struct {\n"
	needMap := make(map[string]bool)
	for _, s := range this.Sheets {
		if len(s.Varnames) == 0 {
			continue
		}
		ret += "	" + s.Name + "Slc []*" + s.Name + "Info `json:\"" + s.Name + "\"`\n"
		if s.Varnames[0] == "ID" && s.Types[0].CType == CellTypeSimple && s.Types[0].ValueType1 == "int32" {
			needMap[s.Name] = true
			ret += "	" + s.Name + "Map map[int32]*" + s.Name + "Info\n"
		}
	}
	ret += "}\n\n"
	for _, s := range this.Sheets {
		if len(s.Varnames) == 0 {
			continue
		}
		ret += "type " + s.Name + "Info struct {\n"
		for i := range s.Varnames {
			if s.Types[i].CType == CellTypeSimple {
				ret += "	" + s.Varnames[i] + "   " + s.Types[i].ValueType1
			} else if s.Types[i].CType == CellTypeSlc {
				ret += "	" + s.Varnames[i] + "   []" + s.Types[i].ValueType1
			} else if s.Types[i].CType == CellTypeDoubleSlc {
				ret += "	" + s.Varnames[i] + "   [][]" + s.Types[i].ValueType1
			} else if s.Types[i].CType == CellTypeMap {
				ret += "	" + s.Varnames[i] + "   map[" + s.Types[i].ValueType1 + "]" + s.Types[i].ValueType2
			}
			ret += "// " + strings.Replace(s.Descs[i], "\n", " ", -1) + "\n"
		}
		ret += "}\n\n"
	}
	ret += "var " + strings.ToLower(this.Name) + "Cfg *" + this.Name + "Cfg\n"
	ret += "var " + strings.ToLower(this.Name) + "Once sync.Once\n\n"
	ret += "func Get" + this.Name + "Cfg() *" + this.Name + "Cfg {\n"
	ret += "	" + strings.ToLower(this.Name) + "Once.Do(func() {\n"
	ret += "		" + strings.ToLower(this.Name) + "Cfg = new(" + this.Name + "Cfg)\n"
	ret += "		" + strings.ToLower(this.Name) + "Cfg.init()\n"
	ret += "	})\n"
	ret += "	return " + strings.ToLower(this.Name) + "Cfg\n"
	ret += "}\n\n"
	ret += "func (this *" + this.Name + "Cfg) init() {\n"
	for _, s := range this.Sheets {
		if len(s.Varnames) == 0 {
			continue
		}
		if needMap[s.Name] {
			ret += "this." + s.Name + "Map = make(map[int32]*" + s.Name + "Info)\n"
		}
	}
	ret += "	rootPath := os.Getenv(define.EnvName)\n"
	ret += "	filePtr, err := os.Open(rootPath + \"/data/json/" + this.Name + ".json\")\n"
	ret += "	if err != nil {\n"
	ret += "		mylog.Error(\"load " + this.Name + "Cfg failed\", err)\n"
	ret += "		return\n"
	ret += "	}\n"
	ret += "	defer filePtr.Close()\n"
	ret += "	data, err := ioutil.ReadAll(filePtr)\n"
	ret += "	if err != nil {\n"
	ret += "		mylog.Error(\"load " + this.Name + "Cfg failed\", err)\n"
	ret += "		return\n"
	ret += "	}\n"
	ret += "	err = json.Unmarshal(data, this)\n"
	ret += "	if err != nil {\n"
	ret += "		mylog.Error(\"load " + this.Name + "Cfg failed\", err)\n"
	ret += "		return\n"
	ret += "	}\n"
	for _, s := range this.Sheets {
		if len(s.Varnames) == 0 {
			continue
		}
		if needMap[s.Name] {
			ret += "	for _, v := range this." + s.Name + "Slc {\n"
			ret += "		this." + s.Name + "Map[v.ID] = v\n"
			ret += "	}\n"
		}
	}
	ret += "}\n\n"
	for _, s := range this.Sheets {
		if len(s.Varnames) == 0 {
			continue
		}
		if needMap[s.Name] {
			ret += "func (this *" + this.Name + "Cfg) Get" + s.Name + "ById(id int32) (*" + s.Name + "Info, bool) {\n"
			ret += "	if ret, ok := this." + s.Name + "Map[id]; ok {\n"
			ret += "		return ret, ok\n"
			ret += "	}\n"
			ret += "	return nil, false\n"
			ret += "}\n\n"
		}
	}
	filePath := path + "/" + this.Name + "Cfg.go"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if nil != err {
		return errors.New("open code file failed " + this.Name + "Cfg.go")
	}
	defer file.Close()
	file.WriteString(ret)
	return nil
}

func (this *ExcelInfo) GenGlobalKey(path string) error {
	if this.Name != "Global" {
		return errors.New("this is not Global.xlsx")
	}
	if len(this.Sheets) != 1 || this.Sheets[0].Name != "Global" {
		return errors.New("Global sheet error")
	}
	if this.Sheets[0].Types[0].CType != CellTypeSimple || this.Sheets[0].Types[0].ValueType1 != "string" {
		return errors.New("Global key type error")
	}
	ret := ""
	ret += "package gencode\n\n"
	ret += "const (\n"
	for _, v := range this.Sheets[0].Content {
		ret += v[0] + "=\"" + v[0] + "\"\n"
	}
	ret += ")\n\n"
	filePath := path + "/GlobalKey.go"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if nil != err {
		return errors.New("open code file failed GlobalKey.go")
	}
	defer file.Close()
	file.WriteString(ret)
	return nil
}

func (this *ExcelInfo) GenTsGlobalKey(path string) error {
	if this.Name != "Global" {
		return errors.New("this is not Global.xlsx")
	}
	if len(this.Sheets) != 1 || this.Sheets[0].Name != "Global" {
		return errors.New("Global sheet error")
	}
	if this.Sheets[0].Types[0].CType != CellTypeSimple || this.Sheets[0].Types[0].ValueType1 != "string" {
		return errors.New("Global key type error")
	}
	ret := ""
	ret += "export class GlobalKey {\n"
	for _, v := range this.Sheets[0].Content {
		ret += "    public static " + v[0] + ": string = \"" + v[0] + "\"\n"
	}
	ret += "}"
	filePath := path + "/GlobalKey.ts"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if nil != err {
		return errors.New("open code file failed GlobalKey.ts")
	}
	defer file.Close()
	file.WriteString(ret)
	return nil
}

func (this *ExcelInfo) GenTsCode(path string) error {
	ret := ""
	ret += "import Res from \"../../common/util/Res\"\n\n"
	needMap := make(map[string]bool)
	for _, s := range this.Sheets {
		if len(s.Varnames) == 0 {
			continue
		}
		if s.Varnames[0] == "ID" && s.Types[0].CType == CellTypeSimple && s.Types[0].ValueType1 == "int32" {
			needMap[s.Name] = true
		}
		ret += "export class " + s.Name + "Info {\n"
		for i := range s.Varnames {
			if s.Types[i].CType == CellTypeSimple {
				ret += "    public " + s.Varnames[i] + ": " + getTsType(s.Types[i].ValueType1)
			} else if s.Types[i].CType == CellTypeSlc {
				ret += "    public " + s.Varnames[i] + ": Array<" + getTsType(s.Types[i].ValueType1) + ">"
			} else if s.Types[i].CType == CellTypeDoubleSlc {
				ret += "    public " + s.Varnames[i] + ": Array<Array<" + getTsType(s.Types[i].ValueType1) + ">>"
			} else if s.Types[i].CType == CellTypeMap {
				ret += "    public " + s.Varnames[i] + ": Map<" + getTsType(s.Types[i].ValueType1) + "," + getTsType(s.Types[i].ValueType2) + ">"
			}
			ret += "// " + strings.Replace(s.Descs[i], "\n", " ", -1) + "\n"
		}
		ret += "}\n\n"
	}
	ret += "export class " + this.Name + "Config {\n"
	ret += "    private static instance: " + this.Name + "Config\n"
	ret += "    public static Get(): " + this.Name + "Config {\n"
	ret += "        if (" + this.Name + "Config.instance == null) {\n"
	ret += "            " + this.Name + "Config.instance = new " + this.Name + "Config()\n"
	ret += "            " + this.Name + "Config.instance.init()\n"
	ret += "        }\n"
	ret += "        return " + this.Name + "Config.instance\n"
	ret += "    }\n\n"
	for _, s := range this.Sheets {
		if len(s.Varnames) == 0 {
			continue
		}
		ret += "    public " + s.Name + "Slc: Array<" + s.Name + "Info>\n"
		if needMap[s.Name] {
			ret += "    public " + s.Name + "Map: Map<number, " + s.Name + "Info>\n"
		}
		ret += "\n"
	}
	ret += "    private init(): void {\n"
	for _, s := range this.Sheets {
		if len(s.Varnames) == 0 {
			continue
		}
		ret += "        this." + s.Name + "Slc = new Array<" + s.Name + "Info>()\n"
		if needMap[s.Name] {
			ret += "        this." + s.Name + "Map = new Map<number, " + s.Name + "Info>()\n"
		}
		ret += "\n"
	}
	ret += "        let jsonData = Res.get<cc.JsonAsset>(\"json/" + this.Name + "\", cc.JsonAsset)\n"
	for _, s := range this.Sheets {
		if len(s.Varnames) == 0 {
			continue
		}
		ret += "        this." + s.Name + "Slc = jsonData.json['" + s.Name + "']\n"
		if needMap[s.Name] {
			ret += "        this." + s.Name + "Slc.forEach(" + strings.ToLower(s.Name) + " => {\n"
			ret += "            this." + s.Name + "Map.set(" + strings.ToLower(s.Name) + ".ID, " + strings.ToLower(s.Name) + ")\n"
			ret += "        })\n"
		}
		ret += "\n"
	}
	ret += "    }\n"

	ret += "}\n"
	filePath := path + "/" + this.Name + "Cfg.ts"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if nil != err {
		return errors.New("open ts code file failed " + this.Name + "Cfg.ts")
	}
	defer file.Close()
	file.WriteString(ret)
	return nil
}

func getTsType(t string) string {
	switch t {
	case "int", "int32", "int64", "float", "float32", "float64":
		return "number"
	case "string":
		return "string"
	case "bool":
		return "boolean"
	}
	return "any"
}
