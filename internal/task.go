package internal

import "github.com/Giovanny472/gMiniPlanner/interfaces"

type subtask struct {
	subtaskname string
	progressbar int
}

type task struct {
	taskname     string
	listsubtasks interfaces.IListSubTasks
}

var ta *task
var sta *subtask

func NewSubTask() interfaces.ISubTask {
	if sta == nil {
		sta = new(subtask)
	}
	return sta
}

func (st *subtask) Name() string {
	return st.subtaskname
}

func (st *subtask) Progress() int {
	return st.progressbar
}

func NewTask(n string) interfaces.ITask {
	if ta == nil {
		ta = &task{taskname: n, listsubtasks: newListSubTask()}
	}
	return ta
}

func newListSubTask() interfaces.IListSubTasks {
	return make(interfaces.IListSubTasks, 0)
}

func (t *task) Name() string {
	return t.taskname
}

func (t *task) ListSubTasks() interfaces.IListSubTasks {
	return t.listsubtasks
}
