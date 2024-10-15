package cisco_xr

type ShowBgpInstanceAllSummary struct {
	Instance_name string `json:"INSTANCE_NAME"`
	Local_as      string `json:"LOCAL_AS"`
	Bgp_neigh     string `json:"BGP_NEIGH"`
	Msg_rcvd      string `json:"MSG_RCVD"`
	In_queue      string `json:"IN_QUEUE"`
	Up_down       string `json:"UP_DOWN"`
	State_pfxrcd  string `json:"STATE_PFXRCD"`
	Instance_id   string `json:"INSTANCE_ID"`
	Router_id     string `json:"ROUTER_ID"`
	Neigh_as      string `json:"NEIGH_AS"`
	Msg_sent      string `json:"MSG_SENT"`
	Out_queue     string `json:"OUT_QUEUE"`
}

var ShowBgpInstanceAllSummary_Template string = `Value Filldown INSTANCE_ID (\d+)
Value Filldown INSTANCE_NAME (\S+)
Value Filldown ROUTER_ID (\d+\.\d+\.\d+\.\d+)
Value Filldown LOCAL_AS (\d+)
Value Required BGP_NEIGH (\d+\.\d+\.\d+\.\d+)
Value NEIGH_AS (\d+)
Value MSG_RCVD (\d+)
Value MSG_SENT (\d+)
Value IN_QUEUE (\d+)
Value OUT_QUEUE (\d+)
Value UP_DOWN (\S+)
Value STATE_PFXRCD (\S+)

Start
  ^\s*BGP instance ${INSTANCE_ID}: '${INSTANCE_NAME}'\s*$$
  ^\s*BGP router identifier ${ROUTER_ID}, local AS number ${LOCAL_AS}\s*$$
  ^\s*Neighbor\s+Spk\s+AS\s+MsgRcvd\s+MsgSent\s+TblVer\s+InQ\s+OutQ\s+Up/Down\s+St/PfxRcd\s*$$
  ^\s*${BGP_NEIGH}\s+\d+\s+${NEIGH_AS}\s+${MSG_RCVD}\s+${MSG_SENT}\s+\d+\s+${IN_QUEUE}\s+${OUT_QUEUE}\s+${UP_DOWN}\s+${STATE_PFXRCD}\s*$$ -> Record
`
