package models

type CiscoIosShowIpEigrpInterfacesDetail struct {
	As	string	`json:"AS"`
	Interface	string	`json:"INTERFACE"`
	Peers	string	`json:"PEERS"`
	Xmit_q_unreliable	string	`json:"XMIT_Q_UNRELIABLE"`
	Xmit_q_reliable	string	`json:"XMIT_Q_RELIABLE"`
	Peer_q_unreliable	string	`json:"PEER_Q_UNRELIABLE"`
	Peer_q_reliable	string	`json:"PEER_Q_RELIABLE"`
	Mean_srtt	string	`json:"MEAN_SRTT"`
	Pacing_time_unreliable	string	`json:"PACING_TIME_UNRELIABLE"`
	Pacing_time_reliable	string	`json:"PACING_TIME_RELIABLE"`
	Mcast_flow_timer	string	`json:"MCAST_FLOW_TIMER"`
	Pending_routes	string	`json:"PENDING_ROUTES"`
	Hello_interval	string	`json:"HELLO_INTERVAL"`
	Hold_time	string	`json:"HOLD_TIME"`
	Split_horizon	string	`json:"SPLIT_HORIZON"`
	Next_xmit_serial	string	`json:"NEXT_XMIT_SERIAL"`
	Packetized_sent	string	`json:"PACKETIZED_SENT"`
	Packetized_expedited	string	`json:"PACKETIZED_EXPEDITED"`
	Hellos_sent	string	`json:"HELLOS_SENT"`
	Hellos_expedited	string	`json:"HELLOS_EXPEDITED"`
	Mcasts_unreliable	string	`json:"MCASTS_UNRELIABLE"`
	Mcasts_reliable	string	`json:"MCASTS_RELIABLE"`
	Ucasts_unreliable	string	`json:"UCASTS_UNRELIABLE"`
	Ucasts_reliable	string	`json:"UCASTS_RELIABLE"`
	Mcast_exceptions	string	`json:"MCAST_EXCEPTIONS"`
	Cr_packets	string	`json:"CR_PACKETS"`
	Acks_suppressed	string	`json:"ACKS_SUPPRESSED"`
	Retransmissions_sent	string	`json:"RETRANSMISSIONS_SENT"`
	Out_of_sequence_rcvd	string	`json:"OUT_OF_SEQUENCE_RCVD"`
	Topology_ids	string	`json:"TOPOLOGY_IDS"`
	Authentication_mode	string	`json:"AUTHENTICATION_MODE"`
	Topologies_advertised	string	`json:"TOPOLOGIES_ADVERTISED"`
	Topologies_not_advertised	string	`json:"TOPOLOGIES_NOT_ADVERTISED"`
}

var CiscoIosShowIpEigrpInterfacesDetail_Template = `Value Filldown AS (\d+)
Value Required INTERFACE (\S+)
Value PEERS (\d+)
Value XMIT_Q_UNRELIABLE (\d+)
Value XMIT_Q_RELIABLE (\d+)
Value PEER_Q_UNRELIABLE (\d+)
Value PEER_Q_RELIABLE (\d+)
Value MEAN_SRTT (\d+)
Value PACING_TIME_UNRELIABLE (\d+)
Value PACING_TIME_RELIABLE (\d+)
Value MCAST_FLOW_TIMER (\d+)
Value PENDING_ROUTES (\d+)
Value HELLO_INTERVAL (\d+)
Value HOLD_TIME (\d+)
Value SPLIT_HORIZON (\S+)
Value NEXT_XMIT_SERIAL (\S+)
Value PACKETIZED_SENT (\d+)
Value PACKETIZED_EXPEDITED (\d+)
Value HELLOS_SENT (\d+)
Value HELLOS_EXPEDITED (\d+)
Value MCASTS_UNRELIABLE (\d+)
Value MCASTS_RELIABLE (\d+)
Value UCASTS_UNRELIABLE (\d+)
Value UCASTS_RELIABLE (\d+)
Value MCAST_EXCEPTIONS (\d+)
Value CR_PACKETS (\d+)
Value ACKS_SUPPRESSED (\d+)
Value RETRANSMISSIONS_SENT (\d+)
Value OUT_OF_SEQUENCE_RCVD (\d+)
Value TOPOLOGY_IDS (\d+)
Value AUTHENTICATION_MODE (.+)
Value TOPOLOGIES_ADVERTISED (\S*)
Value TOPOLOGIES_NOT_ADVERTISED (\S*)

Start
  ^\s*EIGRP-IPv\d\s+Interfaces\s+for\sAS\(${AS}\)\s*$$
  ^\s*Xmit\sQueue\s+PeerQ\s+Mean\s+Pacing\s+Time\s+Multicast\s+Pending\s*$$
  ^\s*Interface\s+Peers\s+Un/Reliable\s+Un/Reliable\s+SRTT\s+Un/Reliable\s+Flow\s+Timer\s+Routes\s*$$
  ^\s*\S+\s+\d+\s+\d+/\d+\s+\d+/\d+\s+\d+\s+\d+/\d+\s+\d+\s+\d+\s*$$ -> Continue.Record
  ^\s*${INTERFACE}\s+${PEERS}\s+${XMIT_Q_UNRELIABLE}/${XMIT_Q_RELIABLE}\s+${PEER_Q_UNRELIABLE}/${PEER_Q_RELIABLE}\s+${MEAN_SRTT}\s+${PACING_TIME_UNRELIABLE}/${PACING_TIME_RELIABLE}\s+${MCAST_FLOW_TIMER}\s+${PENDING_ROUTES}\s*$$
  ^\s*Hello-interval\s+is\s+${HELLO_INTERVAL},\s+Hold-time\s+is\s+${HOLD_TIME}\s*$$
  ^\s*Split-horizon\s+is\s+${SPLIT_HORIZON}\s*$$
  ^\s*Next\s+xmit\s+serial\s+${NEXT_XMIT_SERIAL}\s*$$
  ^\s*Packetized\s+sent/expedited:\s+${PACKETIZED_SENT}/${PACKETIZED_EXPEDITED}\s*$$
  ^\s*Hello's\s+sent/expedited:\s+${HELLOS_SENT}/${HELLOS_EXPEDITED}\s*$$
  ^\s*Un/reliable\s+mcasts:\s+${MCASTS_UNRELIABLE}/${MCASTS_RELIABLE}\s+Un/reliable\s+ucasts:\s+${UCASTS_UNRELIABLE}/${UCASTS_RELIABLE}\s*$$
  ^\s*Mcast\s+exceptions:\s+${MCAST_EXCEPTIONS}\s+CR\s+packets:\s+${CR_PACKETS}\s+ACKs\s+suppressed:\s+${ACKS_SUPPRESSED}\s*$$
  ^\s*Retransmissions\s+sent:\s+${RETRANSMISSIONS_SENT}\s+Out-of-sequence\s+rcvd:\s+${OUT_OF_SEQUENCE_RCVD}\s*$$
  ^\s*Topology-ids\s+on\s+interface\s+-\s+${TOPOLOGY_IDS}\s*$$
  ^\s*Authentication\s+mode\s+is\s+${AUTHENTICATION_MODE}\s*$$
  ^\s*Topologies\s+advertised\s+on\s+this\s+interface:\s*${TOPOLOGIES_ADVERTISED}\s*$$
  ^\s*Topologies\s+not\s+advertised\s+on\s+this\s+interface:\s*${TOPOLOGIES_NOT_ADVERTISED}\s*$$
  ^\s*$$
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is
  ^. -> Error

`