package cisco_asa

type ShowBgpSummary struct {
	State_pfxrcd string `json:"STATE_PFXRCD"`
	Router_id    string `json:"ROUTER_ID"`
	Local_as     string `json:"LOCAL_AS"`
	Bgp_neigh    string `json:"BGP_NEIGH"`
	Neigh_as     string `json:"NEIGH_AS"`
	Up_down      string `json:"UP_DOWN"`
}

var ShowBgpSummary_Template string = `Value Filldown,Required ROUTER_ID ([0-9a-f:\.]+)
Value Filldown LOCAL_AS (\d+)
Value BGP_NEIGH (\d+?\.\d+?\.\d+?\.\d+?)
Value NEIGH_AS (\d+)
Value UP_DOWN (\S+?)
Value STATE_PFXRCD (\S+?\s+\S+?|\S+?)

Start
  ^BGP\s+router\s+identifier\s+${ROUTER_ID},\s+local\s+AS\s+number\s+${LOCAL_AS}\s*$$
  ^BGP\s+table\s+version\s+is\s+\d+,\s+main\s+routing\s+table\s+version\s+\d+\s*$$
  ^\d+\s+network\s+entries\s+using\s+\d+\s+bytes\s+of\s+memory\s*$$
  ^\d+\s+path\s+entries\s+using\s+\d+\s+bytes\s+of\s+memory\s*$$
  ^\d+\/\d+\s+BGP\s+path\/bestpath\s+attribute\s+entries\s+using\s+\d+\s+bytes\s+of\s+memory\s*$$
  ^\d+\s+BGP\s+AS-PATH\s+entries\s+using\s+\d+\s+bytes\s+of\s+memory\s*$$
  ^\d+\s+BGP\s+community\s+entries\s+using\s+\d+\s+bytes\s+of\s+memory\s*$$
  ^\d+\s+BGP\s+extended\s+community\s+entries\s+using\s+\d+\s+bytes\s+of\s+memory\s*$$
  ^\d+\s+BGP\s+route-map\s+cache\s+entries\s+using\s+\d+\s+bytes\s+of\s+memory\s*$$
  ^\d+\s+BGP\s+filter-list\s+cache\s+entries\s+using\s+\d+\s+bytes\s+of\s+memory\s*$$
  ^BGP\s+using\s\d+\s+total\s+bytes\s+of\s+memory\s*$$
  ^BGP\s+activity\s+\d+\/\d+\s+prefixes,\s+\d+\/\d+\s+paths,\s+scan\s+interval\s+\d+\s+secs\s*$$
  ^Neighbor\s+V\s+AS\s+MsgRcvd\s+MsgSent\s+TblVer\s+InQ\s+OutQ\s+Up\/Down\s+State\/PfxRcd
  ^${BGP_NEIGH}\s+\S+\s+${NEIGH_AS}(\s+\d+?){5}\s+${UP_DOWN}\s+${STATE_PFXRCD}\s*$$ -> Record
  ^\s*$$
  ^. -> Error

EOF`
