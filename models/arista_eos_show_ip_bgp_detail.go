package models

type AristaEosShowIpBgpDetail struct {
	Vrf	string	`json:"VRF"`
	Local_pref	string	`json:"LOCAL_PREF"`
	Weight	string	`json:"WEIGHT"`
	Origin	string	`json:"ORIGIN"`
	Prefix	string	`json:"PREFIX"`
	Prefix_length	string	`json:"PREFIX_LENGTH"`
	As_path	string	`json:"AS_PATH"`
	Valid	string	`json:"VALID"`
	Path_type	string	`json:"PATH_TYPE"`
	Active	string	`json:"ACTIVE"`
	Community	string	`json:"COMMUNITY"`
	Next_hop	string	`json:"NEXT_HOP"`
	Neighbor	string	`json:"NEIGHBOR"`
	Neighbor_id	string	`json:"NEIGHBOR_ID"`
	Router_id	string	`json:"ROUTER_ID"`
	Local_as	string	`json:"LOCAL_AS"`
	Metric	string	`json:"METRIC"`
	Backup	string	`json:"BACKUP"`
}

var AristaEosShowIpBgpDetail_Template = `Value Filldown VRF (\S+)
Value LOCAL_PREF (\d+)
Value WEIGHT (\S+)
Value ORIGIN (.+)
Value Filldown PREFIX ([A-Fa-f0-9:\.]+)
Value Filldown PREFIX_LENGTH (\d+)
Value AS_PATH ((?:\d+(\s.*?)*(\s.?)*|(?:\d+\.\d+(\s.*?)*))|(?:Local.*))
Value VALID (\S+)
Value PATH_TYPE (\S+)
Value ACTIVE (best)
Value COMMUNITY ((?:(?:(?:no-export|no-advertise|local-as|internet)|\d+\s?:\s?\d+)\s?)*)
Value NEXT_HOP (\d+\.\d+\.\d+\.\d+)
Value NEIGHBOR (\S+)
Value NEIGHBOR_ID (\S+)
Value Filldown ROUTER_ID (\S+)
Value Filldown LOCAL_AS (\S+)
Value METRIC (\d+)
Value BACKUP (backup)

Start
  ^BGP\s+routing\s+table\s+information\s+for\s+VRF\s+${VRF}\s*$$
  ^Router\s+identifier\s+${ROUTER_ID},\s+local\s+AS\s+number\s+${LOCAL_AS}
  ^BGP\s+routing\s+table\s+entry\s+for\s+${PREFIX}(?:/${PREFIX_LENGTH})?
  ^\s+Paths:\s+\d+\s+available
  ^\s+${AS_PATH}\s*$$
  ^\s+\-\s+from\s+\S+\s+\(${NEIGHBOR_ID}\)
  ^\s+\-\s+from\s+${NEIGHBOR}\s+\(${NEIGHBOR_ID}\)
  ^\s+${NEXT_HOP}\s+from\s+${NEIGHBOR}\s+\(${NEIGHBOR_ID}\)
  ^\s+Origin\s+${ORIGIN},\s+metric\s+${METRIC},\s+localpref\s+${LOCAL_PREF},\s+IGP\s+metric\s+\S+,\s+weight\s+\-,\s+received\s+\S+\s+ago,\s+${VALID},\s+${PATH_TYPE},\s+${ACTIVE}
  ^\s+Origin\s+${ORIGIN},\s+metric\s+${METRIC},\s+localpref\s+${LOCAL_PREF},\s+IGP\s+metric\s+\S+,\s+weight\s+${WEIGHT},\s+received\s+\S+\s+ago,\s+${VALID},\s+${PATH_TYPE},\s+${ACTIVE}
  ^\s+Origin\s+${ORIGIN},\s+metric\s+${METRIC},\s+localpref\s+${LOCAL_PREF},\s+IGP\s+metric\s+\S+,\s+weight\s+${WEIGHT},\s+received\s+\S+\s+ago,\s+${VALID},\s+${PATH_TYPE},\s+${BACKUP}
  ^\s+Origin\s+${ORIGIN},\s+metric\s+${METRIC},\s+localpref\s+${LOCAL_PREF},\s+IGP\s+metric\s+\S+,\s+weight\s+${WEIGHT},\s+received\s+\S+\s+ago,\s+${VALID},\s+${PATH_TYPE},\s+atomic-aggregate
  ^\s+Origin\s+${ORIGIN},\s+metric\s+${METRIC},\s+localpref\s+${LOCAL_PREF},\s+IGP\s+metric\s+\S+,\s+weight\s+${WEIGHT},\s+received\s+\S+\s+ago,\s+${VALID},\s+${PATH_TYPE}
  ^\s+Community:\s+${COMMUNITY}$$
  ^\s+Extended\s+Community:\s+\S+
  ^\s+Rx\s+path\s+id:\s+\S+
  ^\s+Rx\s+SAFI:\s+\S+ -> Record
  ^\s+\d+\s+Contributing\s+routes:
  ^\s+\S+\s+Proto:\s+\S+\s+Origin:\s+\S+\s+AS\s+Path:\s+\S+
  ^. -> Error

EOF

`