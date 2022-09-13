package link

import (
	"testing"
	"fmt"
)

func TestLink(t *testing.T) {
	linkList := LinkedList{}

	// 测试是否为空方法
	fmt.Println(linkList.IsEmpty())
	fmt.Println(linkList.GetLength())

	// 测试尾部添加元素方法
	count := 0
	for count < 10 {
		linkList.Append(count)
		count += 1
	}
	linkList.ShowList()

	// 测试获取长度方法
	fmt.Println(linkList.IsEmpty())
	fmt.Println(linkList.GetLength())
	fmt.Println(linkList.Contains(3))

	// 测试插入元素
	linkList.Insert(5, 200)
	linkList.ShowList()

	// 测试头部添加元素方法
	linkList.Add(100)
	linkList.ShowList()

	// 测试删除头部元素
	linkList.DeleteHead()
	linkList.ShowList()

	// 测试删除尾部元素
	linkList.DeleteTail()
	linkList.ShowList()

	// 测试删除任意元素
	linkList.Delete(3)
	linkList.ShowList()
}
