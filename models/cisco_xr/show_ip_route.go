package cisco_xr

type ShowIpRoute struct {
	Uptime        string `json:"UPTIME"`
	Network       string `json:"NETWORK"`
	Prefix_length string `json:"PREFIX_LENGTH"`
	Distance      string `json:"DISTANCE"`
	Next_hop      string `json:"NEXT_HOP"`
	Interface     string `json:"INTERFACE"`
	Vrf           string `json:"VRF"`
	Protocol      string `json:"PROTOCOL"`
	Metric        string `json:"METRIC"`
	Type          string `json:"TYPE"`
}

var ShowIpRoute_Template string = `Value Filldown VRF (\S+)
Value Filldown PROTOCOL (\S+|\S+\s\S+)
Value Filldown NETWORK (\d+\.\d+\.\d+\.\d+)
Value Filldown PREFIX_LENGTH (\d+)
Value DISTANCE (\d+|is)
Value METRIC (\d+)
Value TYPE (directly|via)
Value Required NEXT_HOP (connected|\d+\.\d+\.\d+\.\d+)
Value INTERFACE (\S+|vrf\s\S+)
Value UPTIME (\S+)


Start
  ^VRF?:\s${VRF}
  ^Codes:
  ^\s+(\S+\s+-\s+.+[,]*)+
  ^Gateway\s+of\s+last\s+resort
  ^${PROTOCOL}\s+${NETWORK}/${PREFIX_LENGTH}\s+(?:\[${DISTANCE}/${METRIC}\]|is)\s+${TYPE}\s+${NEXT_HOP},\s+${UPTIME},\s+${INTERFACE} -> Record
  ^\s+(?:\[${DISTANCE}/${METRIC}\]|is)\s+${TYPE}\s+${NEXT_HOP},\s+${UPTIME},\s+${INTERFACE} -> Record
  ^${PROTOCOL}\s+${NETWORK}/${PREFIX_LENGTH}\s+(?:\[${DISTANCE}/${METRIC}\]|is)\s+${TYPE}\s+${NEXT_HOP},\s+${UPTIME} -> Record
  ^${PROTOCOL}\s+${NETWORK}/${PREFIX_LENGTH}\s+(?:\[${DISTANCE}/${METRIC}\]|is)\s+${TYPE}\s+${NEXT_HOP}\s+\(nexthop\sin\s${INTERFACE}\),\s+${UPTIME} -> Record
  ^\s+(?:\[${DISTANCE}/${METRIC}\]|is)\s+${TYPE}\s+${NEXT_HOP},\s+${UPTIME} -> Record
  ^\s*$$
  ^\S+\s+\S+\s+\d+\s+\d+:\d+:\d+
  ^% No matching routes found
  ^% Network not in table
  ^\s*Routing\s+entry\s+for\s+${NETWORK}/${PREFIX_LENGTH} -> ROUTE_ENTRY
  ^\s+Route\s+metric\s+is\s+\d+
  ^\s+\S+,\s+from\s+\S+,\s+via\s+\S+
  ^\s+No\s+advertising\s+protos
  ^. -> Error "LINE NOT FOUND"

ROUTE_ENTRY
  ^\s+Known\s+via\s+\"${PROTOCOL}\",\s+distance\s+${DISTANCE},\s+metric\s+${METRIC},\s+candidate\s+default\s+path
  ^\s+Tag\s+\d+,\s+type\s+extern\s+\d+
  ^\s+Installed\s\S+\s+\S+\s+\S+\s+for\s+${UPTIME}
  ^\s+Routing\s+Descriptor\s+Blocks
  ^\s+${NEXT_HOP},\s+from\s+\S+,\s+${TYPE}\s+${INTERFACE} -> Record Start
  ^\s+${NEXT_HOP}$$ -> Record Start
  ^. -> Error ROUTE_ENTRY
`
