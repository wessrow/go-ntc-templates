package vyatta_vyos

type ShowIpBgpSummary struct {
	Msg_rcvd     string `json:"MSG_RCVD"`
	Msg_sent     string `json:"MSG_SENT"`
	Up_down      string `json:"UP_DOWN"`
	State_pfxrcd string `json:"STATE_PFXRCD"`
	Router_id    string `json:"ROUTER_ID"`
	Local_as     string `json:"LOCAL_AS"`
	Bgp_neigh    string `json:"BGP_NEIGH"`
	Neigh_as     string `json:"NEIGH_AS"`
}

var ShowIpBgpSummary_Template string = `Value Filldown ROUTER_ID (\S+)
Value Filldown LOCAL_AS (\d+)
Value BGP_NEIGH (\d+?\.\d+?\.\d+?\.\d+?)
Value NEIGH_AS (\d+)
Value MSG_RCVD (\d+)
Value MSG_SENT (\d+)
Value UP_DOWN (\S+?)
Value STATE_PFXRCD (\S+?\s+\S+?|\S+?)

Start
  ^IPv4\s+Unicast\s+Summary:$$
  ^BGP\s+router\s+identifier\s+${ROUTER_ID},\s+[Ll]ocal\s+[Aa][Ss]\s+number\s+${LOCAL_AS}\s+vrf-id\s+\d+$$
  ^BGP\s+table\s+version\s+\d+$$
  ^RIB\s+entries\s+\d+,\s+using\s+\d+\s+MiB\s+of\s+memory$$
  ^Peers\s+\d+,\s+using\s+\d+\s+KiB\s+of\s+memory$$
  ^Peer\s+groups\s+\d+,\s+using\s+\d+\s+bytes\s+of\s+memory$$
  ^Neighbor\s+V\s+AS\s+MsgRcvd\s+MsgSent\s+TblVer\s+InQ\s+OutQ\s+Up/Down\s+State/PfxRcd$$
  ^${BGP_NEIGH}\s+\S+\s+${NEIGH_AS}\s+${MSG_RCVD}\s+${MSG_SENT}(\s+\d+?){3}\s+${UP_DOWN}\s+${STATE_PFXRCD}\s*$$ -> Record
  ^Total\s+number\s+of\s+neighbors\s+\d+$$
  ^\s*$$
  ^. -> Error

EOF
`
