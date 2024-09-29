package models

type DellForce10ShowIpInterfaceBrief struct {
	Interface	string	`json:"INTERFACE"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Ok	string	`json:"OK"`
	Method	string	`json:"METHOD"`
	Status	string	`json:"STATUS"`
	Proto	string	`json:"PROTO"`
}

var DellForce10ShowIpInterfaceBrief_Template = `Value INTERFACE (\S+ \S+)
Value IP_ADDRESS (\S+)
Value OK (YES|NO)
Value METHOD (None|Manual)
Value STATUS (up|down)
Value PROTO (up|down)

Start
  ^${INTERFACE}\s+${IP_ADDRESS}\s+${OK}\s+${METHOD}\s+${STATUS}\s+${PROTO} -> Record

`