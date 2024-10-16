package mikrotik_routeros

type InterfaceEthernetPrint struct {
	Status      string `json:"STATUS"`
	Slave       string `json:"SLAVE"`
	Name        string `json:"NAME"`
	Mtu         string `json:"MTU"`
	Mac_address string `json:"MAC_ADDRESS"`
	Arp         string `json:"ARP"`
	Switch      string `json:"SWITCH"`
	Id          string `json:"ID"`
}

var InterfaceEthernetPrint_Template string = `Value ID (\d+)
Value STATUS (X|R)
Value SLAVE (S)
Value NAME (\S+)
Value MTU (\d+)
Value MAC_ADDRESS ([a-zA-Z0-9]{2}(:[a-zA-Z0-9]{2}){5})
Value ARP (\S+)
Value SWITCH (\S+)

Start
  ^\s*Flags:\s+X\s+-\s+disabled,\s+R\s+-\s+running,\s+S\s+-\s+slave\s*$$
  ^\s*#\s+NAME\s+MTU\s+MAC-ADDRESS\s+ARP\s+SWITCH\s*$$ -> IfacesTable

IfacesTable
  ^\s*${ID}\s+(?:${STATUS})?(?:${SLAVE})?\s+${NAME}\s+${MTU}\s+${MAC_ADDRESS}\s+${ARP}\s+${SWITCH}\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`
