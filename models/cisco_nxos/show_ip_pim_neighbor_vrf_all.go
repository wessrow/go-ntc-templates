package cisco_nxos 

type ShowIpPimNeighborVrfAll struct {
	Vrf	string	`json:"VRF"`
	Neighbor	string	`json:"NEIGHBOR"`
	Interface	string	`json:"INTERFACE"`
	Uptime	string	`json:"UPTIME"`
	Expires	string	`json:"EXPIRES"`
	Dr_priority	string	`json:"DR_PRIORITY"`
	Bidir_capable	string	`json:"BIDIR_CAPABLE"`
	Bfd_state	string	`json:"BFD_STATE"`
	Ecmp_redirect_capable	string	`json:"ECMP_REDIRECT_CAPABLE"`
}

var ShowIpPimNeighborVrfAll_Template = `Value Filldown VRF (\S+)
Value Required NEIGHBOR (\d+\.\d+\.\d+\.\d+)
Value INTERFACE ([\w\./-]+)
Value UPTIME (\S+)
Value EXPIRES (\S+)
Value DR_PRIORITY (\S+)
Value BIDIR_CAPABLE (\S+)
Value BFD_STATE (\S+)
Value ECMP_REDIRECT_CAPABLE (\S+)

Start
  ^\s*\^
  ^%\s+Invalid -> Clear
  ^\s*PIM Neighbor Status for VRF "${VRF}"\s*$$
  ^\s*Neighbor\s+Interface\s+Uptime\s+Expires\s+DR\s+Bidir-\s+BFD\s+ECMP Redirect\s*$$
  ^\s*Priority\s+Capable\s+State\s+Capable\s*$$
  ^\s*${NEIGHBOR}\s+${INTERFACE}\s+${UPTIME}\s+${EXPIRES}\s+${DR_PRIORITY}\s+${BIDIR_CAPABLE}\s+${BFD_STATE}\s+${ECMP_REDIRECT_CAPABLE}\s*$$ -> Record
  ^. -> Error

EOF
`