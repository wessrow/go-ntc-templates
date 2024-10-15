package cisco_xr

type ShowIpBgpSummary struct {
	Up_down      string `json:"UP_DOWN"`
	State_pfxrcd string `json:"STATE_PFXRCD"`
	Neigh_as     string `json:"NEIGH_AS"`
	Msg_rcvd     string `json:"MSG_RCVD"`
	In_queue     string `json:"IN_QUEUE"`
	Out_queue    string `json:"OUT_QUEUE"`
	Router_id    string `json:"ROUTER_ID"`
	Local_as     string `json:"LOCAL_AS"`
	Bgp_neigh    string `json:"BGP_NEIGH"`
	Msg_sent     string `json:"MSG_SENT"`
}

var ShowIpBgpSummary_Template string = `Value Filldown ROUTER_ID (\d+\.\d+\.\d+\.\d+)
Value Filldown LOCAL_AS (\d+)
Value BGP_NEIGH (\d+\.\d+\.\d+\.\d+)
Value NEIGH_AS (\d+)
Value MSG_RCVD (\d+)
Value MSG_SENT (\d+)
Value IN_QUEUE (\d+)
Value OUT_QUEUE (\d+)
Value UP_DOWN (\S+)
Value STATE_PFXRCD (\S+)


Start
  ^.*\s+${ROUTER_ID}.*number\s+${LOCAL_AS} -> Begin

Begin
  ^Neighbor\s+.*Spk -> Neighbors
  
Neighbors
  ^${BGP_NEIGH}\s+\d+\s+${NEIGH_AS}\s+${MSG_RCVD}\s+${MSG_SENT}\s+\d+\s+${IN_QUEUE}\s+${OUT_QUEUE}\s+${UP_DOWN}\s+${STATE_PFXRCD} -> Record
  ^\s+$$
  ^$$
  ^.* -> Error "LINE NOT FOUND"
  
EOF
`
