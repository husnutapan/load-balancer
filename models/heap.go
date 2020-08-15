package models

func (p Pool) Len() int {
	return len(p)
}

func (p Pool) Less(i, j int) bool {
	return p[i].pending < p[j].pending
}

func (p *Pool) Swap(i, j int) {
	a := *p
	a[i], a[j] = a[j], a[i]
	a[i].id = i
	a[j].id = j
}

func (p *Pool) Push(x interface{}) {
	n := len(*p)
	item := x.(*Work)
	item.id = n
	*p = append(*p, item)
}

func (p *Pool) Pop() interface{} {
	old := *p
	n := len(old)
	item := old[n-1]
	item.id = -1
	*p = old[0 : n-1]
	return item
}
