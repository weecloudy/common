##### 介绍

异步执行任务，任务发至channel中，协程异步处理，执行任务错误回调处理

##### 场景

进程启动时，初始化程异步任务

##### 功能

```go
// 初始异步任务
func NewAsyncTask(name string, taskChanNumber int64, goNumber int, onError func(err error)) (*AsyncTask, error)

// 提交任务(实现 Run() error)
func (this *AsyncTask) Post(task ...IAsyncTask) 

type IAsyncTask interface {
	Run() error
}

// 关闭异步任务
func (this *AsyncTask) Close()


```

