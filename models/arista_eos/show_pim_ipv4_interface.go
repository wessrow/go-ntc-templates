package arista_eos

type ShowPimIpv4Interface struct {
	Mode            string `json:"MODE"`
	Hello_interval  string `json:"HELLO_INTERVAL"`
	Dr_address      string `json:"DR_ADDRESS"`
	Interface       string `json:"INTERFACE"`
	Neighbors       string `json:"NEIGHBORS"`
	Dr_priority     string `json:"DR_PRIORITY"`
	Packets_queued  string `json:"PACKETS_QUEUED"`
	Packets_dropped string `json:"PACKETS_DROPPED"`
	Ip_address      string `json:"IP_ADDRESS"`
}

var ShowPimIpv4Interface_Template string = `Value IP_ADDRESS (\d+\.\d+\.\d+\.\d+)
Value Required INTERFACE ([\w\./-]+)
Value MODE (\S+)
Value NEIGHBORS (\d+)
Value HELLO_INTERVAL (\d+)
Value DR_PRIORITY (\d+)
Value DR_ADDRESS (\d+\.\d+\.\d+\.\d+)
Value PACKETS_QUEUED (\d+)
Value PACKETS_DROPPED (\d+)

Start
  ^\s*Address\s+Interface\s+Mode\s+Neighbor\s+Hello\s+DR\s+DR Address\s+PktsQed\s+PktsDropped\s*$$
  ^\s*Count\s+Intvl\s+Pri\s*$$
  ^\s*${IP_ADDRESS}\s+${INTERFACE}\s+${MODE}\s+${NEIGHBORS}\s+${HELLO_INTERVAL}\s+${DR_PRIORITY}\s+${DR_ADDRESS}\s+${PACKETS_QUEUED}\s+${PACKETS_DROPPED}\s*$$ -> Record
`
