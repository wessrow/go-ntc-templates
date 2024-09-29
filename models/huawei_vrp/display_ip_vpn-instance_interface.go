package huawei_vrp 

type DisplayIpVpnInstanceInterface struct {
	Name	string	`json:"NAME"`
	Id	string	`json:"ID"`
	Interface_list	[]string	`json:"INTERFACE_LIST"`
}

var DisplayIpVpnInstanceInterface_Template = `Value Required NAME (\S+)
Value Required ID (\d+)
Value List INTERFACE_LIST ([^,]+)

Start
  ^\s*Total
  ^\s*$$
  ^\s*VPN-Instance -> Continue.Record
  ^\s*VPN-Instance\s+Name\s+and\s+ID\s*:\s+${NAME}, ${ID}
  ^\s+Interface\s+Number
  ^\s+Interface\s+list\s*:\s+${INTERFACE_LIST},*
  ^\s{6,}${INTERFACE_LIST},*
  ^\s*Total
  ^\s*$$
  ^. -> Error
`