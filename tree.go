package interview

type Tree struct {
	left *Tree
	value *int
	right *Tree
}

func (s *Tree) Left() *Tree {
	return s.left
}

func (s *Tree) Right() *Tree {
	return s.right
}

func (s Tree) Value() int {
	return *s.value
}

func (s Tree) Values() []int {
	var ret = []int{}
	if s.Left() != nil {
		for _, v := range s.Left().Values() {
			ret = append(ret, v)
		}
	}
	if s.value != nil {
		ret = append(ret, s.Value())
	}
	if s.Right() != nil {
		for _, v := range s.Right().Values() {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *Tree) Insert(v int) {
	if s.value == nil {
		s.value = new(int);
		*s.value = v
	} else if v <= *s.value {
		if s.Left() == nil {
			s.left = new(Tree)
			*s.left = Tree{nil,nil,nil}
		}
		(*s.left).Insert(v)
	} else {
		if s.Right() == nil {
			s.right = new(Tree)
			*s.right = Tree{nil,nil,nil}
		}
		(*s.right).Insert(v)
	}
}
