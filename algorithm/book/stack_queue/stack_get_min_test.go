package stack_queue

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewGetMinStack(t *testing.T) {
	Convey("test get min stack", t, func() {
		ms := NewGetMinStack()
		ms.Push(3)
		ms.Push(4)
		ms.Push(5)
		ms.Push(1)
		ms.Push(2)
		ms.Push(1)

		ret, err := ms.GetMin()
		So(ret, ShouldEqual, 1)
		So(err, ShouldBeNil)

		ret, err = ms.Pop()
		So(ret, ShouldEqual, 1)
		So(err, ShouldBeNil)

		ret, err = ms.GetMin()
		So(ret, ShouldEqual, 1)
		So(err, ShouldBeNil)

		ret, err = ms.Pop()
		So(ret, ShouldEqual, 2)
		So(err, ShouldBeNil)

		ret, err = ms.Pop()
		So(ret, ShouldEqual, 1)
		So(err, ShouldBeNil)

		ret, err = ms.GetMin()
		So(ret, ShouldEqual, 3)
		So(err, ShouldBeNil)

		ret, err = ms.Pop()
		So(ret, ShouldEqual, 5)
		So(err, ShouldBeNil)

		ret, err = ms.Pop()
		So(ret, ShouldEqual, 4)
		So(err, ShouldBeNil)

	})


}
