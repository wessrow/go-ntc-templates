package zyxel_os

type CfgIntfGroupGet struct {
	Index     string `json:"INDEX"`
	Name      string `json:"NAME"`
	Wan_intf  string `json:"WAN_INTF"`
	Interface string `json:"INTERFACE"`
	Criteria  string `json:"CRITERIA"`
}

var CfgIntfGroupGet_Template string = `Value INDEX (\d+)
Value NAME (.+\w)
Value WAN_INTF (Any\sWAN|((WWAN|ADSL|VDSL),?)+)
Value INTERFACE ((N/A|.+\w))
Value CRITERIA ((N/A|.+>))

Start
  ^Index\s+Name\s+WAN\sInterface\s+LAN\sInterfaces\s+Criteria\s*$$ -> INTFGroupTable
  ^\s*$$
  ^. -> Error

INTFGroupTable
  ^${INDEX}\s+${NAME}\s+${WAN_INTF}\s+${INTERFACE}\s+${CRITERIA}\s*$$ -> Record
  ^Command Successful.\s*$$
  ^\s*$$
  ^. -> Error
`
