package shangguigu

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewQueue(t *testing.T) {
	Convey("test queue", t, func() {
		queue := NewQueue(5)
		So(queue, ShouldNotBeNil)

		err := queue.EnQueue(1)
		So(err, ShouldBeNil)

		err = queue.EnQueue(2)
		So(err, ShouldBeNil)

		err = queue.EnQueue(3)
		So(err, ShouldBeNil)

		err = queue.EnQueue(4)
		So(err, ShouldBeNil)

		err = queue.EnQueue(5)
		So(err, ShouldBeNil)

		err = queue.EnQueue(6)
		So(err, ShouldNotBeNil)
		queue.Print()

		v, err := queue.Head()
		So(err, ShouldBeNil)
		So(v, ShouldEqual, 1)

		v, err = queue.Tail()
		So(err, ShouldBeNil)
		So(v, ShouldEqual, 5)

		v, err = queue.DeQueue()
		So(err, ShouldBeNil)
		So(v, ShouldEqual, 1)

		v, err = queue.DeQueue()
		So(err, ShouldBeNil)
		So(v, ShouldEqual, 2)
		queue.Print()

		err = queue.EnQueue(6)
		So(err, ShouldBeNil)
		err = queue.EnQueue(7)
		So(err, ShouldBeNil)
		queue.Print()

		v, err = queue.DeQueue()
		So(err, ShouldBeNil)
		So(v, ShouldEqual, 3)

		v, err = queue.DeQueue()
		So(err, ShouldBeNil)
		So(v, ShouldEqual, 4)

	})
}
