package Stack_Queue

type QS interface {
	push(v int)
	pop() int
	top() int
	empty() bool
}
type Stack_qs struct {
	queue []int
}

func NewStack_qs() *Stack_qs {
	return &Stack_qs{queue: make([]int, 0)}
}
func (s *Stack_qs) push(v int) {
	s.queue = append(s.queue, v)
}
func (s *Stack_qs) pop() int {
	n := len(s.queue) - 1
	if n != 0 {
		val := s.queue[0]
		s.queue = s.queue[1:]
		s.queue = append(s.queue, val)
	}
	val := s.queue[0]
	s.queue = s.queue[1:]
	return val
}
func (s *Stack_qs) top() int {
	val := s.pop()
	s.queue = append(s.queue, val)
	return val
}
func (s *Stack_qs) empty() bool {
	return len(s.queue) == 0
}
