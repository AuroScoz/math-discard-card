package game

type Player struct {
	Pt int
}

var MyPlayer *Player

func NewPlayer(pt int) {
	MyPlayer = &Player{
		Pt: 100,
	}
}

func (p *Player) AddPt(value int) {
	p.Pt += value
}
