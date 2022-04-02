package requests

type IValidatable interface {
	Validate() bool
}

type RequestConstructor func() interface{}
