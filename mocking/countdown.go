package mocking

// My plan for chapter is not to do 1:1 as written in tutorial but rather take more relax approach to mocking/stubbing

import (
	"fmt"
	"time"
)

// It would be better to use two interfaces but there is a problem
// How can we test in what order do we sleep and write :/
type Spy interface {
	SpySleep(time.Duration)
	SpyWrite(s string)
}

type ProdSpy struct{}

func (n ProdSpy) SpySleep(t time.Duration) {
	time.Sleep(t)
}

func (n ProdSpy) SpyWrite(s string) {
	fmt.Print(s)
}

type Event struct {
	Name    string
	Content string
}

func countdown(spy Spy) {
	for i := 3; i > 0; i-- {
		str := fmt.Sprintf("%d\n", i)

		spy.SpyWrite(str)
		// we could remove this hardcoded 1 second.
		// create some dependancy that is injected with this time as in tutorial
		spy.SpySleep(1 * time.Second)
	}
	spy.SpyWrite("Go!")
}

func Countdown() {
	countdown(&ProdSpy{})
}
