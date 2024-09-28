package models

type BrocadeNetironShowLagBrief struct {
	Lagnameshort	string	`json:"lagnameshort"`
	Lagtype	string	`json:"lagtype"`
	Deployed	string	`json:"deployed"`
	Trunkid	string	`json:"trunkid"`
	Primaryinterface	string	`json:"primaryinterface"`
	Portlist	string	`json:"portlist"`
	Portlistcont	string	`json:"portlistcont"`
}