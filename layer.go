package layers

import tl "github.com/JoelOtter/termloop"

// Layer contains drawable elements
type Layer struct {
	visible  bool
	Entities []tl.Drawable
}

// NewLayer returns new Layer
func NewLayer() *Layer {
	return &Layer{
		visible:  true,
		Entities: make([]tl.Drawable, 0),
	}
}

// Tick processes termloop Tick
func (l *Layer) Tick(e tl.Event) {
	for i := range l.Entities {
		l.Entities[i].Tick(e)
	}
}

// Draw processes termloop Draw
func (l *Layer) Draw(s *tl.Screen) {
	if l.visible {
		for i := range l.Entities {
			l.Entities[i].Draw(s)
		}
	}
}

// AddEntity adds drawable entity to the layer
func (l *Layer) AddEntity(d tl.Drawable) {
	l.Entities = append(l.Entities, d)
}

// RemoveEntity removes entity from the layer
func (l *Layer) RemoveEntity(d tl.Drawable) {
	for i, elem := range l.Entities {
		if elem == d {
			l.Entities = append(l.Entities[:i], l.Entities[i+1:]...)
			return
		}
	}
}

// Show make s the layer visible
func (l *Layer) Show() {
	l.visible = true
}

// Hide hides the layer
func (l *Layer) Hide() {
	l.visible = false
}

// IsVisible returns the current layer visibility
func (l *Layer) IsVisible() bool {
	return l.visible
}

// Len returns the number of elements on the layer
func (l *Layer) Len() int {
	return len(l.Entities)
}
