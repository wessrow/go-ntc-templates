package cisco_nxos

type ShowBfdNeighbors struct {
	Vrf              string `json:"VRF"`
	Neighbor_address string `json:"NEIGHBOR_ADDRESS"`
	Ld_rd            string `json:"LD_RD"`
	Holddown         string `json:"HOLDDOWN"`
	State            string `json:"STATE"`
	Interface        string `json:"INTERFACE"`
	Our_address      string `json:"OUR_ADDRESS"`
	Rh_rs            string `json:"RH_RS"`
	Type             string `json:"TYPE"`
}

var ShowBfdNeighbors_Template string = `Value OUR_ADDRESS (\d+?\.\d+?\.\d+?\.\d+?)
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
