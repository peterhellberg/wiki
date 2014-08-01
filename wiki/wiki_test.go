package wiki

import (
	"testing"

	"github.com/peterhellberg/wiki/db"

	. "github.com/smartystreets/goconvey/convey"
)

func TestWiki(t *testing.T) {
	Convey("Wiki", t, func() {
		Convey("NewWiki", func() {
			db := &db.DB{}

			w, err := NewWiki(db)

			So(err, ShouldBeNil)
			So(w.DB(), ShouldEqual, db)
		})

		Convey("getPageName", func() {
			w := &Wiki{}

			// Default page name as a slice of bytes
			So(w.getPageName(""), ShouldResemble, []byte("root"))

			// Return page name as a slice of bytes
			So(w.getPageName("foo"), ShouldResemble, []byte("foo"))
		})
	})
}
