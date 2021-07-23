package model

type Object interface{}
type IMoney interface {
	Equals(money *Object) bool
}
