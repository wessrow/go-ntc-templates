package models

type CiscoXrShowPimIpv4GroupMap struct {
	Group_range	string	`json:"GROUP_RANGE"`
	Protocol	string	`json:"PROTOCOL"`
	Client	string	`json:"CLIENT"`
	Groups	string	`json:"GROUPS"`
	Rp_address	string	`json:"RP_ADDRESS"`
	Info	string	`json:"INFO"`
}

var CiscoXrShowPimIpv4GroupMap_Template = `Value Required GROUP_RANGE (\d+\.\d+\.\d+\.\d+/\d+)
Value PROTOCOL (\w+)
Value CLIENT (\w+)
Value GROUPS (\d+)
Value RP_ADDRESS (\d+\.\d+\.\d+\.\d+)
Value INFO ((\S+(\s\S+)*))

Start
  ^\s*Group Range\s+Proto\s+Client\s+Groups\s+RP address\s+Info\s*$$
  ^\s*${GROUP_RANGE}(\*)?\s+${PROTOCOL}\s+${CLIENT}\s+${GROUPS}\s+${RP_ADDRESS}(\s+${INFO})?\s*$$ -> Record
`