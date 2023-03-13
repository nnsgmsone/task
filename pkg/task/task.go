package task

import "sync"

func NewTask[T any]() Task[T] {
	return &task[T]{}
}

func (t *task[T]) Run() error {
	var err error
	var wg sync.WaitGroup

	ch := make(chan error, len(t.subTasks))
	for i := range t.subTasks {
		wg.Add(1)
		go func(st *subTask[T]) {
			defer wg.Done()
			ch <- st.Run()
		}(t.subTasks[i])
	}
	wg.Wait()
	for len(ch) > 0 {
		if err0 := <-ch; err0 != nil {
			err = err0
		}
	}
	return err
}

func (t *task[T]) Close() error {
	var err error

	for _, st := range t.subTasks {
		if err0 := st.Close(); err0 != nil {
			err = err0
		}
	}
	return err
}

func (t *task[T]) NewSubTask(address string, op TaskOp[T]) SubTask[T] {
	return &subTask[T]{
		op:   op,
		addr: address,
	}
}
