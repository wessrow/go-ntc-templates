package models

type HpProcurveShowIpRoute struct {
	Destination	string	`json:"DESTINATION"`
	Gateway	string	`json:"GATEWAY"`
	Vlan_name	string	`json:"VLAN_NAME"`
	Type	string	`json:"TYPE"`
	Subtype	string	`json:"SUBTYPE"`
	Metric	string	`json:"METRIC"`
	Distance	string	`json:"DISTANCE"`
}

var HpProcurveShowIpRoute_Template = `Value DESTINATION ([0-9./]+)
Value GATEWAY ([0-9.]*|\S*)
Value VLAN_NAME (\d*)
Value TYPE (\S+)
Value SUBTYPE (\S*)
Value METRIC (\d+)
Value DISTANCE (\d+)

Start
  ^.*IP Route Entries
  ^.*Destination -> Route
  ^\s*$$
  ^. -> Error

Route
  ^\s*----
  ^\s*${DESTINATION}\s+${GATEWAY}\s+${VLAN_NAME}\s+${TYPE}\s+${SUBTYPE}\s+${METRIC}\s+${DISTANCE} -> Record
  ^\s*$$
  ^. -> Error

`