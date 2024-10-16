package huawei_vrp

type DisplayAclAll struct {
	Protocol                  string `json:"PROTOCOL"`
	Source                    string `json:"SOURCE"`
	Source_port_modifier      string `json:"SOURCE_PORT_MODIFIER"`
	Source_port_range         string `json:"SOURCE_PORT_RANGE"`
	Acl_name                  string `json:"ACL_NAME"`
	Acl_number                string `json:"ACL_NUMBER"`
	Sn                        string `json:"SN"`
	Action                    string `json:"ACTION"`
	Matches                   string `json:"MATCHES"`
	Description               string `json:"DESCRIPTION"`
	Destination               string `json:"DESTINATION"`
	Destination_port_modifier string `json:"DESTINATION_PORT_MODIFIER"`
	Destination_port_range    string `json:"DESTINATION_PORT_RANGE"`
}

var DisplayAclAll_Template string = `Value Filldown ACL_NAME (.*?)
Value Filldown,Required ACL_NUMBER (\d+)
Value SN (\d+)
Value ACTION (deny|permit)
Value PROTOCOL (\d+|icmp(v6)?|tcp|udp|gre|igmp|ip(v6)?|ipinip|ospf)
Value SOURCE (((\d+.){3}\d+|[0-9a-fA-F]{1,4}([0-9a-fA-F]{0,4}:){1,7}[0-9a-fA-F]{0,4})[\/\s]((\d+.){3}\d+|\d+))
Value SOURCE_PORT_MODIFIER (eq|lt|gt|range)
Value SOURCE_PORT_RANGE ((\S+)(\s(\S+))?)
Value DESTINATION (((\d+.){3}\d+|[0-9a-fA-F]{1,4}([0-9a-fA-F]{0,4}:){1,7}[0-9a-fA-F]{0,4})[\/\s]((\d+.){3}\d+|\d+))
Value DESTINATION_PORT_MODIFIER (eq|lt|gt|range)
Value DESTINATION_PORT_RANGE ((\S+)(\s(\S+))?)
Value MATCHES (\d+)
Value DESCRIPTION (.*)

Start
  ^\s*rule\s\d+\s(deny|permit) -> Continue.Record
  ^\S+\sIPv6\sACL\s -> Continue.Record
  ^\S+\sACL\s -> Continue.Record
  ^\s*Total\squantity\sof\snonempty\sACL\snumber\sis\s\d+\s*$$
  ^\s*Total\snonempty\sacl6\snumber\sis\s\d+\s*$$
  ^\S+\sIPv6\sACL\s${ACL_NUMBER}(\sname\s${ACL_NAME})?,\s\d+\srules?.*$$
  ^\S+\sACL\s(${ACL_NAME}\s)?${ACL_NUMBER},\s\d+\srules?.*$$
  ^\s*Acl's\sstep\sis\s\d+\s*$$
  ^\s*rule\s${SN}\s${ACTION}(\s${PROTOCOL})?(\ssource\s${SOURCE})?(\ssource-port\s${SOURCE_PORT_MODIFIER}\s${SOURCE_PORT_RANGE})?(\sdestination\s${DESTINATION})?(\sdestination-port\s${DESTINATION_PORT_MODIFIER}\s${DESTINATION_PORT_RANGE})?(\s\(${MATCHES}\smatch(es)?\))?\s*$$
  ^\s*rule\s\d+\sdescription\s"${DESCRIPTION}"\s*$$
  ^\s*$$
  ^. -> Error
`
