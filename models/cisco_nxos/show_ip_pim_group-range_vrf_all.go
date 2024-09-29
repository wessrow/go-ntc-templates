package cisco_nxos 

type ShowIpPimGroupRangeVrfAll struct {
	Vrf	string	`json:"VRF"`
	Group_range	string	`json:"GROUP_RANGE"`
	Action	string	`json:"ACTION"`
	Mode	string	`json:"MODE"`
	Rp_address	string	`json:"RP_ADDRESS"`
	Shared_tree_range	string	`json:"SHARED_TREE_RANGE"`
	Origin	string	`json:"ORIGIN"`
}

var ShowIpPimGroupRangeVrfAll_Template = `Value Filldown VRF (\S+)
Value Required GROUP_RANGE (\d+\.\d+\.\d+\.\d+/\d+)
Value ACTION (\S+)
Value MODE (\S+)
Value RP_ADDRESS (\S+)
Value SHARED_TREE_RANGE (\S+)
Value ORIGIN (\S+)


Start
  ^\s*PIM Group-Range Configuration for VRF "${VRF}"\s*$$
  ^\s*Group-range\s+Action\s+Mode\s+RP-address\s+Shrd-tree-range\s+Origin\s*$$
  ^\s*${GROUP_RANGE}\s+${ACTION}\s+${MODE}\s+${RP_ADDRESS}\s+${SHARED_TREE_RANGE}\s+${ORIGIN}\s*$$ -> Record
  ^. -> Error
`