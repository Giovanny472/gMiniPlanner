package interfaces

type ISubTask interface {
	Name() string
	Progress() int
}

type IListSubTasks []ISubTask

type ITask interface {
	Name() string
	ListSubTasks() IListSubTasks
}
