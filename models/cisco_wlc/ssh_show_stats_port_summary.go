package cisco_wlc 

type SshShowStatsPortSummary struct {
	Port	string	`json:"PORT"`
	Link_status	string	`json:"LINK_STATUS"`
	Packets_in	string	`json:"PACKETS_IN"`
	Packets_out	string	`json:"PACKETS_OUT"`
	Broadcast_in	string	`json:"BROADCAST_IN"`
	Errors_in	string	`json:"ERRORS_IN"`
	Errors_out	string	`json:"ERRORS_OUT"`
	Collisions	string	`json:"COLLISIONS"`
	Total_packets_in	string	`json:"TOTAL_PACKETS_IN"`
	Total_packets_out	string	`json:"TOTAL_PACKETS_OUT"`
	Total_broadcast_in	string	`json:"TOTAL_BROADCAST_IN"`
	Total_errors_in	string	`json:"TOTAL_ERRORS_IN"`
	Total_errors_out	string	`json:"TOTAL_ERRORS_OUT"`
	Total_collisions	string	`json:"TOTAL_COLLISIONS"`
}

var SshShowStatsPortSummary_Template = `Value Required PORT (\d)
Value LINK_STATUS (\S+)
Value PACKETS_IN (\d+)
Value PACKETS_OUT (\d+)
Value BROADCAST_IN (\d+)
Value ERRORS_IN (\d+)
Value ERRORS_OUT (\d+)
Value COLLISIONS (\d+)
Value Fillup TOTAL_PACKETS_IN (\d+)
Value Fillup TOTAL_PACKETS_OUT (\d+)
Value Fillup TOTAL_BROADCAST_IN (\d+)
Value Fillup TOTAL_ERRORS_IN (\d+)
Value Fillup TOTAL_ERRORS_OUT (\d+)
Value Fillup TOTAL_COLLISIONS (\d+)

Start
  ^\s+Link\s+(Pkts In\s+){2}\s*Pkts Out\s*$$
  ^\s*Pr\s+Status\s+(Pkts (In|Out)\s+){2}\s*Bcast\s+(?:Errors\s+){2}\s*Collisions\s*$$ -> Port_Stats
  #
  # Port Statistics (after the dashed line)
  #^\s*[-\s]+$$ -> Port_Stats
  ^\s*$$
  ^. -> Error

Port_Stats
  ^-+
  ^\s*${PORT}\s+${LINK_STATUS}\s+${PACKETS_IN}\s+${PACKETS_OUT}\s+${BROADCAST_IN}\s+${ERRORS_IN}\s+${ERRORS_OUT}\s+${COLLISIONS}\s*$$ -> Record
  ^\s*Total\s+${TOTAL_PACKETS_IN}\s+${TOTAL_PACKETS_OUT}\s+${TOTAL_BROADCAST_IN}\s+${TOTAL_ERRORS_IN}\s+${TOTAL_ERRORS_OUT}\s+${TOTAL_COLLISIONS}\s*$$
  ^\s*$$
  ^. -> Error
`