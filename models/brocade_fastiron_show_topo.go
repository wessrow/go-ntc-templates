package models

type BrocadeFastironShowTopo struct {
	Topogroup	string	`json:"topogroup"`
	Mastervlan	string	`json:"mastervlan"`
	L2proto	string	`json:"l2proto"`
	Membervlans	string	`json:"membervlans"`
	Controlports	string	`json:"controlports"`
	Freeports	string	`json:"freeports"`
}