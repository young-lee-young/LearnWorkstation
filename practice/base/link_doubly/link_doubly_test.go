package link_doubly

import (
	"testing"
	"fmt"
)

func TestDoublyLink(t *testing.T) {
	link := LinkDoublyList{}

	// 测试右侧插入元素
	i := 0
	for i < 0 {
		link.RightPush(i)
		i ++
	}
	link.ShowList()


	// 测试左侧删除元素
	result := link.LeftPop()
	if result != nil {
		fmt.Println(result)
	}
	link.ShowList()

	// 测试获取长度方法
	fmt.Println(link.GetLength())
}
