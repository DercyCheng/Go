package Stack_Queue

type Stack_sq []int

func (s *Stack_sq) push(v int) {
	*s = append(*s, v)
}
func (s *Stack_sq) pop() int {
	val := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return val
}
func (s *Stack_sq) peek() int {
	return (*s)[len(*s)-1]
}
func (s *Stack_sq) size() int {
	return len(*s)
}
func (s *Stack_sq) empty() bool {
	return len(*s) == 0
}

type Queue_sq struct {
	stIn  *Stack_sq
	stOut *Stack_sq
}

func Constructor_sq() Queue_sq {
	return Queue_sq{stIn: &Stack_sq{}, stOut: &Stack_sq{}}
}

func (q *Queue_sq) push(v int) {
	q.stIn.push(v)
}
func (q *Queue_sq) pop() int {
	q.fillstackout()
	return q.stOut.pop()
}
func (q *Queue_sq) peek() int {
	q.fillstackout()
	return q.stOut.peek()
}
func (q *Queue_sq) empty() bool {
	return q.stIn.empty() && q.stOut.empty()
}
func (q *Queue_sq) fillstackout() {
	if q.stOut.empty() {
		for !q.stIn.empty() {
			val := q.stIn.pop()
			q.stOut.push(val)
		}
	}
}
