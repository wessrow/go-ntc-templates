package cisco_xr

type ShowInterfaceBrief struct {
	Int_bw      string `json:"INT_BW"`
	Interface   string `json:"INTERFACE"`
	Intf_state  string `json:"INTF_STATE"`
	Linep_state string `json:"LINEP_STATE"`
	Encap_type  string `json:"ENCAP_TYPE"`
	Mtu         string `json:"MTU"`
}

var ShowInterfaceBrief_Template string = `Value INTERFACE ([\w+/]+)
Value INTF_STATE (up|down|admin-down)
Value LINEP_STATE (up|down|admin-down)
Value ENCAP_TYPE (\S+)
Value MTU (\d+)
Value INT_BW (\d+)


Start
  ^\s+${INTERFACE}\s+${INTF_STATE}\s+${LINEP_STATE}\s+${ENCAP_TYPE}\s+${MTU}\s+${INT_BW} -> Record
`
