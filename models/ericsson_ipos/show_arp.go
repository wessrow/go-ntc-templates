package ericsson_ipos

type ShowArp struct {
	Context     string `json:"CONTEXT"`
	Context_id  string `json:"CONTEXT_ID"`
	Host        string `json:"HOST"`
	Mac_address string `json:"MAC_ADDRESS"`
	Ttl         string `json:"TTL"`
	Type        string `json:"TYPE"`
	Circuit     string `json:"CIRCUIT"`
}

var ShowArp_Template string = `Value Required HOST (\d+\.\d+\.\d+\.\d+)
Value Required MAC_ADDRESS (\S+)
Value Required TTL (\S+)
Value Required TYPE (\S+)
Value Required CIRCUIT (\S+\s+\S+\s+\d+|\S+)
Value Filldown CONTEXT (\S+)
Value Filldown CONTEXT_ID (\S+)

Start
  ^Context\s+:${CONTEXT}\s+Context\s+id\s+:\s+${CONTEXT_ID}
  ^-+
  ^Total\s+number\s+of\s+arp\s+entries\s+in\s+cache:\s+\d+
  ^\s+Resolved\s+entry\s+:\s+\d+
  ^\s+Incomplete\s+entry\s+:\s+\d+
  ^Host\s+Hardware\s+address\s+Ttl\s+Type\s+Circuit
  ^${HOST}\s+${MAC_ADDRESS}\s+${TTL}\s+${TYPE}\s+${CIRCUIT} -> Record
  ^Showing\s+ARP\s+entries\s+in\s+RP\s+OS\s+kernel:
  ^Host\s+Hardware\s+address\s+Flags\s+Type
  ^. -> Error
`
