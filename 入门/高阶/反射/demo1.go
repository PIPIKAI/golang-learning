package main

import (
	"fmt"
	"reflect"
)

// 自定义类型
type myInt int

// 自定义结构体
type Student struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

func (s Student) run() {
	fmt.Printf("func run : %v\n", s)
}

func fn(o interface{}) {
	// 当我们空接口的时候可以传入任意类型
	// 我们可以利用反射来获得该变量的类型，更改该变量

	typeOf := reflect.TypeOf(o)

	fmt.Printf("name :%v kind: %v\n", typeOf.Name(), typeOf.Kind())

	// 对于指针类型 用Elem 来获得取值
	if typeOf.Kind() == reflect.Ptr {
		fmt.Printf("--是指针 该元素的name :%v kind: %v\n", typeOf.Elem().Name(), typeOf.Elem().Kind())
		if typeOf.Elem().Kind() == reflect.Struct {
			s := reflect.ValueOf(o).Elem()
			typeOfT := s.Type()
			for i := 0; i < s.NumField(); i++ {
				f := s.Field(i)
				fmt.Printf("idx:%d: name : \"%s\"type:%s =  value:%v\n", i,
					typeOfT.Field(i).Name, f.Type(), f.Interface())
			}
			// 更改 struct 里的字段
			fmt.Printf("更改 struct 里的 Id 字段\n")
			s.FieldByName("Id").SetInt(4321)
			hh, _ := o.(Student)
			hh.run()
		}

	}
	// 用反射来获得结构体中的字段，和方法
	if typeOf.Kind() == reflect.Struct {
		// 获得结构体中的字段
		// 第一种方法：Field()
		for i := 0; i < typeOf.NumField(); i++ {
			fieldType := typeOf.Field(i)
			fmt.Printf("name: %v tag: %v\n", fieldType.Name, fieldType.Tag)
		}

	}
	fmt.Println("------------------")
}

func main() {
	var myint myInt = 10
	fn(myint)
	// 可以看到并不是所有的有name ，但是都有kind
	fn(&myint)
	fn([4]int{1, 2, 3, 4})
	fn([]int{1, 2, 3, 4})
	fn([]string{"str1", "str2", "str3"})
	std := Student{
		Name: "小明",
		Id:   1234,
	}
	fn(std)
	fmt.Println(std)

	fn(&std)

	fmt.Println(std)

}
