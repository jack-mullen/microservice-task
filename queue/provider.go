package queue

//a queue provider must support the following functions
type Provider interface {
	Publish(data interface{})
	Consume(c Consumer)
	//for closing connection etc.
	Cleanup()
}

