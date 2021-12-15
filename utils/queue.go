package utils

type Queue struct {
	q chan interface{}
}

func (this *Queue) Init() {
	this.q = make(chan interface{}, 100)
}

func (this *Queue) PushBack(data interface{}) {
	this.q <- data
}

func (this *Queue) Pop() chan interface{} {
	return this.q
}
