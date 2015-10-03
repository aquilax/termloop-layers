package layers

import tl "github.com/JoelOtter/termloop"

type Manager struct {
	layers map[string]tl.Drawable
}

func NewManager() *Manager {
	return &Manager{
		layers: make(map[string]tl.Drawable),
	}
}

func (m *Manager) Add(name string, d tl.Drawable) {
	m.layers[name] = d
}

func (m *Manager) Remove(name string) {
	delete(m.layers, name)
}

func (m *Manager) Len() int {
	return len(m.layers)
}

func (m *Manager) Tick(e tl.Event) {
	for i := range m.layers {
		m.layers[i].Tick(e)
	}
}

func (m *Manager) Draw(s *tl.Screen) {
	for i := range m.layers {
		m.layers[i].Draw(s)
	}
}
