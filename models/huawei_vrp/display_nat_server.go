package huawei_vrp

type DisplayNatServer struct {
	Protocol       string `json:"PROTOCOL"`
	Vpn            string `json:"VPN"`
	Acl            string `json:"ACL"`
	Vrrp_id        string `json:"VRRP_ID"`
	Description    string `json:"DESCRIPTION"`
	Interface      string `json:"INTERFACE"`
	Global_ip_port string `json:"GLOBAL_IP_PORT"`
	Inside_ip_port string `json:"INSIDE_IP_PORT"`
}

var DisplayNatServer_Template string = `Value Filldown INTERFACE (\S+)
Value Required GLOBAL_IP_PORT (.+)
Value INSIDE_IP_PORT ((\d{1,3}.){3}\d{1,3}\/\d+(\(\w+\))?)
Value PROTOCOL (\d+(\(\w+\))?|\w+)
Value VPN (\S+)
Value ACL (\S+)
Value VRRP_ID (\S+)
Value DESCRIPTION (.+)

Start
  ^\s*Nat\s+Server\s+Information:\s*$$
  ^\s+Interface\s+:\s+${INTERFACE}\s*$$
  ^\s+Global\s+IP\/Port\s+:\s+${GLOBAL_IP_PORT}\s*$$
  ^\s+Inside\s+IP\/Port\s+:\s+${INSIDE_IP_PORT}\s*$$
  ^\s+Protocol\s+:\s+${PROTOCOL}\s*$$
  ^\s+VPN\s+instance-name\s+:\s+${VPN}\s*$$
  ^\s+Acl\s+number\s+:\s+${ACL}\s*$$
  ^\s+Vrrp\s+id\s+:\s+${VRRP_ID}\s*$$
  ^\s+Description\s+:\s+${DESCRIPTION}\s*$$ -> Record
  ^\s*Total\s+:\s+\d+\s*$$
  ^\s*$$
  ^.*$$ -> Error
`
