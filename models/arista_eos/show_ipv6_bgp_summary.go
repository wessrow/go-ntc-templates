package arista_eos 

type ShowIpv6BgpSummary struct {
	Router_id	string	`json:"ROUTER_ID"`
	Local_as	string	`json:"LOCAL_AS"`
	Vrf	string	`json:"VRF"`
	Bgp_neigh	string	`json:"BGP_NEIGH"`
	Neigh_as	string	`json:"NEIGH_AS"`
	Msg_rcvd	string	`json:"MSG_RCVD"`
	Msg_sent	string	`json:"MSG_SENT"`
	In_queue	string	`json:"IN_QUEUE"`
	Out_queue	string	`json:"OUT_QUEUE"`
	Up_down	string	`json:"UP_DOWN"`
	State	string	`json:"STATE"`
	State_pfxrcd	string	`json:"STATE_PFXRCD"`
	State_pfxacc	string	`json:"STATE_PFXACC"`
}

var ShowIpv6BgpSummary_Template = `Value Filldown ROUTER_ID (\d+\.\d+\.\d+\.\d+)
Value Filldown LOCAL_AS (\d+)
Value Filldown VRF (\S+)
Value Required BGP_NEIGH (\S+)
Value NEIGH_AS (\d+)
Value MSG_RCVD (\d+)
Value MSG_SENT (\d+)
Value IN_QUEUE (\d+)
Value OUT_QUEUE (\d+)
Value UP_DOWN (\S+)
Value STATE (\S+)
Value STATE_PFXRCD (\d+)
Value STATE_PFXACC (\d+)

Start
  ^BGP.+?VRF\s+${VRF}\s*$$
  ^\s*Neighbor\s+V\s+AS\s+MsgRcvd\s+MsgSent\s+InQ\s+OutQ\s+Up/Down\s+State/PfxRcd\s*$$
  ^\s*Neighbor\s+V\s+AS\s+MsgRcvd\s+MsgSent\s+InQ\s+OutQ\s+Up/Down\s+State\s+PfxRcd\s+PfxAcc\s*$$
  ^.+\s+${ROUTER_ID},\s+[Ll]ocal\s+[Aa][Ss]\s+[Nn]umber\s+${LOCAL_AS}
  ^\s+${BGP_NEIGH}\s+\d+\s+${NEIGH_AS}\s+${MSG_RCVD}\s+${MSG_SENT}\s+${IN_QUEUE}\s+${OUT_QUEUE}\s+${UP_DOWN}\s+${STATE}\s+${STATE_PFXRCD}\s+${STATE_PFXACC} -> Record
  ^\s+${BGP_NEIGH}\s+\d+\s+${NEIGH_AS}\s+${MSG_RCVD}\s+${MSG_SENT}\s+${IN_QUEUE}\s+${OUT_QUEUE}\s+${UP_DOWN}\s+${STATE}\s* -> Record
  ^${BGP_NEIGH}\s+\d+\s+${NEIGH_AS}\s+${MSG_RCVD}\s+${MSG_SENT}\s+${IN_QUEUE}\s+${OUT_QUEUE}\s+${UP_DOWN}\s+${STATE_PFXRCD} -> Record
  ^Neighbor\s+Status\s+Codes:
  ^\s*$$
  ^. -> Error
`