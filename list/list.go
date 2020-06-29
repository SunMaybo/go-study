package list

type List struct {
	Val  int
	Next *List
}

func New() *List {
	root := new(List)
	list := root
	for i := 0; i < 8; i++ {
		list.Val = i + 1
		list.Next = new(List)
		list = list.Next
	}
	return root
}
func NewTest() *List {
	root := new(List)
	list := root
	for i := 0; i < 8; i++ {
		if i != 0 {
			p := new(List)
			p.Val = i + 1
			list = p
		}else {
			list.Val = i + 1
		}
		list = list.Next
	}
	return root
}

func Flip(list *List, k int) *List {
	length := 0
	var first *List

	for list.Next != nil {
		if length == 0 {
			first = new(List)
			first.Val = list.Val

		} else {
			bef := new(List)
			bef.Val = list.Val
			bef.Next = first
			first = bef
		}
		list = list.Next
		length++
	}
	count := length / k
	index := 0
	countIndex := 0

	var secondPtr *List
	var secondrootPtr *List
	var startPtr *List
	for first != nil {
		if index == k {
			index = 0
			countIndex++
			if secondrootPtr != nil {
				secondPtr.Val = secondrootPtr.Val
				secondPtr.Next = secondrootPtr.Next
				//secondPtr = secondrootPtr

			} else {
				secondrootPtr = new(List)
			}
			secondrootPtr.Val = startPtr.Val
			secondrootPtr.Next = startPtr.Next
		}
		if count == countIndex {
			p := new(List)
			p.Val = first.Val
			p.Next = secondrootPtr
			secondrootPtr = p
		}

		if index == 0 {
			startPtr = new(List)
			secondPtr = startPtr
		}
		secondPtr.Val = first.Val
		secondPtr.Next = new(List)
		secondPtr = secondPtr.Next
		first = first.Next
		index++
	}
	return secondrootPtr
}
