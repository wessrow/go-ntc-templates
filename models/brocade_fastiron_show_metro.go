package models

type BrocadeFastironShowMetro struct {
	Ringid	string	`json:"ringid"`
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
	Port1activeint	string	`json:"port1activeint"`
	Port1inttype	string	`json:"port1inttype"`
	Port2interface	string	`json:"port2interface"`
	Port2role	string	`json:"port2role"`
	Port2state	string	`json:"port2state"`
	Port2activeint	string	`json:"port2activeint"`
	Port2inttype	string	`json:"port2inttype"`
	Rhpssent	string	`json:"rhpssent"`
	Rhpsrecvd	string	`json:"rhpsrecvd"`
	Tcrbpdusrecvd	string	`json:"tcrbpdusrecvd"`
	Statechanges	string	`json:"statechanges"`
}

var BrocadeFastironShowMetro_Template = `Value ringid (\d+)
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