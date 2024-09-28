package models

type CiscoNxosShowIpMsdpSummaryVrfAll struct {
	Vrf	string	`json:"VRF"`
	Local_asn	string	`json:"LOCAL_ASN"`
	Originator_id	string	`json:"ORIGINATOR_ID"`
	Peer_address	string	`json:"PEER_ADDRESS"`
	Peer_asn	string	`json:"PEER_ASN"`
	Connection_state	string	`json:"CONNECTION_STATE"`
	Uptime_downtime	string	`json:"UPTIME_DOWNTIME"`
	Last_msg_received	string	`json:"LAST_MSG_RECEIVED"`
	S_g_received	string	`json:"S_G_RECEIVED"`
}