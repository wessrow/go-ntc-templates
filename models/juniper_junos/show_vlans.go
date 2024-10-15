package juniper_junos

type ShowVlans struct {
	Interfaces       []string `json:"INTERFACES"`
	Routing_instance string   `json:"ROUTING_INSTANCE"`
	Vlan_name        string   `json:"VLAN_NAME"`
	Tag              string   `json:"TAG"`
}

var ShowVlans_Template string = `Value ROUTING_INSTANCE (\S+)
Value VLAN_NAME (\S+)
Value TAG (\d+)
Value List INTERFACES (\S+)
		   
Start
  ^Routing\sinstance\s+VLAN\sname\s+Tag\s+Interfaces$$ -> VLAN
		   
VLAN
  ^\S -> Continue.Record
  ^${ROUTING_INSTANCE}\s+${VLAN_NAME}\s+${TAG}\s*$$
  ^\s+${INTERFACES}$$
  ^\s*$$
  ^{master:\d+}
  ^. -> Error`
