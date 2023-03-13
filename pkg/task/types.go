package task

import "encoding"

// Task is a basic distributed task that can run on different nodes
type Task[T any] interface {
	// Run the task until the entire task is finished or an error is encountered
	//  	fill is use to receive the results of the task
	Run(fill func(T) error) error
	// Close the task, this action will force the task to end
	Close() error

	// NewSubTask create a new subtask:
	// 	producers: producers of the current subtask,
	//		empty producers means that the task does not accept input from other tasks
	// 	consumers: consumers of the current subtask,
	//  	if empty means no task needs to consume the message of the current task
	//  controls: indicates those tasks that will send control signals to the current task
	//  blocking: blocking task means that the current subtask will be blocked by other tasks,
	// 		blocking signal may be a useful information
	//  address: the current subtask's running address
	//  op: subtask's operator
	NewSubTask(producers, consumers, controls []SubTask[T],
		blocking SubTask[T], address string, op TaskOp[T]) SubTask[T]
}

type TaskOp[T any] interface {
	// Close the current operator and clean up the resources
	Close() error
	// Run when nil or error is returned,
	// 	the op will end and the whole task will be automatically ended
	Run() (T, error)

	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}

type SubTask[T any] struct {
	id int
	op TaskOp[T]
}

type task[T any] struct {
}
