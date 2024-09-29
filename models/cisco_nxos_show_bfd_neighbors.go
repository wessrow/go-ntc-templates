package models

type CiscoNxosShowBfdNeighbors struct {
	Our_address	string	`json:"OUR_ADDRESS"`
	Neighbor_address	string	`json:"NEIGHBOR_ADDRESS"`
	Ld_rd	string	`json:"LD_RD"`
	Rh_rs	string	`json:"RH_RS"`
	Holddown	string	`json:"HOLDDOWN"`
	State	string	`json:"STATE"`
	Interface	string	`json:"INTERFACE"`
	Vrf	string	`json:"VRF"`
	Type	string	`json:"TYPE"`
}

var CiscoNxosShowBfdNeighbors_Template = `Value OUR_ADDRESS (\d+?\.\d+?\.\d+?\.\d+?)
Value NEIGHBOR_ADDRESS (\d+?\.\d+?\.\d+?\.\d+?)
Value LD_RD (\S+)
Value RH_RS (\S+)
Value HOLDDOWN (\S+)
Value STATE (\S+)
Value INTERFACE (\S+)
Value VRF (\S+)
Value TYPE (\S+)

Start
  ^OurAddr\s+NeighAddr\s+LD\/RD\s+RH\/RS\s+Holdown\(mult\)\s+State\s+Int\s+Vrf\s+Type\s*$$ -> BFDNeigh
  ^. -> Error

BFDNeigh
  ^${OUR_ADDRESS}\s+${NEIGHBOR_ADDRESS}\s+${LD_RD}\s+${RH_RS}\s+${HOLDDOWN}\s+${STATE}\s+${INTERFACE}\s+${VRF}\s+${TYPE}\s*$$ -> Record
  ^. -> Error

`