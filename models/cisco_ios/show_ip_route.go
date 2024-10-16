package cisco_ios

type ShowIpRoute struct {
	Nexthop_vrf   string `json:"NEXTHOP_VRF"`
	Uptime        string `json:"UPTIME"`
	Flag          string `json:"FLAG"`
	Vrf           string `json:"VRF"`
	Protocol      string `json:"PROTOCOL"`
	Nexthop_ip    string `json:"NEXTHOP_IP"`
	Distance      string `json:"DISTANCE"`
	Metric        string `json:"METRIC"`
	Nexthop_if    string `json:"NEXTHOP_IF"`
	Type          string `json:"TYPE"`
	Network       string `json:"NETWORK"`
	Prefix_length string `json:"PREFIX_LENGTH"`
}

var ShowIpRoute_Template string = `Value Filldown VRF (\S+)
Value Filldown PROTOCOL (\w)
Value Filldown TYPE (\w{0,2})
Value Required,Filldown NETWORK (\d{1,3}.\d{1,3}.\d{1,3}.\d{1,3})
Value Filldown PREFIX_LENGTH (\d{1,2})
Value DISTANCE (\d+)
Value METRIC (\d+)
Value NEXTHOP_IP (\d{1,3}.\d{1,3}.\d{1,3}.\d{1,3})
Value NEXTHOP_VRF (\S+)
Value NEXTHOP_IF ([A-Za-z][\w\-\.:/]+)
Value UPTIME (\d[\w:\.]+)
Value FLAG ([\*%p])

Start
  ^Routing\s+Table:\s${VRF}\s*$$
  ^Gateway.* -> Routes
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is

Routes
  ^Routing\s+Table: -> Continue.Clearall
  ^Routing\s+Table:\s+${VRF}\s*$$
  # For "is (variably )subnetted" line, capture mask, clear all values.
  ^\s+\d{1,3}.\d{1,3}.\d{1,3}.\d{1,3}\/${PREFIX_LENGTH}\sis -> Clear
  #
  # Match directly connected route with explicit mask
  ^${PROTOCOL}(\s|${FLAG})${TYPE}\s+${NETWORK}\/${PREFIX_LENGTH}\sis\sdirectly\sconnected(,\s${UPTIME})?(,\s${NEXTHOP_IF})? -> Record
  #
  # Match directly connected route (mask is inherited from "is subnetted")
  ^${PROTOCOL}(\s|${FLAG})${TYPE}\s+${NETWORK}\sis\sdirectly\sconnected,\s${NEXTHOP_IF} -> Record
  #
  # Match regular routes, with mask, where all data in same line
  ^${PROTOCOL}(\s|${FLAG})${TYPE}\s+${NETWORK}\/${PREFIX_LENGTH}\s\[${DISTANCE}/${METRIC}\]\svia\s${NEXTHOP_IP}(\s\(${NEXTHOP_VRF}\))?(,\s${UPTIME})?(,\s${NEXTHOP_IF})? -> Record
  #
  # Match regular route, all one line, where mask is learned from "is subnetted" line
  ^${PROTOCOL}(\s|${FLAG})${TYPE}\s+${NETWORK}\s\[${DISTANCE}\/${METRIC}\]\svia\s${NEXTHOP_IP}(,\s${UPTIME})?(,\s${NEXTHOP_IF})? -> Record
  #
  # Match route with no via statement (Null via protocol)
  ^${PROTOCOL}(\s|${FLAG})${TYPE}\s+${NETWORK}\/${PREFIX_LENGTH}\s\[${DISTANCE}/${METRIC}\],\s${UPTIME},\s${NEXTHOP_IF} -> Record
  #
  # Match "is a summary" routes (often Null0)
  ^${PROTOCOL}(\s|${FLAG})${TYPE}\s+${NETWORK}\/${PREFIX_LENGTH}\sis\sa\ssummary,\s${UPTIME},\s${NEXTHOP_IF} -> Record
  #
  # Match regular routes where the network/mask is on the line above the rest of the route
  ^${PROTOCOL}(\s|${FLAG})${TYPE}\s+${NETWORK}\/${PREFIX_LENGTH}
  #
  # Match regular routes where the network only (mask from subnetted line) is on the line above the rest of the route
  ^${PROTOCOL}(\s|${FLAG})${TYPE}\s+${NETWORK}
  #
  # Match the rest of the route information on line below network (and possibly mask)
  ^\s+\[${DISTANCE}\/${METRIC}\]\svia\s${NEXTHOP_IP}(,\s${UPTIME})?(,\s${NEXTHOP_IF})? -> Record
  #
  # Match load-balanced routes
  ^\s+\[${DISTANCE}\/${METRIC}\]\svia\s${NEXTHOP_IP} -> Record

EOF
`
