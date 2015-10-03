package layers

import tl "github.com/JoelOtter/termloop"

type Layer struct {
	visible  bool
	Entities []tl.Drawable
}

func NewLayer() *Layer {
	return &Layer{
		visible:  true,
		Entities: make([]tl.Drawable, 0),
	}
}

func (l *Layer) Tick(e tl.Event) {
	for i := range l.Entities {
		l.Entities[i].Tick(e)
	}
}

func (l *Layer) Draw(s *tl.Screen) {
	if l.visible {
		for i := range l.Entities {
			l.Entities[i].Draw(s)
		}
	}
}

func (l *Layer) AddEntity(d tl.Drawable) {
	l.Entities = append(l.Entities, d)
}

func (l *Layer) RemoveEntity(d tl.Drawable) {
	for i, elem := range l.Entities {
		if elem == d {
			l.Entities = append(l.Entities[:i], l.Entities[i+1:]...)
			return
		}
	}
}

func (l *Layer) Show() {
	l.visible = false
}

func (l *Layer) Hide() {
	l.visible = false
}

func (l *Layer) IsVisible() bool {
	return l.visible
}
