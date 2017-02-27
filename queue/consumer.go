package queue

//a consumer must have a Process function that takes a message as it's only
//argument.
type Consumer interface {
	Process(*Message) bool
}
