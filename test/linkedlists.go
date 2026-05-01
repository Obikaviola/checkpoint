package main

import "fmt"

type node struct{
	data int
	next *node
}

type linkedList struct{
	head *node
	length int
}

//  {Method receiver}     {what it accepts}
func (l *linkedList) prepend(n *node){
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l linkedList) printListData(){
	toPrint := l.head
	for l.length != 0{
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

func (l *linkedList) deleteWithValue(value int){
	if l.length == 0{
		return 
	}

	if l.head.data == value{
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head
	for previousToDelete.next.data != value{
		if previousToDelete.next.next == nil{
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func main(){
	mylist := linkedList{}
	node1 := &node{data:48}
	node2 := &node{data:10}
	node3 := &node{data:45}
	node4 := &node{data:16}
	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)
	mylist.prepend(node4)
	fmt.Println(mylist)
	mylist.printListData()
	mylist.deleteWithValue(100)
	mylist.deleteWithValue(16)
	mylist.printListData()
	emptyList := linkedList{}
	emptyList.deleteWithValue(10)


	// var mylist linkedList
	// mylist.addNodeToEnd(1)
	// mylist.addNodeToEnd(2)
	// mylist.addNodeToEnd(3)
	// mylist.addNodeToEnd(4)
	// mylist.printList()
}

func (list *linkedList) addNodeToEnd (data int){
	n := &node{data, nil}
	if list.head == nil{
		list.head = n
	} else {
		// cn is current node 
		//start at the beginning
		cn := list.head
		// loop through the list to find the end of the list
		for cn.next != nil{
			cn = cn.next
		}
		cn.next = n
	}
}

func (list *linkedList) printList(){
	// start at the beginning
	cn := list.head
	// loop through the list
	for cn != nil{
		fmt.Printf("current node address: %p\t node data: %d \t next node: %p\n", cn, cn.data, cn.next)
		cn = cn.next
	}
}