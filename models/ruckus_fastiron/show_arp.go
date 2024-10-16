package ruckus_fastiron

type ShowArp struct {
	Number      string `json:"NUMBER"`
	Ip_address  string `json:"IP_ADDRESS"`
	Mac_address string `json:"MAC_ADDRESS"`
	Type        string `json:"TYPE"`
	Age         string `json:"AGE"`
	Port        string `json:"PORT"`
	Status      string `json:"STATUS"`
	Vlan_id     string `json:"VLAN_ID"`
}

var ShowArp_Template string = `Value NUMBER (\S+)
Value Required IP_ADDRESS ([A-Fa-f0-9:\.]+)
Value Required MAC_ADDRESS (([A-Fa-f0-9\.]{14}|None))
Value TYPE (\S+)
Value AGE (\S+)
Value Required PORT (\S+)
Value STATUS (\S+)
Value VLAN_ID (\S+)

Start
  ^All\s+ARPs:\s+\d+,\s+maximum\s+capacity:\s+\d+
  ^Total\s+number\s+of\s+ARP\s+entries\s*:\s+\d+
  ^Entries\s+in\s+default\s+routing\s+instance:
  ^No\.\s+(IP\s+Address|IP)\s+(MAC\s+Address|MAC)\s+Type\s+Age\s+Port\s+Status\s+VLAN -> Layer2
  ^No\.\s+(IP\s+Address|IP)\s+(MAC\s+Address|MAC)\s+Type\s+Age\s+Port\s+Status -> Layer3
  ^\s*$$
  ^. -> Error 

Layer2
  ^${NUMBER}\s+${IP_ADDRESS}\s+${MAC_ADDRESS}\s+${TYPE}\s+${AGE}\s+${PORT}\s+${STATUS}\s+${VLAN_ID} -> Record
  ^Total\s+ARP\s+Entries\s*:\s+\d+
  ^\s*$$
  ^. -> Error 
  
Layer3
  ^${NUMBER}\s+${IP_ADDRESS}\s+${MAC_ADDRESS}\s+${TYPE}\s+${AGE}\s+${PORT}\s+${STATUS} -> Record
  ^\s*$$
  ^. -> Error 
`
