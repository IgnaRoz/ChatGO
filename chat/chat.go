package chat

import (
	"fmt"
)

type Mensaje struct {
	From    string
	Mensaje string
}

func (m Mensaje) String() string {
	return fmt.Sprintf("%s: %s", m.From, m.Mensaje)
}

type Sender interface {
	Send(m Mensaje)
}
type Grupo struct {
	Participantes []Sender
}

func (g Grupo) NewMessage(m Mensaje, from Sender) {
	for _, p := range g.Participantes {
		if p == from {
			continue
		}
		p.Send(m)
	}
}
func (g *Grupo) AddParticipant(p Sender) {
	g.Participantes = append(g.Participantes, p)
}

//falta implementar el metodo RemoveParticipant
