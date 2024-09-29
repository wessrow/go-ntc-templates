package models

type FortinetExecuteTraceroute struct {
	Destination_address	string	`json:"DESTINATION_ADDRESS"`
	Hops_max	string	`json:"HOPS_MAX"`
	Packets_per_hop	string	`json:"PACKETS_PER_HOP"`
	Packet_size	string	`json:"PACKET_SIZE"`
	Hop_num	string	`json:"HOP_NUM"`
	Address	string	`json:"ADDRESS"`
	Details	string	`json:"DETAILS"`
}

var FortinetExecuteTraceroute_Template = `Value DESTINATION_ADDRESS (\S+)
Value HOPS_MAX (\d+)
Value PACKETS_PER_HOP (\d+)
Value PACKET_SIZE (\d+)
Value HOP_NUM (\d+)
Value ADDRESS ((?!\*(?:\s+\*)*)\S+)
Value DETAILS (.*)

Start
  ^\s*traceroute\s+to\s+${DESTINATION_ADDRESS}\s+\(.*?\),\s+${HOPS_MAX}\s+hops\s+max,\s+${PACKETS_PER_HOP}\s+probe\s+packets\s+per\s+hop,\s+${PACKET_SIZE}\s+byte\s+packets\s*$$
  ^\s*${HOP_NUM}(?:\s+${ADDRESS})?\s+${DETAILS}\s*$$ -> Record
  ^\s*\*\s*$$
  ^\s*$$
  ^. -> Error

`