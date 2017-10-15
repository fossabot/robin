package core

import "reflect"

type IExecutor interface {
	ExecuteTasks(t []Task)
	ExecuteTasksWithGoroutine(t []Task)
}

type Task struct {
	Func   interface{}
	Params []interface{}
}

func NewTask(t interface{}, p ...interface{}) Task {
	return Task{Func: t, Params: p}
}

type defaultExecutor struct {
}

func NewDefaultExecutor() defaultExecutor { return defaultExecutor{} }

func (d defaultExecutor) ExecuteTasks(tasks []Task) {
	for _, task := range tasks {
		task.Run()
	}
}

func (d defaultExecutor) ExecuteTasksWithGoroutine(tasks []Task) {
    for _, task := range tasks {
        go task.Run()
    }
}

func (t Task) Run() {
	execFunc := reflect.ValueOf(t.Func)
	params := make([]reflect.Value, len(t.Params))
	for k, param := range t.Params {
		params[k] = reflect.ValueOf(param)
	}
	func(in []reflect.Value) { _ = execFunc.Call(in) }(params)
}
