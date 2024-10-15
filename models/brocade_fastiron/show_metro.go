package brocade_fastiron

type ShowMetro struct {
	Prefwingtime   string `json:"prefwingtime"`
	Port1interface string `json:"port1interface"`
	Port2state     string `json:"port2state"`
	Rhpsrecvd      string `json:"rhpsrecvd"`
	Ringid         string `json:"ringid"`
	Ringstate      string `json:"ringstate"`
	Mastervlan     string `json:"mastervlan"`
	Port2inttype   string `json:"port2inttype"`
	Hellotime      string `json:"hellotime"`
	Port1inttype   string `json:"port1inttype"`
	Port2role      string `json:"port2role"`
	Tcrbpdusrecvd  string `json:"tcrbpdusrecvd"`
	Topogroup      string `json:"topogroup"`
	Port1role      string `json:"port1role"`
	Port2interface string `json:"port2interface"`
	Port1activeint string `json:"port1activeint"`
	Port2activeint string `json:"port2activeint"`
	Rhpssent       string `json:"rhpssent"`
	Statechanges   string `json:"statechanges"`
	Ringname       string `json:"ringname"`
	Ringrole       string `json:"ringrole"`
	Port1state     string `json:"port1state"`
}

var ShowMetro_Template string = `Value ringid (\d+)
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
Value port1activeint ([0-9\/]+)
Value port1inttype (\w+)
Value port2interface ([0-9\/]+)
Value port2role (primary|secondary)
Value port2state (\w+)
Value port2activeint ([0-9\/]+)
Value port2inttype (\w+)
Value rhpssent (\d+)
Value rhpsrecvd (\d+)
Value tcrbpdusrecvd (\d+)
Value statechanges (\d+)

Start
  ^Metro Ring ${ringid} -> Continue
  ^Metro Ring ${ringid} - ${ringname}
  ^\d+\s+${ringstate}\s+${ringrole}\s+${mastervlan}\s+${topogroup}\s+${hellotime}\s+${prefwingtime}
  ^${rhpssent}\s+${rhpsrecvd}\s+${tcrbpdusrecvd}\s+${statechanges} -> Record
  ^ethernet\s+${port1interface}\s+${port1role}\s+${port1state}\s+ethernet\s+${port1activeint}\s+${port1inttype} -> PORT2

PORT2
  ^ethernet\s+${port2interface}\s+${port2role}\s+${port2state}\s+ethernet\s+${port2activeint}\s+${port2inttype} -> Start

`
