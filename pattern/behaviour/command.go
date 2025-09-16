package behaviour

type Command[T any] interface {
	Execute(T) error
	Undo(T) error
}
