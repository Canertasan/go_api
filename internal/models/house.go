package models

type House struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Rooms  int    `json:"rooms,omitempty"`
	Square int    `json:"square,omitempty"`
}
