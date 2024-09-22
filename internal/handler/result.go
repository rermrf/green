package handler

type Result[T any] struct {
	Code int
	Msg  string
	Data T
}
