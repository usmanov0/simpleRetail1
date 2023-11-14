package foobar

type FooBar struct {
}

func NewFooBarService() *FooBar {
	return &FooBar{}
}

func (s FooBar) DoSomeDummyAction() error {

	return nil
}
