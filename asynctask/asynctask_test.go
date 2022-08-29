package asynctask_test

import (
	"errors"
	"github.com/weecloudy/common/asynctask"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type TestJob struct {
	name string
}

func NewTestJob(name string) *TestJob {
	return &TestJob{name: name}
}

func (t *TestJob) Run() error {
	if t.name == "err" {
		return errors.New("error")
	}

	return nil
}

var _ = Describe("Asynctask", func() {
	Context("new", func() {
		It("one", func() {
			task, err := asynctask.NewAsyncTask("new", 3, 3, func(err error) {
			})
			Expect(task).NotTo(BeNil())
			Expect(err).To(BeNil())
			task.Close()
		})
		It("many same name", func() {
			task, err := asynctask.NewAsyncTask("new", 3, 3, func(err error) {
			})
			Expect(task).NotTo(BeNil())
			Expect(err).To(BeNil())
			task, err = asynctask.NewAsyncTask("new", 3, 3, func(err error) {
				if err != nil {
					println(err.Error())
					return
				}
			})
			Expect(task).To(BeNil())
			Expect(err).NotTo(BeNil())
			task.Close()
		})
	})

	Context("post task", func() {
		It("run ok", func() {
			task, _ := asynctask.NewAsyncTask("new", 3, 3, func(err error) {
				Expect(err).To(BeNil())
			})
			task.Post(NewTestJob("hello"))
			task.Close()
		})
		It("run error", func() {
			task, _ := asynctask.NewAsyncTask("new", 3, 3, func(err error) {
				Expect(err).NotTo(BeNil())
			})
			task.Post(NewTestJob("err"))
			task.Close()
		})
	})
})
