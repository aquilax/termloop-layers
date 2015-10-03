package layers

import tl "github.com/JoelOtter/termloop"

type Manager struct {
	layers map[string]*Layer
}

func NewManager() *Manager {
	return &Manager{
		layers: make(map[string]*Layer),
	}
}

func (m *Manager) Add(name string, layer *Layer) {
	m.layers[name] = layer
}

func (m *Manager) Remove(name string) {
	delete(m.layers, name)
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
