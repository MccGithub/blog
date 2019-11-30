package main

import (
	"fmt"
	"sync"
)

func deferAndGo() {
	s := sync.WaitGroup{}
	s.Add(10)
	for i := 0; i < 10; i++ {
		defer fmt.Println("[fmt defer]:", i)
		defer func() {
			fmt.Println("[func defer]:", i)
		}()
		go fmt.Println("[fmt go]:", i)
		go func() {
			fmt.Println("[func go]:", i)
			s.Done()
		}()
	}
	s.Wait()
}

type S struct {
	a, b int
}

// String implements the fmt.Stringer interface
func (s *S) String() string {
	return fmt.Sprintf("%s", s) // Sprintf will call s.String()
}

func main() {
	s := &S{a: 1, b: 2}
	fmt.Println(s)
}
