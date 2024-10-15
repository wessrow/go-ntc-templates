package huawei_vrp

type DisplayArpBrief struct {
	Expire      string `json:"EXPIRE"`
	Type        string `json:"TYPE"`
	Interface   string `json:"INTERFACE"`
	Vlan_id     string `json:"VLAN_ID"`
	Pvc         string `json:"PVC"`
	Ip_address  string `json:"IP_ADDRESS"`
	Mac_address string `json:"MAC_ADDRESS"`
}

var DisplayArpBrief_Template string = `Value IP_ADDRESS (\d+.\d+.\d+.\d+)
Value MAC_ADDRESS ((\w+-?)+)
Value EXPIRE (\d+)
Value TYPE (I\s-|[DS][F-][\d\s])
Value INTERFACE (\S+)
Value VLAN_ID (\d+/(\d+|-))
Value PVC (\S+)

Start
  ^\s*IP\sADDRESS.+$$ -> Next
  ^\s+VLAN.+$$ -> Next
  ^-+ -> Next
  ^\s*\d+.\d+.\d+.\d+ -> Continue.Record
  ^\s*${IP_ADDRESS}\s+(${MAC_ADDRESS}\s+)?(${EXPIRE}\s+)?${TYPE}\s+${INTERFACE}(\s+${VLAN_ID})?(\s+${PVC})?\s*$$
  ^\s*${VLAN_ID}(\s+${PVC})?\s*$$
  ^Total:.+$$ -> Next
  ^.*$$ -> Error
`
