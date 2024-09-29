package cisco_nxos 

type ShowIpMsdpSummaryVrfAll struct {
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

var ShowIpMsdpSummaryVrfAll_Template = `Value Filldown VRF (\S+)
Value Filldown LOCAL_ASN (\d+)
Value Filldown ORIGINATOR_ID (\d+\.\d+\.\d+\.\d+)
Value Required PEER_ADDRESS (\d+\.\d+\.\d+\.\d+)
Value Required PEER_ASN (\d+)
Value Required CONNECTION_STATE (\S+)
Value Required UPTIME_DOWNTIME (\S+)
Value Required LAST_MSG_RECEIVED (\S+)
Value Required S_G_RECEIVED (\S+)

Start
  ^\s*MSDP Peer Status Summary for VRF "${VRF}"\s*$$
  ^\s*Local ASN: ${LOCAL_ASN}, originator-id: ${ORIGINATOR_ID}\s*$$
  ^\s*Number.+\s*$$
  ^\s*Peer.+\s*$$
  ^\s*Address.+\s*$$
  ^\s*${PEER_ADDRESS}\s+${PEER_ASN}\s+${CONNECTION_STATE}\s+${UPTIME_DOWNTIME}\s+${LAST_MSG_RECEIVED}\s+${S_G_RECEIVED} -> Record  
  ^. -> Error`