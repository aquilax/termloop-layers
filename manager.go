package layers

import tl "github.com/JoelOtter/termloop"

// Manager manages collection of named Drawable objects
type Manager struct {
	keys   []string
	layers map[string]tl.Drawable
}

// NewManager returns new Manager
func NewManager() *Manager {
	return &Manager{
		keys:   make([]string, 0),
		layers: make(map[string]tl.Drawable),
	}
}

// Add adds named Drawable to the manager
func (m *Manager) Add(name string, d tl.Drawable) {
	m.keys = append(m.keys, name)
	m.layers[name] = d
}

// Remove removes named drawable from manager
func (m *Manager) Remove(name string) bool {
	i := m.keyIndex(name)
	if i != -1 {
		m.keys = append(m.keys[:i], m.keys[i+1:]...)
		delete(m.layers, name)
		return true
	}
	return false
}

// Len returns the number of elements in the manager
func (m *Manager) Len() int {
	return len(m.layers)
}

// Tick processes termloop Tick
func (m *Manager) Tick(e tl.Event) {
	for _, key := range m.keys {
		m.layers[key].Tick(e)
	}
}

// Draw processes termloop Draw
func (m *Manager) Draw(s *tl.Screen) {
	for _, key := range m.keys {
		m.layers[key].Draw(s)
	}
}

func (m *Manager) keyIndex(name string) int {
	for i, key := range m.keys {
		if key == name {
			return i
		}
	}
	return -1
}
