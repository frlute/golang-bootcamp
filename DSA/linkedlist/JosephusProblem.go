package linkedlist

type JosephusSolutaion struct {
	cll   *CircularLinkedList
	total int
}

func newJosephusSolutaion(size int) *JosephusSolutaion {
	solution := &JosephusSolutaion{
		cll:   newCircularLinkedList(),
		total: 0,
	}
	solution.init(size)
	return solution
}

func (jose *JosephusSolutaion) init(size int) {
	for i := 0; i < size; i++ {
		jose.cll.insert(i + 1)
	}
	jose.total = size
	jose.cll.display()
}

func (jose *JosephusSolutaion) start(hops int) {
	jose.cll.setMarkerPosition(1)
	for jose.total > 1 {
		for i := 1; i <= hops; i++ {
			jose.cll.move()
		}
		jose.cll.delete()
		jose.total--
	}
	jose.cll.display()
}
