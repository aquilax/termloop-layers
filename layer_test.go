package layers

import (
	tl "github.com/JoelOtter/termloop"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type TestEntity struct{}

func (te *TestEntity) Tick(e tl.Event) {}

func (te *TestEntity) Draw(s *tl.Screen) {}

func TestLayer(t *testing.T) {
	Convey("Given Layer", t, func() {
		l := NewLayer()

		Convey("Layer is not null", func() {
			So(l, ShouldNotBeNil)
		})

		Convey("Layer is hidden", func() {
			l.Hide()
			Convey("Layer is not visible", func() {
				So(l.IsVisible(), ShouldBeFalse)
				Convey("Layer is Shown", func() {
					l.Show()
					Convey("Layer is visible", func() {
						So(l.IsVisible(), ShouldBeTrue)
					})
				})
			})
		})
		Convey("Given test entity", func() {
			en := &TestEntity{}
			Convey("Adding entity to the layer works", func() {
				l.AddEntity(en)
				So(len(l.Entities), ShouldEqual, 1)
				Convey("Removing  entity from the layer works", func() {
					l.RemoveEntity(en)
					So(len(l.Entities), ShouldEqual, 0)
				})
			})
		})
	})
}
