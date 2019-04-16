package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/Luxurioust/excelize"
)

func ListDir(dirPth string, suffix string) (files []string, err error) {
	files = make([]string, 0, 10)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	PthSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}

	return files, nil
}

func GetClassDict(sysName string) (map[string]string, map[string]string, map[string]string, map[string]string) {
	classPath := "C:/Users/Chao/Documents/github/gakki_python/NBL3_autogen_csv/Class"
	classFile := []string{}
	classFile, _ = ListDir(classPath, ".xlsx")
	for _, v := range classFile {
		fmt.Println(v)
	}
	sysClassifyDict := make(map[string]string)
	fmt.Println(sysClassifyDict)
	if len(classFile) > 1 {
		for _, v := range classFile {
			re := regexp.MustCompile("NBL3 (.*)CLASS (.*) V(.*)")
			sys := re.FindStringSubmatch(v)[2]
			fmt.Println(sys)
			_, ok := sysClassifyDict[sys]
			if ok == false {
				sysClassifyDict[sys] = v
			}
		}
	}
	classListPtname := make(map[string]string)
	classListEqptype := make(map[string]string)
	classListDeadband := make(map[string]string)
	classListScaling := make(map[string]string)
	v, ok := sysClassifyDict[sysName]
	if ok {
		xlsx, err := excelize.OpenFile(v)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		rows, _ := xlsx.GetRows("CLASS")
		xlsx.GetCellStyle()
		for _, row := range rows[2:] {
			EqptCode := row[8]
			AttributeDescription := row[11]
			DCDataType := row[13]
			EqpNumStr := EqptCode + AttributeDescription + DCDataType
			PointName := row[122]
			CFGElementName := row[120]
			Deadband := row[60]
			Scaling := row[58]
			_, ok := classListPtname[EqpNumStr]
			if ok == false {
				classListPtname[EqpNumStr] = PointName
			}
			_, ok = classListEqptype[EqpNumStr]
			if ok == false {
				classListEqptype[EqpNumStr] = CFGElementName
			}
			_, ok = classListDeadband[EqpNumStr]
			if ok == false {
				classListDeadband[EqpNumStr] = Deadband
			}
			_, ok = classListScaling[EqpNumStr]
			if ok == false {
				classListScaling[EqpNumStr] = Scaling
			}
			/*for _, colCell := range row {
				fmt.Print(colCell, "\t")
			}*/
		}

	} else {
		fmt.Println("无对应class,请检查!")
		os.Exit(1)
	}
	return classListPtname, classListEqptype, classListDeadband, classListScaling
}

//Get value from cell by given sheet index and axis
/*cell, err := xlsx.GetCellValue("CLASS", "L3")
fmt.Println(cell, err)*/
// Get all the rows in a sheet.
//index := xlsx.GetSheetIndex("CLASS")

func main() {
	classListPtname := make(map[string]string)
	classListEqptype := make(map[string]string)
	classListDeadband := make(map[string]string)
	classListScaling := make(map[string]string)
	classListPtname, classListEqptype, classListDeadband, classListScaling = GetClassDict("BAS")
	fmt.Println(classListPtname, classListEqptype, classListDeadband, classListScaling)
}
