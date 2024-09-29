package arista_eos 

type ShowMacSecurityParticipantsDetail struct {
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

var ShowMacSecurityParticipantsDetail_Template = `Value Filldown INTERFACE (\S+)
Value CKN (\S+)
Value MESSAGE_ID (\S+)
Value ELECTED_SELF (\S+)
Value SUCCESS (\S+)
Value PRINCIPAL (\S+)
Value DEFAULT (\S+)
Value KEY_SERVER_SCI (\S+)
Value SAK_TRANSMIT (\S+)
Value LLPN_EXHAUSTION (\S+)
Value DISTRIBUTED_KEY_ID (\S+)
Value LIVE_PEER_LIST (\S+)
Value POTENTIAL_PEER_LIST (\S+)

Start
  ^Interface:\s+${INTERFACE} -> Ckn1

Ckn1
  ^\s+CKN:\s+${CKN}
  ^\s+Message ID:\s+${MESSAGE_ID}
  ^\s+Elected self:\s+${ELECTED_SELF}
  ^\s+Success:\s+${SUCCESS}
  ^\s+Principal:\s+${PRINCIPAL}
  ^\s+Default:\s+${DEFAULT}
  ^\s+KeyServer SCI:\s+${KEY_SERVER_SCI}
  ^\s+SAK transmit:\s+${SAK_TRANSMIT}
  ^\s+LLPN exhaustion:\s+${LLPN_EXHAUSTION}
  ^\s+Distributed key identifier:\s+${DISTRIBUTED_KEY_ID}
  ^\s+Live peer list:\s+${LIVE_PEER_LIST}
  ^\s+Potential peer list:\s+${POTENTIAL_PEER_LIST} -> Record
  ^\s+$$
  ^$$ -> Ckn2
  ^.* -> Error "LINE NOT FOUND"
 
Ckn2
  ^\s+CKN:\s+${CKN}
  ^\s+Message ID:\s+${MESSAGE_ID}
  ^\s+Elected self:\s+${ELECTED_SELF}
  ^\s+Success:\s+${SUCCESS}
  ^\s+Principal:\s+${PRINCIPAL}
  ^\s+Default:\s+${DEFAULT}
  ^\s+KeyServer SCI:\s+${KEY_SERVER_SCI}
  ^\s+SAK transmit:\s+${SAK_TRANSMIT}
  ^\s+LLPN exhaustion:\s+${LLPN_EXHAUSTION}
  ^\s+Distributed key identifier:\s+${DISTRIBUTED_KEY_ID}
  ^\s+Live peer list:\s+${LIVE_PEER_LIST}
  ^\s+Potential peer list:\s+${POTENTIAL_PEER_LIST} -> Record
  ^\s+$$
  ^$$ -> Start
  ^.* -> Error "LINE NOT FOUND"
  
EOF`