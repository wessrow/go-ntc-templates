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