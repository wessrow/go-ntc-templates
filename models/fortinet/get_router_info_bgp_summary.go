package fortinet

type GetRouterInfoBgpSummary struct {
	Bgp_neigh    string `json:"BGP_NEIGH"`
	Neigh_as     string `json:"NEIGH_AS"`
	Up_down      string `json:"UP_DOWN"`
	State_pfxrcd string `json:"STATE_PFXRCD"`
}

var GetRouterInfoBgpSummary_Template string = `Value BGP_NEIGH (\d+?\.\d+?\.\d+?\.\d+?)
Value NEIGH_AS (\d+)
Value UP_DOWN (\S+)
Value STATE_PFXRCD (\w+)

Start
  ^BGP\s+router\s+identifier\s+\d+?\.\d+?\.\d+?\.\d+?,\s+local\s+AS\s+number\s+\d+\s*$$
  ^BGP\s+table\s+version\s+is\s+\d+?$$
  ^\d+\s+BGP\s+AS-PATH\s+entries\s*$$
  ^\d+\s+BGP\s+community\s+entries\s*$$
  ^Neighbor\s+V\s+AS\s+MsgRcvd\s+MsgSent\s+TblVer\s+InQ\s+OutQ\s+Up\/Down\s+State\/PfxRcd.*$$
  ^${BGP_NEIGH}\s+\S+\s+${NEIGH_AS}(?:\s+\S+\s+){5}${UP_DOWN}\s+${STATE_PFXRCD}\s*$$ -> Record
  ^Total number\s+of\s+neighbors\s+\d*\s*$$
  ^\s*$$
  ^. -> Error`
