package task

func (st *subTask[T]) AddBlocking(blocking SubTask[T]) error {
	return nil
}

func (st *subTask[T]) AddControls(controls ...SubTask[T]) error {
	return nil
}

func (st *subTask[T]) AddProducers(producers ...SubTask[T]) error {
	return nil
}

func (st *subTask[T]) AddConsumers(consumers ...SubTask[T]) error {
	return nil
}
