package cisco_nxos

type ShowIpPimInterfaceBriefVrfAll struct {
	Vrf              string `json:"VRF"`
	Interface        string `json:"INTERFACE"`
	Ip_address       string `json:"IP_ADDRESS"`
	Pim_dr_address   string `json:"PIM_DR_ADDRESS"`
	Neighbor_count   string `json:"NEIGHBOR_COUNT"`
	Border_interface string `json:"BORDER_INTERFACE"`
}

var ShowIpPimInterfaceBriefVrfAll_Template string = `Value Filldown VRF (\S+)
Value Required INTERFACE ([\w\./-]+)
Value IP_ADDRESS (\d+\.\d+\.\d+\.\d+)
Value PIM_DR_ADDRESS (\d+\.\d+\.\d+\.\d+)
Value NEIGHBOR_COUNT (\d+)
Value BORDER_INTERFACE (\S+)

Start
  ^\s*PIM Interface Status for VRF "${VRF}"\s*$$
  ^\s*Interface\s+IP Address\s+PIM DR Address\s+Neighbor\s+Border\s*$$
  ^\s*Count\s+Interface\s*$$
  ^\s*${INTERFACE}\s+${IP_ADDRESS}\s+${PIM_DR_ADDRESS}\s+${NEIGHBOR_COUNT}\s+${BORDER_INTERFACE}\s*$$ -> Record
  ^. -> Error
`
