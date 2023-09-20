package player

type Player struct {
	Name string
	Char string
}

func NewPlayer(name string,char string) *Player {
	return &Player{
		Name:name,
		Char:char,
	}
}

func (player *Player) GetChar() string{
	return player.Char
}

func (player *Player) GetName() string{
	return player.Name
}