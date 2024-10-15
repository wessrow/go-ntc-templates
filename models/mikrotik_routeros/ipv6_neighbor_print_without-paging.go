package mikrotik_routeros

type Ipv6NeighborPrintWithoutPaging struct {
	Mac_address  string `json:"MAC_ADDRESS"`
	Status       string `json:"STATUS"`
	Index        string `json:"INDEX"`
	Flags        string `json:"FLAGS"`
	Ipv6_address string `json:"IPV6_ADDRESS"`
	Interface    string `json:"INTERFACE"`
}

var Ipv6NeighborPrintWithoutPaging_Template string = `Value Key INDEX (\d+)
Value FLAGS ([R]+)
Value IPV6_ADDRESS (\S+)
Value INTERFACE (\S+)
Value MAC_ADDRESS ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5})
Value STATUS (noarp|incomplete|stale|reachable|delay|probe|failed)

Start
  ^Flags:.*$$ -> NeighborsTable

NeighborsTable
  ^\s*${INDEX}\s*(${FLAGS}\s+)?address=${IPV6_ADDRESS}\s+interface=${INTERFACE}\s+(mac-address=${MAC_ADDRESS}\s+)?status="${STATUS}"\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`
