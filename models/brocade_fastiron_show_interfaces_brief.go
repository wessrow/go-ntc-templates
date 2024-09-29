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

var BrocadeFastironShowInterfacesBrief_Template = `Value interface (\S+)
Value linkstate (Up|Disable|Down|ERR-DIS)
Value portstate ([a-zA-Z\/]+)
Value duplex ([a-zA-Z\/]+)
Value speed ([a-zA-Z0-9\/]+)
Value trunkid (\d+|None)
Value tagonly (Yes|No|N\/A)
Value pvid (\d+|N\/A|None)
Value priority ([a-zA-Z0-9\/]+)
Value mac ([A-Fa-f0-9\.]{14})
Value name (\S+)

Start
  ^Port\s+Link\s+State\s+Dupl\s+Speed\s+Trunk\s+Tag\s+Pvid\s+Pri\s+MAC\s+Name
  ^${interface}\s+${linkstate}\s+${portstate}\s+${duplex}\s+${speed}\s+${trunkid}\s+${tagonly}\s+${pvid}\s+${priority}\s+${mac}\s+${name} -> Next.Record
  ^${interface}\s+${linkstate}\s+${portstate}\s+${duplex}\s+${speed}\s+${trunkid}\s+${tagonly}\s+${pvid}\s+${priority}\s+${mac} -> Next.Record
  ^. -> Error

`