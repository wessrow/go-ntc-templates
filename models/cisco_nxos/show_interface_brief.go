package cisco_nxos

type ShowInterfaceBrief struct {
	Vrf             string `json:"VRF"`
	Ip_address      string `json:"IP_ADDRESS"`
	Type            string `json:"TYPE"`
	Mtu             string `json:"MTU"`
	Vlan_id         string `json:"VLAN_ID"`
	Portch          string `json:"PORTCH"`
	Description     string `json:"DESCRIPTION"`
	Interface       string `json:"INTERFACE"`
	Speed           string `json:"SPEED"`
	Mode            string `json:"MODE"`
	Protocol        string `json:"PROTOCOL"`
	Status          string `json:"STATUS"`
	Peer_ip_address string `json:"PEER_IP_ADDRESS"`
	Reason          string `json:"REASON"`
	Vcid            string `json:"VCID"`
}

var ShowInterfaceBrief_Template string = `Value INTERFACE ([\w+/\.]+)
Value VRF (\S+)
Value STATUS (up|down)
Value IP_ADDRESS (\S+)
Value PEER_IP_ADDRESS (\d+\.\d+\.\d+\.\d+)
Value SPEED (\S+)
Value MTU (\d+)
Value VLAN_ID (\S+)
Value TYPE (\S+)
Value MODE (\S+)
Value REASON (\S+((\s\w+)+)?)
Value PORTCH (\S+)
Value DESCRIPTION (\S+((\s\w+)+)?)
Value PROTOCOL (\S+)
Value VCID (\S+)

Start
  ^Port\s+VRF\s+Status\s+IP\s+Address\s+Speed\s+MTU -> Management
  ^Ethernet\s+VLAN\s+Type\sMode\s+Status\s+Reason\s+Speed\s+Port\s*$$ -> Ethernet
  ^Interface\s+Ch\s+ -> Ethernet
  ^Interface\s+Status\s+Description\s*$$ -> Loopback
  ^Interface\s+Secondary\s+VLAN\(Type\)\s+Status\s+Reason -> VLAN
  ^Port-channel\s+VLAN\s+Type\sMode\s+Status\s+Reason\s+Speed\s+Protocol -> PORTCHANNEL
  ^Port\s+Status\s+Reason\s+MTU\s+VRF -> NVE
  ^Port\s+Status\s+Reason\s+MTU\s*$$ -> NVE
  ^Interface\s+Status\s+IP\s+Address\s+Encap\s+type\s+MTU -> TUNNEL
  ^Interface\s+Status\s+Peer\s+IP\s+Address\s+VC\s+ID -> PSEUDOWIRE
  ^---+$$
  ^. -> Error Start

Management
  ^${INTERFACE}\s+${VRF}\s+${STATUS}\s+${IP_ADDRESS}\s+${SPEED}\s+${MTU} -> Record
  ^Ethernet\s+VLAN\s+Type\s+Mode\s+Status\s+Reason\s+Speed\s+Port -> Start
  ^---+$$
  ^\s*$$ -> Start
  ^. -> Error Management

Ethernet
  ^Interface\s+Ch\s+#\s*$$
  ^${INTERFACE}\s+${VLAN_ID}\s+${TYPE}\s+${MODE}\s+${STATUS}\s+${REASON}\s+${SPEED}\s+${PORTCH} -> Record
  ^---+$$
  ^\s*$$ -> Start
  ^. -> Error Etherenet

Loopback
  ^${INTERFACE}\s+${STATUS}\s+${DESCRIPTION} -> Record
  ^---+$$
  ^\s*$$ -> Start
  ^. -> Error Loopback

VLAN
  ^${INTERFACE}\s+${TYPE}\s+${STATUS}\s+${REASON} -> Record
  ^---+$$
  ^\s*$$ -> Start
  ^. -> Error VLAN

PORTCHANNEL
  ^Interface\s*$$
  ^${INTERFACE}\s+${VLAN_ID}\s+${TYPE}\s+${MODE}\s+${STATUS}\s+${REASON}\s+${SPEED}\s+${PROTOCOL} -> Record
  ^---+$$
  ^\s*$$ -> Start
  ^. -> Error PORTCHANNEL

TUNNEL
  ^${INTERFACE}\s+${STATUS}\s+${IP_ADDRESS}\s+${TYPE}\s+${MTU} -> Record
  ^---+$$
  ^\s*$$ -> Start
  ^. -> Error TUNNEL

NVE
  ^${INTERFACE}\s+${STATUS}\s+${REASON}\s+${MTU}\s+${VRF} -> Record
  ^${INTERFACE}\s+${STATUS}\s+${REASON}\s+${MTU} -> Record
  ^---+$$
  ^\s*$$ -> Start
  ^. -> Error NVE

PSEUDOWIRE
  ^${INTERFACE}\s+${STATUS}\s+${PEER_IP_ADDRESS}\s+${VCID}
`
