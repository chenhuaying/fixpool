package fixpool

type Pool struct {
	num   int
	queue []chan Command
	leaky chan int
}

type Command struct {
	function interface{}
	args     []interface{}
}

func NewPool(num int) *Pool {
	queue := make([]chan Command, 0, num)
	for i := 0; i < num; i++ {
		queue = append(queue, make(chan Command))
	}
	leaky := make(chan int, num)
	for idx := 0; idx < num; idx++ {
		leaky <- idx
	}
	pool := &Pool{num: num, queue: queue, leaky: leaky}

	for l := 0; l < num; l++ {
		go func(idx int) {
			for {
				select {
				case cmd := <-pool.queue[idx]:
					if len(cmd.args) > 1 {
						cmd.function.(func(...interface{}))(cmd.args...)
					} else if len(cmd.args) == 1 {
						cmd.function.(func(interface{}))(cmd.args[0])
					} else {
						cmd.function.(func())()
					}
					leaky <- idx
				}
			}
		}(l)
	}

	return pool
}

func (p *Pool) AddTask(fc interface{}, args ...interface{}) {
	var cmd Command
	cmd.function = fc
	cmd.args = args
	idx := <-p.leaky
	p.queue[idx] <- cmd
}
