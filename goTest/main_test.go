package main

import (
	"testing"
	"fmt"
	"os"
)

func IsShengxu(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		if i != 0 && arr[i-1] > arr[i] {
			return false
		}
	}
	return true
}

// 最基本的使用
// func TestSort(t *testing.T) {
// 	testArr1 := []int{4, 3, 5, 2, 1}
// 	Sort(testArr1)

// 	if !IsShengxu(testArr1) {
// 		t.Errorf("Sort arr expect [1 2 3 4 5],but %v got", testArr1)
// 	}
// }
// 最基本的使用
// func TestSort(t *testing.T) {

// 	testArr1 := []int{4, 3, 5, 2, 1}
// 	copyTestArr := make([]int, len(testArr1))
// 	copy(copyTestArr, testArr1)
// 	Sort(testArr1)
// 	sort.Ints(copyTestArr)
// 	if !IsShengxu(testArr1) {
// 		t.Errorf("Sort arr expect %v,but %v got", copyTestArr, testArr1)
// 	}
// }

// 子测试(Subtests)
// 可以在某个测试用例中，根据测试场景使用 t.Run创建不同的子测试用例
// func TestSort(t *testing.T) {
// 	t.Helper()
// 	t.Run("arr1", func(t *testing.T) {
// 		testArr1 := []int{4, 3, 5, 2, 1}
// 		copyTestArr := make([]int, len(testArr1))
// 		copy(copyTestArr, testArr1)
// 		Sort(testArr1)
// 		sort.Ints(copyTestArr)
// 		if !IsShengxu(testArr1) {
// 			// 使用Error/Errorf 遇到错误不会停止，会继续完成剩下的测试，
// 			// t.Errorf("Sort arr expect %v,but %v got", copyTestArr, testArr1)
// 			//  使用Fatal/Fatalf 遇到错误会停止剩下的测试
// 			// t.Fatalf("Sort arr expect %v,but %v got", copyTestArr, testArr1)
// 			t.Fatal("fail")
// 		}
// 	})
// 	t.Run("arr2", func(t *testing.T) {
// 		testArr1 := []int{1, 3, 2, 4, 5}
// 		copyTestArr := make([]int, len(testArr1))
// 		copy(copyTestArr, testArr1)
// 		Sort(testArr1)
// 		sort.Ints(copyTestArr)
// 		if !IsShengxu(testArr1) {
// 			// t.Fatalf("Sort arr expect %v,but %v got", copyTestArr, testArr1)
// 			// t.Errorf("Sort arr expect %v,but %v got", copyTestArr, testArr1)
// 			t.Fatal("fail")
// 		}
// 	})

// }

// calc_test.go
// func TestMul(t *testing.T) {
// 	t.Run("pos", func(t *testing.T) {
// 		if Mul(2, 3) != 6 {
// 			t.Fatal("fail")
// 		}

// 	})
// 	t.Run("neg", func(t *testing.T) {
// 		if Mul(2, -3) != -6 {
// 			t.Fatal("fail")
// 		}
// 	})
// }
// type calcCase struct{ A, B, Expected int }

// func createMulTestCase(t *testing.T, c *calcCase) {
// 	t.Helper()
// 	if ans := Mul(c.A, c.B); ans != c.Expected {
// 		t.Fatalf("%d * %d expected %d, but %d got",
// 			c.A, c.B, c.Expected, ans)
// 	}

// }
// func TestMul(t *testing.T) {
// 	createMulTestCase(t, &calcCase{2, 3, 6})
// 	createMulTestCase(t, &calcCase{2, -3, -6})
// 	createMulTestCase(t, &calcCase{2, 0, 1}) // wrong case
// }

// setup 和 teardown
// 如果在同一个测试文件中，每一个测试用例运行前后的逻辑是相同的，一般会写在 setup 和 teardown 函数中。
// 例如执行前需要实例化待测试的对象，如果这个对象比较复杂，很适合将这一部分逻辑提取出来；
// 执行后，可能会做一些资源回收类的工作，例如关闭网络连接，释放文件等。标准库 testing 提供了这样的机制：

func setup() {
	fmt.Println("Before all tests")
}

func teardown() {
	fmt.Println("After all tests")
}

func Test1(t *testing.T) {
	fmt.Println("I'm test1")
}

func Test2(t *testing.T) {
	fmt.Println("I'm test2")
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
