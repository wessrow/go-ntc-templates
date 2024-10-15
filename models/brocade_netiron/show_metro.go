package brocade_netiron

type ShowMetro struct {
	Port2state     string `json:"port2state"`
	Statechanges   string `json:"statechanges"`
	Ringid         string `json:"ringid"`
	Vlantype       string `json:"vlantype"`
	Ringrole       string `json:"ringrole"`
	Port2role      string `json:"port2role"`
	Rhpssent       string `json:"rhpssent"`
	Tcrbpdusrecvd  string `json:"tcrbpdusrecvd"`
	Ringstate      string `json:"ringstate"`
	Topogroup      string `json:"topogroup"`
	Port1interface string `json:"port1interface"`
	Port2interface string `json:"port2interface"`
	Port1role      string `json:"port1role"`
	Port1state     string `json:"port1state"`
	Port1inttype   string `json:"port1inttype"`
	Port2inttype   string `json:"port2inttype"`
	Ringname       string `json:"ringname"`
	Mastervlan     string `json:"mastervlan"`
	Hellotime      string `json:"hellotime"`
	Prefwingtime   string `json:"prefwingtime"`
	Rhpsrecvd      string `json:"rhpsrecvd"`
}

var ShowMetro_Template string = `Value ringid (\d+)
Value Required vlantype (\w+)
Value ringname ([a-zA-Z0-9\-\"\ ]+)
Value ringstate (enabled|disabled)
Value ringrole (member|master)
Value mastervlan (\d+)
Value topogroup (\d+|not conf)
Value hellotime (\d+)
Value prefwingtime (\d+)
Value port1interface ([0-9\/]+)
Value port1role (primary|secondary)
Value port1state (\w+)
Value port1inttype (\w+)
Value port2interface ([0-9\/]+)
Value port2role (primary|secondary)
Value port2state (\w+)
Value port2inttype (\w+)
Value rhpssent (\d+)
Value rhpsrecvd (\d+)
Value tcrbpdusrecvd (\d+)
Value statechanges (\d+)

Start
  ^Metro Ring ${ringid} - VLAN Type ${vlantype} -> Continue
  ^Metro Ring ${ringid} - VLAN Type ${vlantype} - ${ringname}
  ^\d+\s+${ringstate}\s+${ringrole}\s+${mastervlan}\s+${topogroup}\s+${hellotime}\s+${prefwingtime}
  ^${rhpssent}\s+${rhpsrecvd}\s+${tcrbpdusrecvd}\s+${statechanges} -> Record
  ^ethernet\s+${port1interface}\s+${port1role}\s+${port1state}\s+${port1inttype} -> PORT2

PORT2
  ^ethernet\s+${port2interface}\s+${port2role}\s+${port2state}\s+${port2inttype} -> Start
`
