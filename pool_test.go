package fixpool

import (
	"fmt"
	"sync"
	"testing"
)

var wg sync.WaitGroup

func function() {
	defer wg.Done()
	fmt.Println("funcion no arg")
}

func function1(arg interface{}) {
	defer wg.Done()
	fmt.Printf("function with one arg: %s\n", arg.(string))
}

func functionn(args ...interface{}) {
	defer wg.Done()
	fmt.Printf("args len:%d, ", len(args))
	for i := 0; i < len(args); i++ {
		fmt.Printf("%v ", args[i])
	}
	fmt.Println("")
}

func TestFunc(t *testing.T) {
	pool := NewPool(5)
	wg.Add(1)
	pool.AddTask(function)
	wg.Add(1)
	pool.AddTask(function1, "test one arg")
	wg.Add(1)
	pool.AddTask(functionn, "test", "one", "arg")
	wg.Wait()
}

func TestMany(t *testing.T) {
	pool := NewPool(5)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		pool.AddTask(function)
		wg.Add(1)
		pool.AddTask(function1, "test one arg")
		wg.Add(1)
		pool.AddTask(functionn, "test", "one", "arg")
	}
	wg.Wait()
}
