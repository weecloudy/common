package asynctask

import (
	"fmt"
	"runtime/debug"
	"sync"

	"github.com/weecloudy/common/container/set"
	"github.com/weecloudy/common/strings"
	"github.com/weecloudy/logger"
)

var asyncTaskNames *set.HashSet

func init() {
	asyncTaskNames = set.NewSet()
}

type IAsyncTask interface {
	Run() error
}

type AsyncTask struct {
	name string
	ch   chan IAsyncTask
	wg   *sync.WaitGroup
}

func (t *AsyncTask) Close() {
	asyncTaskNames.Remove(t.name)
	close(t.ch)
	t.wg.Wait()
}

func NewAsyncTask(name string, taskChanNumber int64, goNumber int, onError func(err error)) (*AsyncTask, error) {
	if ok := asyncTaskNames.Contains(name); ok {
		return nil, fmt.Errorf("asynctask name duplicated: %v", name)
	}
	asyncTaskNames.Add(name)

	asyncTask := new(AsyncTask)
	asyncTask.name = name
	asyncTask.ch = make(chan IAsyncTask, taskChanNumber)
	asyncTask.wg = &sync.WaitGroup{}
	for i := 0; i < goNumber; i++ {
		asyncTask.wg.Add(1)
		asyncTask.run(onError)
	}
	return asyncTask, nil
}

func Recover() {
	if e := recover(); e != nil {
		logger.Errorf("panic: %v, stack: %v", e, strings.BytesToString(debug.Stack()))
	}
}

func (t *AsyncTask) run(onError func(err error)) {
	go func(name string, ch chan IAsyncTask) {
		defer t.wg.Done()
		defer Recover()

		logger.Infof("AsyncTask [%v] created", name)
		for {
			asyncTask, ok := <-ch
			if !ok {
				break
			}
			if err := asyncTask.Run(); nil != err {
				if nil != onError {
					onError(err)
				}
			}
		}
		logger.Infof("AsyncTask [%v] quit", name)
	}(t.name, t.ch)
}

func (t *AsyncTask) Post(tasks ...IAsyncTask) {
	for _, task := range tasks {
		t.ch <- task
	}
}
