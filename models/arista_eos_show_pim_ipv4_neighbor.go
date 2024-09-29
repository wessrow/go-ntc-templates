package models

type AristaEosShowPimIpv4Neighbor struct {
	Vrf	string	`json:"VRF"`
	Neighbor	string	`json:"NEIGHBOR"`
	Interface	string	`json:"INTERFACE"`
	Uptime	string	`json:"UPTIME"`
	Expires	string	`json:"EXPIRES"`
	Mode	string	`json:"MODE"`
	Transport	string	`json:"TRANSPORT"`
}

var AristaEosShowPimIpv4Neighbor_Template = `Value Filldown VRF (\S+)
Value Required NEIGHBOR (\d+\.\d+\.\d+\.\d+)
Value INTERFACE ([\w\./-]+)
Value UPTIME (\S+)
Value EXPIRES (\S+)
Value MODE (\S+)
Value TRANSPORT (\S+)

Start
  ^\s*PIM Neighbor Table for ${VRF} VRF\s*$$
  ^\s*Neighbor Address\s+Interface\s+Uptime\s+Expires\s+Mode\s+Transport\s*$$
  ^\s*${NEIGHBOR}\s+${INTERFACE}\s+${UPTIME}\s+${EXPIRES}\s+${MODE}\s+${TRANSPORT}\s*$$ -> Record
  ^%\s
  ^. -> Error

EOF

`