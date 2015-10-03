package layers

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestLayer(t *testing.T) {
	Convey("Given Layer", t, func() {
		l := NewLayer()

		Convey("Layer is not null", func() {
			So(l, ShouldNotBeNil)
		})
	})
}
