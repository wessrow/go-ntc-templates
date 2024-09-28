package models

type BrocadeNetironShowMetro struct {
	Ringid	string	`json:"ringid"`
	Vlantype	string	`json:"vlantype"`
	Ringname	string	`json:"ringname"`
	Ringstate	string	`json:"ringstate"`
	Ringrole	string	`json:"ringrole"`
	Mastervlan	string	`json:"mastervlan"`
	Topogroup	string	`json:"topogroup"`
	Hellotime	string	`json:"hellotime"`
	Prefwingtime	string	`json:"prefwingtime"`
	Port1interface	string	`json:"port1interface"`
	Port1role	string	`json:"port1role"`
	Port1state	string	`json:"port1state"`
	Port1inttype	string	`json:"port1inttype"`
	Port2interface	string	`json:"port2interface"`
	Port2role	string	`json:"port2role"`
	Port2state	string	`json:"port2state"`
	Port2inttype	string	`json:"port2inttype"`
	Rhpssent	string	`json:"rhpssent"`
	Rhpsrecvd	string	`json:"rhpsrecvd"`
	Tcrbpdusrecvd	string	`json:"tcrbpdusrecvd"`
	Statechanges	string	`json:"statechanges"`
}