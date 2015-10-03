package layers

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestManager(t *testing.T) {
	Convey("Given Layer manager", t, func() {
		m := NewManager()

		Convey("Manager is not null", func() {
			So(m, ShouldNotBeNil)
		})
	})
}
