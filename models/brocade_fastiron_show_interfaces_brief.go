package models

type BrocadeFastironShowInterfacesBrief struct {
	Interface	string	`json:"interface"`
	Linkstate	string	`json:"linkstate"`
	Portstate	string	`json:"portstate"`
	Duplex	string	`json:"duplex"`
	Speed	string	`json:"speed"`
	Trunkid	string	`json:"trunkid"`
	Tagonly	string	`json:"tagonly"`
	Pvid	string	`json:"pvid"`
	Priority	string	`json:"priority"`
	Mac	string	`json:"mac"`
	Name	string	`json:"name"`
}