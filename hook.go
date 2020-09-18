package coralogix

import "fmt"

type Foo struct {

}

func (f Foo) FooFoo() {
	fmt.Println("foo")
}