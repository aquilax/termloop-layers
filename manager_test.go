package layers

import (
	tl "github.com/JoelOtter/termloop"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type TestEntity struct {
	DrawCalled int
	TickCalled int
}

func (te *TestEntity) Tick(e tl.Event) {
	te.TickCalled++
}

func (te *TestEntity) Draw(s *tl.Screen) {
	te.DrawCalled++
}

func TestManager(t *testing.T) {
	Convey("Given Layer manager", t, func() {
		m := NewManager()

		Convey("Manager is not null", func() {
			So(m, ShouldNotBeNil)
		})

		Convey("Adding new layer works", func() {
			l := NewLayer()
			m.Add("name", l)
			So(m.Len(), ShouldEqual, 1)
			Convey("Removing layer works", func() {
				m.Remove("name")
				So(m.Len(), ShouldEqual, 0)
			})
		})

		Convey("Adding new test entity works", func() {
			te := &TestEntity{}
			m.Add("name", te)
			So(m.Len(), ShouldEqual, 1)

			Convey("Tick works", func() {
				m.Tick(tl.Event{})
				So(te.TickCalled, ShouldEqual, 1)
			})

			Convey("Draw works", func() {
				m.Draw(nil)
				So(te.DrawCalled, ShouldEqual, 1)
			})

			Convey("Removing entity works", func() {
				m.Remove("name")
				So(m.Len(), ShouldEqual, 0)
			})
		})

	})
}
