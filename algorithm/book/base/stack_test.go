package base

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewStack(t *testing.T) {
	Convey("test stack", t, func() {
		s := NewStack()
		So(s.IsEmpty(), ShouldBeTrue)
		ret, err := s.Pop()
		So(ret, ShouldEqual, -1)
		So(err, ShouldEqual, ErrStackEmpty)
		s.Push(1)
		s.Push(2)
		s.Push(3)
		s.Push(4)
		s.Push(5)
		ret, err = s.Pop()
		So(ret, ShouldEqual, 5)
		So(err, ShouldBeNil)
		ret, err = s.Pop()
		So(ret, ShouldEqual, 4)
		So(err, ShouldBeNil)
		ret, err = s.Top()
		So(ret, ShouldEqual, 3)
		So(err, ShouldBeNil)
		So(s.IsEmpty(), ShouldBeFalse)

		s.Push(8)
		s.Push(9)
		ret, err = s.Top()
		So(ret, ShouldEqual, 9)
		So(err, ShouldBeNil)
		ret, err = s.Pop()
		So(ret, ShouldEqual, 9)
		So(err, ShouldBeNil)

		t.Log(s.All())

	})
}
