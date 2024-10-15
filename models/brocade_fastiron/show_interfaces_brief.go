package brocade_fastiron

type ShowInterfacesBrief struct {
	Duplex    string `json:"duplex"`
	Tagonly   string `json:"tagonly"`
	Linkstate string `json:"linkstate"`
	Portstate string `json:"portstate"`
	Trunkid   string `json:"trunkid"`
	Pvid      string `json:"pvid"`
	Priority  string `json:"priority"`
	Mac       string `json:"mac"`
	Name      string `json:"name"`
	Interface string `json:"interface"`
	Speed     string `json:"speed"`
}

var ShowInterfacesBrief_Template string = `Value interface (\S+)
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
