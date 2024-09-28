package models

type CiscoIosShowCryptoSessionDetail struct {
	Interface	string	`json:"INTERFACE"`
	Session_status	string	`json:"SESSION_STATUS"`
	Profile	string	`json:"PROFILE"`
	Uptime	string	`json:"UPTIME"`
	Peer	string	`json:"PEER"`
	Port	string	`json:"PORT"`
	Fvrf	string	`json:"FVRF"`
	Ivrf	string	`json:"IVRF"`
	Description	string	`json:"DESCRIPTION"`
	Phase1_id	string	`json:"PHASE1_ID"`
	Session_id	string	`json:"SESSION_ID"`
	Local_ip	string	`json:"LOCAL_IP"`
	Local_port	string	`json:"LOCAL_PORT"`
	Remote_ip	string	`json:"REMOTE_IP"`
	Remote_port	string	`json:"REMOTE_PORT"`
	Ikev1_status	string	`json:"IKEV1_STATUS"`
	Capabilities	string	`json:"CAPABILITIES"`
	Conn_id	string	`json:"CONN_ID"`
	Lifetime	string	`json:"LIFETIME"`
	Permit	string	`json:"PERMIT"`
	Src_host	string	`json:"SRC_HOST"`
	Src_mask	string	`json:"SRC_MASK"`
	Dst_host	string	`json:"DST_HOST"`
	Dst_mask	string	`json:"DST_MASK"`
	Active_sa	string	`json:"ACTIVE_SA"`
	Origin	string	`json:"ORIGIN"`
}