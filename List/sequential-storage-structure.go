package main

//线性表中的顺序存储结构

import (
	"fmt"
)

type Elem int

type SqList struct {
	//最大长度
	maxsize int
	// 当前长度
	length int
	//保存数据
	data   []Elem
}

//初始化
func New(maxsize int) *SqList {
	return &SqList{maxsize: maxsize, data: make([]Elem, maxsize)}
}

//检查线性表是否为空
func (list *SqList) IsEmpty() bool {
	return 0 == list.length
}

//判断线性表是否已满
func (list *SqList) IsFull() bool {
	return list.length == list.maxsize
}

//在i个位置之前插入新的元素e
func (list *SqList) Insert(i int, e Elem) bool {
	if i < 1 || i > list.length {
		fmt.Println("pls check i:", i)
		return false
	}

	for k := list.length; k > i-1; k-- {
		list.data[k] = list.data[k-1]
	}
	list.data[i-1] = e
	list.length++
	return true
}

//删除位置为i的元素
func (list *SqList) Del(i int) bool {
	if i < 1 || i > list.length {
			fmt.Println("pls check i:", i)
		return false
	}
	for k := i - 1; k < list.length-1; k++ {
		list.data[k] = list.data[k+1]
	}

	list.data[list.length-1] = 0
	list.length--
	return true
}

//获取第i个位置的元素
func (list SqList) GetElem(i int) Elem{
	if i < 1 || i > list.length {
		fmt.Println("pls check i:", i)
		return -1
	}
	return list.data[i-1]
}

//追加一个元素
func (list *SqList) append(e Elem) bool {
	if list.IsFull() {
		fmt.Println("list is fulle")
		return false
	}
	list.data[list.length] = e
	list.length++
	return true
}

func main() {
	sq := New(10)

	sq.append(99)
	sq.append(999)
	sq.append(9999)
	sq.append(99999)
	fmt.Println(sq)

	sq.Insert(4, 888)
	fmt.Println(sq)

	//fmt.Println(r)
	sq.Del(2)
	fmt.Println(sq)

	fmt.Println(sq.GetElem(3))
}
