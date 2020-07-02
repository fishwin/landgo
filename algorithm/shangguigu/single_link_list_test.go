package shangguigu

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewSingleLinkList(t *testing.T) {
	Convey("test queue", t, func() {
		sl := NewSingleLinkList()
		sl.Print()

		err := sl.DeleteTail()
		So(err, ShouldNotBeNil)

		sl.AddToHead(1)
		err = sl.DeleteHead()
		So(err, ShouldBeNil)

		sl.AddToTail(1)
		err = sl.DeleteTail()
		So(err, ShouldBeNil)

		sl.AddToHead(2)
		sl.AddToTail(5)
		sl.AddAtIndex(1, 4)
		sl.AddAtIndex(1, 3)
		sl.AddToHead(1)
		sl.AddToTail(6)
		sl.Print()

		err = sl.UpdateTail(123)
		So(err, ShouldBeNil)

		err = sl.UpdateHead(222)
		So(err, ShouldBeNil)

		sl.Print()

		err = sl.UpdateTail(6)
		So(err, ShouldBeNil)

		err = sl.UpdateHead(1)
		So(err, ShouldBeNil)

		sl.Print()

		err = sl.UpdateAtIndex(3, 999)
		So(err, ShouldBeNil)

		sl.Print()

		err = sl.DeleteHead()
		So(err, ShouldBeNil)

		err = sl.DeleteTail()
		So(err, ShouldBeNil)

		sl.Print()

		err = sl.DeleteAtIndex(2)
		So(err, ShouldBeNil)

		sl.Print()

		err = sl.DeleteHead()
		So(err, ShouldBeNil)

		sl.Print()

		err = sl.DeleteTail()
		So(err, ShouldBeNil)

		sl.Print()

		err = sl.DeleteTail()
		So(err, ShouldBeNil)

		sl.Print()

		err = sl.DeleteTail()
		So(err, ShouldNotBeNil)

		sl.Print()

		So(sl.Length(), ShouldEqual, 0)

	})
}
