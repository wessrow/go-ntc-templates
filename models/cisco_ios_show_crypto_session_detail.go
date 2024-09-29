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

var CiscoIosShowCryptoSessionDetail_Template = `Value Required INTERFACE (\S+)
Value Required SESSION_STATUS (\S+)
Value PROFILE (\S+)
Value UPTIME (\S+)
Value Required PEER (\S+)
Value PORT (\d+)
Value FVRF (\S+)
Value IVRF (\S+)
Value DESCRIPTION (\S+)
Value PHASE1_ID (\S+)
Value SESSION_ID (\d+)
Value LOCAL_IP (\S+)
Value LOCAL_PORT (\d+)
Value REMOTE_IP (\S+)
Value REMOTE_PORT (\S+)
Value IKEV1_STATUS (\S+)
Value CAPABILITIES (\S+)
Value CONN_ID (\d+)
Value LIFETIME (\S+)
Value PERMIT (\S+)
Value SRC_HOST (\S+)
Value SRC_MASK (\S+)
Value DST_HOST (\S+)
Value DST_MASK (\S+)
Value ACTIVE_SA (\d+)
Value ORIGIN (.+)

Start
 ^Crypto\s+.*
 ^Code:
 ^K\s+-
 ^X\s+-
 ^R\s+-
 ^S\s+-
 ^Interface: -> Continue.Record
 ^Interface:\s+${INTERFACE}
 ^Profile:\s+${PROFILE}
 ^Session\s+status:\s+${SESSION_STATUS}
 ^Uptime:\s+${UPTIME}
 ^Peer:\s+${PEER}\s+port\s+${PORT}\s+fvrf:\s+${FVRF}\s+ivrf:\s+${IVRF}
 ^\s+Desc:\s+${DESCRIPTION}
 ^\s+Phase1_id:\s+${PHASE1_ID}
 ^\s+Session\s+ID:\s+${SESSION_ID}
 ^\s+IKEv[1|2]\s+SA:\s+local\s+${LOCAL_IP}/${LOCAL_PORT}\s+remote\s+${REMOTE_IP}/${REMOTE_PORT}\s+${IKEV1_STATUS}
 ^\s+Capabilities:${CAPABILITIES}\s+connid:${CONN_ID}\s+lifetime:${LIFETIME}
 ^\s+IPSEC\s+FLOW:\s+permit\s+${PERMIT}\s+host\s+${SRC_HOST}\s+host\s+${DST_HOST}
 ^\s+IPSEC\s+FLOW:\s+permit\s+${PERMIT}\s+${SRC_HOST}/${SRC_MASK}\s+${DST_HOST}/${DST_MASK}
 ^\s+Active\s+SAs:\s+${ACTIVE_SA},\s+origin:\s+${ORIGIN}
 ^\s+Inbound:\s+#.*
 ^\s+Outbound:\s+#.*
 ^\s*$$
 ^. -> Error

`