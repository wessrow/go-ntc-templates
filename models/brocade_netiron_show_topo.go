package models

type BrocadeNetironShowTopo struct {
	Topogroup	string	`json:"topogroup"`
	Topohwindex	string	`json:"topohwindex"`
	Mastervlan	string	`json:"mastervlan"`
	L2proto	string	`json:"l2proto"`
	Vplsvlanpresent	string	`json:"vplsvlanpresent"`
	Membervlans	string	`json:"membervlans"`
	Membergroup	string	`json:"membergroup"`
	Controlports	string	`json:"controlports"`
	Freeports	string	`json:"freeports"`
}