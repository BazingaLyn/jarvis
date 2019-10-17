package model

type Movie struct {
	Id        string
	Name      string
	Score     float32
	Actors    []string
	Directors []string
	MovieTime int16
}
