package task

func NewTask[T any]() Task[T] {
	return &task[T]{}
}

func (t *task[T]) Run(fill func(T) error) error {
	return nil
}

func (t *task[T]) Close() error {
	return nil
}

func (t *task[T]) NewSubTask(producers, consumers, controls []SubTask[T],
	blocking SubTask[T], address string, op TaskOp[T]) SubTask[T] {
	return SubTask[T]{}
}
