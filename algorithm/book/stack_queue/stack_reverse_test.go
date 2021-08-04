package stack_queue

import (
	"github.com/fishwin/landgo/algorithm/book/base"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestReverseStack(t *testing.T) {
	Convey("test reverse stack", t, func() {

		s := base.NewStack()
		s.Push(1)
		s.Push(2)
		s.Push(3)
		s.Push(4)
		s.Push(5)

		ReverseStack(s)

		ret, err := s.Pop()
		So(ret, ShouldEqual, 1)
		So(err, ShouldBeNil)

		ret, err = s.Pop()
		So(ret, ShouldEqual, 2)
		So(err, ShouldBeNil)

		ret, err = s.Pop()
		So(ret, ShouldEqual, 3)
		So(err, ShouldBeNil)

		ret, err = s.Pop()
		So(ret, ShouldEqual, 4)
		So(err, ShouldBeNil)

		ret, err = s.Pop()
		So(ret, ShouldEqual, 5)
		So(err, ShouldBeNil)

	})
}
