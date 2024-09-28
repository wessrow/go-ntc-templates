package models

type AristaEosShowMacSecurityParticipantsDetail struct {
	Interface	string	`json:"INTERFACE"`
	Ckn	string	`json:"CKN"`
	Message_id	string	`json:"MESSAGE_ID"`
	Elected_self	string	`json:"ELECTED_SELF"`
	Success	string	`json:"SUCCESS"`
	Principal	string	`json:"PRINCIPAL"`
	Default	string	`json:"DEFAULT"`
	Key_server_sci	string	`json:"KEY_SERVER_SCI"`
	Sak_transmit	string	`json:"SAK_TRANSMIT"`
	Llpn_exhaustion	string	`json:"LLPN_EXHAUSTION"`
	Distributed_key_id	string	`json:"DISTRIBUTED_KEY_ID"`
	Live_peer_list	string	`json:"LIVE_PEER_LIST"`
	Potential_peer_list	string	`json:"POTENTIAL_PEER_LIST"`
}