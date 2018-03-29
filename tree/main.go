package main

import "fmt"

type tree struct {
	left  *tree
	right *tree
	value int
}

func (t *tree) Add(v int) {
	if v > t.value {
		if t.right == nil {
			t.right = &tree{value: v}
		} else {
			t.right.Add(v)
		}
	} else {
		if t.left == nil {
			t.left = &tree{value: v}
		} else {
			t.left.Add(v)
		}
	}
}

func NewTree(in []int) *tree {
	t := &tree{value: in[0]}

	for _, v := range in[1:] {
		t.Add(v)
	}

	return t
}

func (t *tree) Walk(c chan [2]int) {
	ch := make(chan [2]int)
	t.walk(c, 0)
	close(c)
}

func (t *tree) walk(c chan [2]int, d int) {
	if t.left != nil {
		t.left.walk(c, d+1)
	}
	c <- [...]int{t.value, d}
	if t.right != nil {
		t.right.walk(c, d+1)
	}
}

func (t *tree) String() string {
	ch := make(chan [2]int)
	go func() { t.Walk(ch) }()
	x := []int{}
	for v := range ch {
		x = append(x, v[0])
	}
	return fmt.Sprintf("%v", x)
}

func Eq(x, y *tree) bool {
	xChan := make(chan [2]int)
	yChan := make(chan [2]int)

	go func() { x.Walk(xChan) }()
	go func() { y.Walk(yChan) }()
	for {
		l, lok := <-xChan
		r, rok := <-yChan

		fmt.Printf("l: %d, r: %d, lok: %v, rok: %v\n", l, r, lok, rok)
		if !lok && !rok {
			return true
		}
		if !lok || !rok {
			return false
		}
		if l[0] != r[0] {
			return false
		}
	}
}

func main() {
	t := NewTree([]int{7, 5, 3, 9, 5, 8, 3, 2})
	v := NewTree([]int{2, 3, 3, 5, 5, 7, 8, 9})

	fmt.Println(t)
	if !Eq(t, v) {
		fmt.Println(v)
	}

}
