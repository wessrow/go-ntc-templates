package edgecore 

type ShowVlan struct {
	Vlan_id	string	`json:"VLAN_ID"`
	Type	string	`json:"TYPE"`
	Vlan_name	string	`json:"VLAN_NAME"`
	Status	string	`json:"STATUS"`
	Ports	[]string	`json:"PORTS"`
}

var ShowVlan_Template = `Value VLAN_ID (\d+)
Value TYPE (\S+)
Value VLAN_NAME (.*)
Value STATUS (\S+)
Value List PORTS (\S+\s*\d+\s*/\s*\d+)

Start
  ^\s*Default\s+VLAN\s+ID\s*:\s*\d+\s*$$
  ^\s*VLAN\s+ID:\s*\d+\s*$$ -> Continue.Record
  ^\s*VLAN\s+ID:\s*${VLAN_ID}\s*$$
  ^\s*Type:\s*${TYPE}\s*$$
  ^\s*Name:\s*(?:${VLAN_NAME})?\s*$$
  ^\s*Status:\s*${STATUS}\s*$$
  ^\s*Ports/Port\s+Channels:\s*${PORTS}\s*\(\w+\).*$$ -> Continue
  ^\s*Ports/Port\s+Channels:\s*(\S+\s*\d+\s*/\s*\d+\s*\(\w+\)\s*){1}${PORTS}\s*\(\w+\).*$$ -> Continue
  ^\s*Ports/Port\s+Channels:\s*(\S+\s*\d+\s*/\s*\d+\s*\(\w+\)\s*){2}${PORTS}\s*\(\w+\).*$$ -> Continue
  ^\s*Ports/Port\s+Channels:\s*(\S+\s*\d+\s*/\s*\d+\s*\(\w+\)\s*){3}${PORTS}\s*\(\w+\).*$$ -> Continue
  ^\s*Ports/Port\s+Channels:\s*(\S+\s*\d+\s*/\s*\d+\s*\(\w+\)\s*){4}${PORTS}\s*\(\w+\).*$$ -> Continue
  ^\s*Ports/Port\s+Channels:\s*(\S+\s*\d+\s*/\s*\d+\s*\(\w+\)\s*){5}${PORTS}\s*\(\w+\).*$$ -> Continue
  ^\s*Ports/Port\s+Channels:\s*(\S+\s*\d+\s*/\s*\d+\s*\(\w+\)\s*){6}${PORTS}\s*\(\w+\).*$$ -> Continue
  ^\s*Ports/Port\s+Channels:\s*(\S+\s*\d+\s*/\s*\d+\s*\(\w+\)\s*){7}${PORTS}\s*\(\w+\).*$$ -> Continue
  ^\s*Ports/Port\s+Channels:\s*(\S+\s*\d+\s*/\s*\d+\s*\(\w+\)\s*){8}${PORTS}\s*\(\w+\).*$$ -> Continue
  ^\s*Ports/Port\s+Channels:\s*(\S+\s*\d+\s*/\s*\d+\s*\(\w+\)\s*){9}${PORTS}\s*\(\w+\).*$$ -> Continue
  ^\s*Ports/Port\s+Channels:\s*(\S+\s*\d+\s*/\s*\d+\s*\(\w+\)\s*)*$$
  ^\s*${PORTS}\s*\(\w+\).*$$ -> Continue
  ^\s*(\S+\s*\d+\s*/\s*\d+\s*\(\w+\)\s*){1}${PORTS}\s*\(\w+\).*$$ -> Continue
  ^\s*(\S+\s*\d+\s*/\s*\d+\s*\(\w+\)\s*){2}${PORTS}\s*\(\w+\).*$$ -> Continue
  ^\s*(\S+\s*\d+\s*/\s*\d+\s*\(\w+\)\s*){3}${PORTS}\s*\(\w+\).*$$ -> Continue
  ^\s*(\S+\s*\d+\s*/\s*\d+\s*\(\w+\)\s*){4}${PORTS}\s*\(\w+\).*$$ -> Continue
  ^\s*(\S+\s*\d+\s*/\s*\d+\s*\(\w+\)\s*){5}${PORTS}\s*\(\w+\).*$$ -> Continue
  ^\s*(\S+\s*\d+\s*/\s*\d+\s*\(\w+\)\s*){6}${PORTS}\s*\(\w+\).*$$ -> Continue
  ^\s*(\S+\s*\d+\s*/\s*\d+\s*\(\w+\)\s*){7}${PORTS}\s*\(\w+\).*$$ -> Continue
  ^\s*(\S+\s*\d+\s*/\s*\d+\s*\(\w+\)\s*){8}${PORTS}\s*\(\w+\).*$$ -> Continue
  ^\s*(\S+\s*\d+\s*/\s*\d+\s*\(\w+\)\s*){9}${PORTS}\s*\(\w+\).*$$ -> Continue
  ^\s*(\S+\s*\d+\s*/\s*\d+\s*\(\w+\)\s*)*$$
  ^\s*$$
  ^. -> Error
`