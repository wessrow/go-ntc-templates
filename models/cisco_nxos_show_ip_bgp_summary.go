package models

type CiscoNxosShowIpBgpSummary struct {
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
	State_pfxrcd	string	`json:"STATE_PFXRCD"`
}

var CiscoNxosShowIpBgpSummary_Template = `Value Filldown ROUTER_ID (\S+)
Value Filldown LOCAL_AS (\d+)
Value Filldown VRF ([A-Za-z0-9\-_]+)
Value Required BGP_NEIGH (\d+?\.\d+?\.\d+?\.\d+?)
Value NEIGH_AS (\d+)
Value MSG_RCVD (\d+)
Value MSG_SENT (\d+)
Value IN_QUEUE (\d+)
Value OUT_QUEUE (\d+)
Value UP_DOWN (\S+)
Value STATE_PFXRCD (\S+?\s+\S+?|\S+?)

Start
  ^BGP.+VRF\s+${VRF},?
  ^BGP\s+router\s+identifier\s+${ROUTER_ID},\s+[Ll]ocal\s+[Aa][Ss]\s+number\s+${LOCAL_AS}
  ^Neighbor\s+V\s+AS\s+MsgRcvd\s+MsgSent\s+TblVer\s+InQ\s+OutQ\s+Up/Down\s+State/PfxRcd\s*$$
  ^${BGP_NEIGH}\s+\d+\s+${NEIGH_AS}\s+${MSG_RCVD}\s+${MSG_SENT}\s+\d+\s+${IN_QUEUE}\s+${OUT_QUEUE}\s+${UP_DOWN}\s+${STATE_PFXRCD}\s*$$ -> Record
  ^${BGP_NEIGH}\s+\d+\s+${NEIGH_AS}
  ^\s+${MSG_RCVD}\s+${MSG_SENT}\s+\d+\s+${IN_QUEUE}\s+${OUT_QUEUE}\s+${UP_DOWN}\s+${STATE_PFXRCD}\s*$$ -> Record
  ^BGP\s+table\s+version
  ^\d+\s+network\s+entries
  ^BGP attribute entries
  ^BGP community entries
  ^\d+\s+(received|identical)
  ^\s*$$
  ^. -> Error

`