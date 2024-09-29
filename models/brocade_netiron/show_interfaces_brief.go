package brocade_netiron 

type ShowInterfacesBrief struct {
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
}

var ShowInterfacesBrief_Template = `Value interface ([0-9\/velb]+)
Value linkstate (Up|Disabled|Down)
Value portstate ([a-zA-Z\/]+)
Value duplex ([a-zA-Z\/]+)
Value speed ([a-zA-Z0-9\/]+)
Value trunkid (\d+|None)
Value tagonly (Yes|No|N\/A)
Value pvid (\d+|N\/A)
Value priority ([a-zA-Z0-9\/]+)
Value mac (\w+\.\w+\.\w+|N\/A)

Start
  ^${interface}\s+${linkstate}\s+${portstate}\s+${duplex}\s+${speed}\s+${trunkid}\s+${tagonly}\s+${priority}\s+${mac} -> Next.Record
`