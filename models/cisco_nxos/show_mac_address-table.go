package cisco_nxos

type ShowMacAddressTable struct {
	Vlan_id     string `json:"VLAN_ID"`
	Mac_address string `json:"MAC_ADDRESS"`
	Type        string `json:"TYPE"`
	Age         string `json:"AGE"`
	Secure      string `json:"SECURE"`
	Ntfy        string `json:"NTFY"`
	Ports       string `json:"PORTS"`
}

var ShowMacAddressTable_Template string = `Value VLAN_ID (\S+)
Value MAC_ADDRESS (\S+)
Value TYPE (\S+)
Value AGE (\S+)
Value SECURE ([TF])
Value NTFY ([TF])
Value PORTS (\S+)

Start
  ^VLAN\s+MAC\s+Address\s+Type\s+age\s+Secure\s+NTFY\s+Ports -> Continue
  ^.*\s${VLAN_ID}\s+${MAC_ADDRESS}\s+${TYPE}\s+${AGE}\s+${SECURE}\s+${NTFY}\s+${PORTS} -> Record
`
