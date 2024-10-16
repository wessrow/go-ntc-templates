package cisco_nxos

type ShowIpBgpSummaryVrf struct {
	Address_family string `json:"ADDRESS_FAMILY"`
	Router_id      string `json:"ROUTER_ID"`
	Up_down        string `json:"UP_DOWN"`
	Msg_sent       string `json:"MSG_SENT"`
	In_queue       string `json:"IN_QUEUE"`
	Vrf            string `json:"VRF"`
	Bgp_neigh      string `json:"BGP_NEIGH"`
	Neigh_as       string `json:"NEIGH_AS"`
	Tblver         string `json:"TBLVER"`
	Local_as       string `json:"LOCAL_AS"`
	Bgp_ver        string `json:"BGP_VER"`
	Msg_rcvd       string `json:"MSG_RCVD"`
	Out_queue      string `json:"OUT_QUEUE"`
	State_pfxrcd   string `json:"STATE_PFXRCD"`
}

var ShowIpBgpSummaryVrf_Template string = `Value Filldown VRF (\S+)
Value Filldown ADDRESS_FAMILY (\S+\s\S+)
Value Filldown ROUTER_ID (\d+?\.\d+?\.\d+?\.\d+?)
Value Filldown LOCAL_AS (\d+)
Value Required BGP_NEIGH (\d+?\.\d+?\.\d+?\.\d+?)
Value BGP_VER (\d)
Value Required NEIGH_AS (\S+)
Value MSG_RCVD (\d+)
Value MSG_SENT (\d+)
Value TBLVER (\d+)
Value IN_QUEUE (\d+)
Value OUT_QUEUE (\d+)
Value UP_DOWN (\S+)
Value STATE_PFXRCD (\S+?\s+\S+?|\S+?)

Start
  # Match BGP VRF
  ^BGP summary information for VRF ${VRF}, address family ${ADDRESS_FAMILY}
  # Match RID and Local AS
  ^BGP router identifier ${ROUTER_ID}, local AS number ${LOCAL_AS}
  # Match Neighbor lines
  ^${BGP_NEIGH}\s+${BGP_VER}\s+${NEIGH_AS}.* -> Continue
  ^\s+${MSG_RCVD}\s+${MSG_SENT}\s+${TBLVER}\s+${IN_QUEUE}\s+${OUT_QUEUE}\s+${UP_DOWN}\s+${STATE_PFXRCD}\s*$$ -> Record
  ^${BGP_NEIGH}\s+${BGP_VER}\s+${NEIGH_AS}\s+${MSG_RCVD}\s+${MSG_SENT}\s+${TBLVER}\s+${IN_QUEUE}\s+${OUT_QUEUE}\s+${UP_DOWN}\s+${STATE_PFXRCD}\s*$$ -> Record
  # Match lines that are spaces
  ^\s+$$

Done
  ^.*
`
