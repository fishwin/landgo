package stack_queue

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewQueueVia2Stack(t *testing.T) {

	Convey("test queue via 2 stack", t, func() {

		q := NewQueueVia2Stack()
		q.EnQueue(1)
		q.EnQueue(2)
		q.EnQueue(3)

		ret, err := q.DeQueue()
		So(ret, ShouldEqual, 1)
		So(err, ShouldBeNil)

		q.EnQueue(4)

		ret, err = q.DeQueue()
		So(ret, ShouldEqual, 2)
		So(err, ShouldBeNil)

		ret, err = q.DeQueue()
		So(ret, ShouldEqual, 3)
		So(err, ShouldBeNil)

		q.EnQueue(5)

		ret, err = q.DeQueue()
		So(ret, ShouldEqual, 4)
		So(err, ShouldBeNil)

		ret, err = q.DeQueue()
		So(ret, ShouldEqual, 5)
		So(err, ShouldBeNil)
	})

}
