package db

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestError(t *testing.T) {
	Convey("Error", t, func() {
		Convey("with cause", func() {
			e := &Error{message: "Foo", cause: errors.New("Bar")}

			So(e.Error(), ShouldEqual, "Foo: Bar")
		})

		Convey("without cause", func() {
			e := &Error{message: "Baz"}

			So(e.Error(), ShouldEqual, "Baz")
		})
	})
}
