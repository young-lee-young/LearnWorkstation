/**
 * @Time:    2023/1/28 11:45 
 * @Author:  leeyoung
 * @File:    main.go
 * @Content:

剑指Offer No005 替换空格

标签：双指针
 */
package main

import "fmt"

func main() {
	s := "We are happy."

	ret := solution(s)

	fmt.Println("ret:", ret)
}

func solution(s string) string {
	sb := ([]byte)(s)

	strLen := len(sb)

	// 计算空格数量
	count := 0
	for i := 0; i < strLen; i ++ {
		if sb[i] == ' ' {
			count ++
		}
	}

	// 扩容切片：从 ' ' -> '%20'，需要扩容 2 个空格
	expand := make([]byte, count*2)
	sb = append(sb, expand...)

	// 双指针：i 指向旧字符串的尾部，j 指向扩容后字符串的尾部
	i := strLen - 1
	j := len(sb) - 1
	// ⚠️：这里是 >= 0
	for i >= 0 {
		if sb[i] != ' ' {
			sb[j] = sb[i]
			i --
			j --
		} else {
			sb[j] = '0'
			sb[j - 1] = '2'
			sb[j - 2] = '%'

			i --
			j = j - 3
		}
	}

	return string(sb)
}
