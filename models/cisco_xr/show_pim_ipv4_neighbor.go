package cisco_xr

type ShowPimIpv4Neighbor struct {
	Vrf         string `json:"VRF"`
	Neighbor    string `json:"NEIGHBOR"`
	Interface   string `json:"INTERFACE"`
	Uptime      string `json:"UPTIME"`
	Expires     string `json:"EXPIRES"`
	Dr_priority string `json:"DR_PRIORITY"`
	Flags       string `json:"FLAGS"`
}

var ShowPimIpv4Neighbor_Template string = `Value Filldown VRF (\S+)
Value Required NEIGHBOR (\d+\.\d+\.\d+\.\d+)
Value INTERFACE ([\w\./-]+)
Value UPTIME ((\S+(\s\S+)*))
Value EXPIRES (\S+)
Value DR_PRIORITY (\d+(\s\(DR\))?)
Value FLAGS ((\S+(\s\S+)*))

Start
  ^\s*PIM neighbors in VRF ${VRF}\s*$$
  ^\s*Neighbor Address\s+Interface\s+Uptime\s+Expires\s+DR\s+pri\s+Flags\s*$$
  ^\s*${NEIGHBOR}(\*)?\s+${INTERFACE}\s+${UPTIME}\s+${EXPIRES}\s+${DR_PRIORITY}(\s+${FLAGS})?\s*$$ -> Record




`
