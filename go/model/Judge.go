package model
type Judge interface{
	Check(play[]*Player)
}