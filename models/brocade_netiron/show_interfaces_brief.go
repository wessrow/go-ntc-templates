package brocade_netiron

type ShowInterfacesBrief struct {
	Portstate string `json:"portstate"`
	Duplex    string `json:"duplex"`
	Tagonly   string `json:"tagonly"`
	Pvid      string `json:"pvid"`
	Mac       string `json:"mac"`
	Interface string `json:"interface"`
	Linkstate string `json:"linkstate"`
	Speed     string `json:"speed"`
	Trunkid   string `json:"trunkid"`
	Priority  string `json:"priority"`
}

var ShowInterfacesBrief_Template string = `Value interface ([0-9\/velb]+)
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
