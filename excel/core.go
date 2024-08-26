package excel

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type SheetData struct {
	Name string      // 中文描述，表示excel的sheet
	Data interface{} // 完整的struct，表示sheet中的数据
}

func GenerateExcel(sheets []SheetData, fileName string) error {
	f, err := makeExcel(sheets)
	if err != nil {
		return err
	}
	// 保存excel文件
	if err = f.SaveAs(fileName); err != nil {
		return err
	}
	// 返回nil，表示成功
	return nil
}

func MemoryCacheExcel(sheets []SheetData) (buffer *bytes.Buffer, err error) {
	buffer = &bytes.Buffer{}
	f, err := makeExcel(sheets)
	if err != nil {
		return nil, err
	}
	if err = f.Write(buffer); err != nil {
		return
	}
	return
}

// 定义一个方法，接受一个struct集合，生成一个有多个sheet的excel
// 参数：
// sheets: struct集合，每个元素包含中文描述和完整的struct
// 返回值：
// error: 错误信息，如果有的话
func makeExcel(sheets []SheetData) (f *excelize.File, err error) {
	// 创建一个新的excel文件
	f = excelize.NewFile()
	// 遍历struct集合，为每个元素创建一个sheet
	for i, sheet := range sheets {
		// 获取中文描述和完整的struct
		name := sheet.Name
		data := sheet.Data
		// 如果是第一个元素，就使用默认的Sheet1，否则就新建一个sheet
		if i == 0 {
			f.SetSheetName("Sheet1", name)
		} else {
			f.NewSheet(name)
		}
		// 获取完整的struct的类型和值
		dataType := reflect.TypeOf(data)
		dataValue := reflect.ValueOf(data)
		// 如果完整的struct不是一个切片，就返回错误
		if dataType.Kind() != reflect.Slice {
			err = fmt.Errorf("data is not a slice")
			return
		}
		// 获取切片的长度和元素类型
		length := dataValue.Len()
		elemType := dataType.Elem()
		// 如果元素类型不是一个结构体，就返回错误
		if elemType.Kind() != reflect.Struct {
			err = fmt.Errorf("slice element is not a struct")
			return
		}
		// 获取结构体的字段数量
		numField := elemType.NumField()
		// 定义一个变量，用来记录当前的行号
		row := 1
		// 定义一个切片，用来存储表头
		headers := make([]string, numField)
		// 遍历结构体的字段，获取每个字段的tag
		for j := 0; j < numField; j++ {
			// 获取字段的tag
			tag := elemType.Field(j).Tag.Get("excel")
			// 如果tag为空，就使用字段的名称
			if tag == "" {
				tag = elemType.Field(j).Name
			}
			// 将tag添加到表头切片中
			headers[j] = tag
			// 将tag写入excel的第一行，对应的列
			f.SetCellValue(name, fmt.Sprintf("%c%d", 'A'+j, row), tag)
		}
		// 行号加一
		row++
		// 遍历切片的元素，获取每个元素的字段值
		for k := 0; k < length; k++ {
			// 获取元素的值
			elemValue := dataValue.Index(k)
			// 遍历结构体的字段，获取每个字段的值
			for j := 0; j < numField; j++ {
				// 获取字段的值
				fieldValue := elemValue.Field(j).Interface()
				// 将字段的值写入excel的对应的行和列
				f.SetCellValue(name, fmt.Sprintf("%c%d", 'A'+j, row), fieldValue)
			}
			// 行号加一
			row++
		}
	}
	return
}

//// 必须定义成如下结构体
//// 完整的struct，表示sheet中的数据
//// excel表示列名
//type Student struct {
//	Name  string `excel:"姓名"` // 姓名
//	Age   int    `excel:"年龄"` // 年龄
//	Score int    `excel:"成绩"` // 成绩
//}
//
//type Teacher struct {
//	Name    string `excel:"姓名"` // 姓名
//	Subject string `excel:"科目"` // 科目
//	Salary  int    `excel:"工资"` // 工资
//}
//
//// 主函数，测试用例
//func main() {
//	// 创建一个struct集合，包含两个元素
//	sheets := []SheetData{
//		{
//			Name: "学生信息", // 中文描述，表示excel的sheet
//			Data: []Student{ // 完整的struct，表示sheet中的数据
//				{Name: "张三", Age: 18, Score: 90},
//				{Name: "李四", Age: 19, Score: 80},
//				{Name: "王五", Age: 20, Score: 70},
//			},
//		},
//		{
//			Name: "教师信息", // 中文描述，表示excel的sheet
//			Data: []Teacher{
//				{Name: "赵老师", Subject: "语文", Salary: 5000},
//				{Name: "钱老师", Subject: "数学", Salary: 6000},
//				{Name: "孙老师", Subject: "英语", Salary: 7000},
//			},
//		},
//	}
//	// 调用方法，生成excel文件
//	err := GenerateExcel(sheets, "data.xlsx")
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Println("生成成功")
//	}
//}
