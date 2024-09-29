package models

type CiscoXrShowLptsPifibHardwarePoliceLocation struct {
	Location	string	`json:"LOCATION"`
	Flowtype	string	`json:"FLOWTYPE"`
	Policer	string	`json:"POLICER"`
	Type	string	`json:"TYPE"`
	Cur_rate	string	`json:"CUR_RATE"`
	Def_rate	string	`json:"DEF_RATE"`
	Accepted	string	`json:"ACCEPTED"`
	Dropped	string	`json:"DROPPED"`
	Tos_value	string	`json:"TOS_VALUE"`
}

var CiscoXrShowLptsPifibHardwarePoliceLocation_Template = `Value Filldown LOCATION (\S+?)
Value FLOWTYPE (\S+)
Value POLICER (\d+)
Value TYPE (\S+)
Value CUR_RATE (\d+)
Value DEF_RATE (\d+)
Value ACCEPTED (\d+)
Value DROPPED (\d+)
Value TOS_VALUE (\d+)

Start
  ^\s+Node\s+${LOCATION}:
  ^FlowType\s+Policer\s+Type\s+Cur.\s+Rate\s+Def.\s+Rate\s+Accepted\s+Dropped\s+TOS\s+Value -> Parse

Parse
  ^${FLOWTYPE}\s+${POLICER}\s+${TYPE}\s+${CUR_RATE}\s+${DEF_RATE}\s+${ACCEPTED}\s+${DROPPED}\s+${TOS_VALUE} -> Record
  ^statistics:
  ^Packets\s+accepted\s+by\s+deleted\s+entries:
  ^Packets\s+dropped\s+by\s+deleted\s+entries:
  ^Run\s+out\s+of\s+statistics\s+counter\s+errors:
  ^Statistics\s+last\s+cleared:
  ^[ -]*$$
  ^.* -> Error "LINE NOT FOUND"

EOF

`