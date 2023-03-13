package task

func (st *subTask[T]) Close() error {
	return st.op.Close()
}

func (st *subTask[T]) Run() error {
	blocking := st.getBlockingData()
	for {
		r, err := st.op.Run(blocking, st.getControlsData(),
			st.getProducersData())
		if err != nil {
			return err
		}
		if st.op.IsEnd(r) {
			return nil
		}
		st.setConsumersData(r)
	}
}

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

func (st *subTask[T]) MarshalBinary() ([]byte, error) {
	return nil, nil
}

func (st *subTask[T]) UnmarshalBinary(data []byte) error {
	return nil
}

func (st *subTask[T]) getBlockingData() T {
	var data T
	return data
}

func (st *subTask[T]) getControlsData() T {
	var data T
	return data
}

func (st *subTask[T]) getProducersData() T {
	var data T
	return data
}

func (st *subTask[T]) setConsumersData(data T) {
}
