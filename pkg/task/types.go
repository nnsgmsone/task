package task

import "encoding"

// Task is a basic distributed task that can run on different nodes
type Task[T any] interface {
	// Run the task until the entire task is finished or an error is encountered
	Run() error
	// Close the task, this action will force the task to stop
	Close() error

	// NewSubTask create a new subtask:
	//  address: the current subtask's running address
	//  op: subtask's operator
	NewSubTask(address string, op TaskOp[T]) SubTask[T]
}

type TaskOp[T any] interface {
	// Close the current operator and clean up the resources
	Close() error
	// Run when nil or error is returned,
	// 	the op will end and the whole task will be automatically ended
	Run(T, T, T) (T, error)

	IsEnd(T) bool

	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}

type SubTask[T any] interface {
	Run() error
	Close() error
	//  AddBlocking: blocking task means that the task will be blocked by other tasks,
	// 		blocking signal may be a useful information the the task must be consumed
	AddBlocking(blocking SubTask[T]) error
	//  AddControls: indicates those tasks that will send control signals to the task
	AddControls(controls ...SubTask[T]) error
	// 	AddProducers: add producers of the subtask,
	//		empty producers means that the task does not accept input from other tasks
	AddProducers(producers ...SubTask[T]) error
	// 	AddConsumers: add consumers of the subtask,
	//  	if empty means no task needs to consume the message of the task
	AddConsumers(consumers ...SubTask[T]) error

	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}

type task[T any] struct {
	subTasks []*subTask[T]
}

type subTask[T any] struct {
	id   int
	addr string
	t    *task[T]
	op   TaskOp[T]
}
