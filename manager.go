package layers

import tl "github.com/JoelOtter/termloop"

// Manager manages collection of named Drawable objects
type Manager struct {
	layers map[string]tl.Drawable
}

// NewManager returns new Manager
func NewManager() *Manager {
	return &Manager{
		layers: make(map[string]tl.Drawable),
	}
}

// Add adds named Drawable to the manager
func (m *Manager) Add(name string, d tl.Drawable) {
	m.layers[name] = d
}

// Remove removes named drawable from manager
func (m *Manager) Remove(name string) {
	delete(m.layers, name)
}

// Len returns the number of elements in the manager
func (m *Manager) Len() int {
	return len(m.layers)
}

// Tick processes termloop Tick
func (m *Manager) Tick(e tl.Event) {
	for i := range m.layers {
		m.layers[i].Tick(e)
	}
}

// Draw processes termloop Draw
func (m *Manager) Draw(s *tl.Screen) {
	for i := range m.layers {
		m.layers[i].Draw(s)
	}
}
