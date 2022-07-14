/*练习 4.9： 编写一个程序wordfreq程序，报告输入文本中每个单词出现的频率。
在第一次调用Scan前先调用input.Split(bufio.ScanWords)函数，
这样可以按单词而不是按行输入。*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	// input.Split(bufio.ScanWords)
	wordfreq := make(map[string]int, 0)
	for input.Scan() {
		for _, word := range strings.Split(input.Text(), " ") {
			wordfreq[word]++
		}
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "4.9 %v \n", err)
	}
	for key, value := range wordfreq {
		fmt.Printf("word: %s count: %d\n", key, value)
	}
}

/*

The length of an array is a component of the array type, so [3]int and [4]int are two different array types. The length of the array must be a constant expression, because the length of the array needs to be determined at the compilation stage.
^Z
word: stage. count: 1
word: is count: 1
word: [4]int count: 1
word: expression, count: 1
word: compilation count: 1
word: the count: 5
word: an count: 1
word: to count: 1
word: length count: 3
word: array count: 5
word: different count: 1
word: two count: 1
word: needs count: 1
word: determined count: 1
word: at count: 1
word: of count: 4
word: type, count: 1
word: and count: 1
word: are count: 1
word: component count: 1
word: [3]int count: 1
word: must count: 1
word: be count: 2
word: constant count: 1
word: because count: 1
word: The count: 2
word: a count: 2
word: so count: 1
word: types. count: 1
*/
