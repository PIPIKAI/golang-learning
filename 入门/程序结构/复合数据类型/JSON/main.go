package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	// 用 json:name 这种形式定义结构体转化为json后的key的名称 这叫做结构体的tag
	Id       string `json:"id"`
	Name     string `json:"name"`
	Age      int
	birthDay string // 首字母小写代表私有，json这个包不能访问
}

func main() {
	stud := Student{
		"1234",
		"王小明",
		4,
		"2018-6-19",
	}
	fmt.Printf("stud: %v\n", stud)
	// 结构体转换为字符串
	// 返回值的类型是byte
	studByte, _ := json.Marshal(stud)
	// 输出的时候强转为字符串型
	fmt.Printf("studJson: %v\n", string(studByte))

	// 字符串 转化为结构体类型  // 反序列化

	studStr := string(studByte)
	var studStruct Student
	err := json.Unmarshal([]byte(studStr), &studStruct)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("studStruct: %#v\n", studStruct)
}
