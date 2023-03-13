package task

func NewTask[T any]() Task[T] {
	return &task[T]{}
}

func (t *task[T]) Run() error {
	return nil
}

func (t *task[T]) Close() error {
	return nil
}

func (t *task[T]) NewSubTask(address string, op TaskOp[T]) SubTask[T] {
	return &subTask[T]{
		op:   op,
		addr: address,
	}
}
