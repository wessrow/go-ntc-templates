package cisco_asa 

type ShowRoute struct {
	Protocol	string	`json:"PROTOCOL"`
	Type	string	`json:"TYPE"`
	Network	string	`json:"NETWORK"`
	Netmask	string	`json:"NETMASK"`
	Distance	string	`json:"DISTANCE"`
	Metric	string	`json:"METRIC"`
	Nexthopip	string	`json:"NEXTHOPIP"`
	Nexthopif	string	`json:"NEXTHOPIF"`
	Uptime	string	`json:"UPTIME"`
}

var ShowRoute_Template = `Value Filldown PROTOCOL (C|S|R|B|D|O|i|L)
Value Filldown TYPE (\w{0,2})
Value Required,Filldown NETWORK (\d+\.\d+\.\d+\.\d+|\S+)
Value Filldown NETMASK (\d+\.\d+\.\d+\.\d+)
Value DISTANCE (\d+)
Value METRIC (\d+)
Value NEXTHOPIP (\d+\.\d+\.\d+\.\d+)
Value NEXTHOPIF (\S+)
Value UPTIME (\d\S+?)

Start
  # Skips over the code line that explains what each code means
  ^Codes:
  # Skips over the definitions for the codes
  ^\s{6,}\S*\s-\s\S.*
  ^Gateway -> ROUTES
  ^. -> Error

ROUTES
  # Match regular routes with all data in same line
  ^${PROTOCOL}(\s|\*)${TYPE}\s+${NETWORK}\s+${NETMASK}\s\[${DISTANCE}\/${METRIC}\]\svia\s${NEXTHOPIP}(,\s${UPTIME}){0,1}(,\s${NEXTHOPIF}){0,1}\s*$$ -> Record
  #
  # Clear all non Filldown variables when line started with network that is variably subnetted
  ^\s+[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}${NETMASK}\s*$$ -> Clear
  ^${PROTOCOL}(\s|\*)${TYPE}\s+${NETWORK}\s\[${DISTANCE}\/${METRIC}\]\svia\s${NEXTHOPIP}(,\s${UPTIME}){0,1}(,\s${NEXTHOPIF}){0,1}\s*$$ -> Record
  #
  # Match multiline route statements
  ^${PROTOCOL}(\s|\*)${TYPE}\s+${NETWORK}\s+${NETMASK}\s*$$
  #
  # Match load-balanced routes
  ^\s+\[${DISTANCE}\/${METRIC}\]\s+via\s+${NEXTHOPIP},(?:\s+${UPTIME},)?\s+${NEXTHOPIF}\s*$$ -> Record
  #
  # Match directly connected routes
  ^${PROTOCOL}\s${TYPE}\s+${NETWORK}\sis\sdirectly\sconnected,\s${NEXTHOPIF} -> Record
  ^${PROTOCOL}(\*){0,1}\s${TYPE}\s+${NETWORK}\s+${NETMASK}\sis\sdirectly\sconnected,\s${NEXTHOPIF} -> Record
  ^\s+is\sdirectly\sconnected,\s${NEXTHOPIF} -> Record
  #
  # Clear all variables on empty lines
  ^\s*$$ -> Clearall
  #^${TYPE} -> Continue.Record
  #^${TYPE}\s+${NETWORK}\s+${NETMASK}\s+\[\d+\/\d+\]\s+via\s+${GATEWAY}\,\s+${UPTIME},\s+${INTFC}\s*$$ -> Record
  #^\s+\[\d+\/\d+\]\s+via\s+${GATEWAY}\,\s+${UPTIME},\s+${INTFC}\s*$$ -> Record
  ^. -> Error

EOF
`