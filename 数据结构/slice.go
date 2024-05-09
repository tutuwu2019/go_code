package main

import "fmt"

func ps(s []int) {
	fmt.Println("len = ", len(s), " cap = ", cap(s), "s = ", s)
}

// 切片的长度和容量
// 长度值得是其包含元素个数
// 容量指的是从它的第一个元素开始数，到其底层数组元素末尾的个数
func main() {
	s := []int{2, 3, 5, 7, 11, 15}
	fmt.Print("原始的s: ")
	ps(s)

	s = s[:0]
	fmt.Print("s = s[:0]操作后的s: ")
	ps(s)

	s = s[:4]
	fmt.Print("s = s[:4]操作后的s: ")
	ps(s)

	s = s[2:]
	fmt.Print("s = s[2:]操作后的s: ")
	ps(s)
	s = append(s, 1, 2, 3, 4, 5)
	fmt.Print("s = append(s, 1, 2, 3, 4, 5)操作后的s: ")
	ps(s)
	new_ := make([]int, len(s), cap(s)*2)
	copy(new_, s)
	fmt.Print("copy(new_, s)操作后的new_: ")
	ps(new_)

}



/*
原始的s: len =  6  cap =  6 s =  [2 3 5 7 11 15]
s = s[:0]操作后的s: len =  0  cap =  6 s =  []
s = s[:4]操作后的s: len =  4  cap =  6 s =  [2 3 5 7]
s = s[2:]操作后的s: len =  2  cap =  4 s =  [5 7]
s = append(s, 1, 2, 3, 4, 5)操作后的s: len =  7  cap =  8 s =  [5 7 1 2 3 4 5]
copy(new_, s)操作后的new_: len =  7  cap =  16 s =  [5 7 1 2 3 4 5]
*/
