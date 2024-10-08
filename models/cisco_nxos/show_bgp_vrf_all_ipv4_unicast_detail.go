package cisco_nxos 

type ShowBgpVrfAllIpv4UnicastDetail struct {
	Active	string	`json:"ACTIVE"`
	As_path	string	`json:"AS_PATH"`
	Community	string	`json:"COMMUNITY"`
	Filtered	string	`json:"FILTERED"`
	Local_pref	string	`json:"LOCAL_PREF"`
	Metric	string	`json:"METRIC"`
	Neighbor	string	`json:"NEIGHBOR"`
	Neighbor_id	string	`json:"NEIGHBOR_ID"`
	Next_hop	string	`json:"NEXT_HOP"`
	Origin	string	`json:"ORIGIN"`
	Path_type	string	`json:"PATH_TYPE"`
	Prefix	string	`json:"PREFIX"`
	Prefix_length	string	`json:"PREFIX_LENGTH"`
	Valid	string	`json:"VALID"`
	Vrf	string	`json:"VRF"`
	Weight	string	`json:"WEIGHT"`
}

var ShowBgpVrfAllIpv4UnicastDetail_Template = `Value ACTIVE (best)
Value Required AS_PATH ((?:\d+(\s.*?)*(\s.?)*|(?:\d+\.\d+(\s.*?)*))|(?:Local.*)|(?:NONE))
Value COMMUNITY ((?:(?:(?:no-export|no-advertise|local-as|internet)|\d+\s?:\s?\d+)\s?)*)
Value FILTERED (received\s+only)
Value LOCAL_PREF (\d+)
Value METRIC (\d+)
Value NEIGHBOR (\S+)
Value NEIGHBOR_ID (\S+)
Value NEXT_HOP (\d+\.\d+\.\d+\.\d+)
Value ORIGIN (\S+)
Value PATH_TYPE (\S+)
Value Filldown PREFIX ([A-Fa-f0-9:\.]+)
Value Filldown PREFIX_LENGTH (\d+)
Value VALID (\w+)
Value Filldown VRF (\S+)
Value WEIGHT (\S+)

Start
  ^BGP\s+routing\s+table\sinformation\s+for\s+VRF\s+${VRF}
  ^BGP\s+routing\s+table\s+entry\s+for\s+${PREFIX}/${PREFIX_LENGTH}
  ^Paths:
  ^Flags:
  ^\s+Advertised\s+path-id\s
  ^\s+Path\s+type: -> Continue.Record
  ^\s+Path\s+type:\s+${PATH_TYPE},\s+path\s+is\s+${VALID},\s+is\s+${ACTIVE}\s+path
  ^\s+Path\s+type:\s+${PATH_TYPE},\s+path\s+is\s+${VALID},\s+${FILTERED},
  ^\s+Path\s+type:\s+${PATH_TYPE},\s+path\s+is\s+${VALID}
  ^\s+AS-Path:\s+${AS_PATH}\s+,
  ^\s+AS-Path:\s+${AS_PATH},
  ^\s+${NEXT_HOP}\s+\(metric\s+\d+\)\s+from\s+${NEIGHBOR}\s+\(${NEIGHBOR_ID}\)
  ^\s+Origin\s+${ORIGIN},\s+MED\s+${METRIC},\s+localpref\s+${LOCAL_PREF},\s+weight\s+${WEIGHT}
  ^\s+Origin\s+${ORIGIN},\s+MED\s+not\s+set,\s+localpref\s+${LOCAL_PREF},\s+weight\s+${WEIGHT}
  ^\s+Aggregated\s+by
  ^\s+Community:\s+${COMMUNITY}\s+$$
  ^\s+Community:\s+${COMMUNITY}$$
  ^\s+Extcommunity:\s+ -> Record
  ^\s+Path-id\s+\d+\s+
  ^\s+\b(?:\d{1,3}\.){3}\d{1,3}\b
  ^. -> Error 
`