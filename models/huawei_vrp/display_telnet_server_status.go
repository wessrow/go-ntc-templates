package huawei_vrp 

type DisplayTelnetServerStatus struct {
	Telnet_ipv4_server	string	`json:"TELNET_IPV4_SERVER"`
	Telnet_ipv6_server	string	`json:"TELNET_IPV6_SERVER"`
	Telnet_service_port	string	`json:"TELNET_SERVICE_PORT"`
	Telnet_server_source_address	string	`json:"TELNET_SERVER_SOURCE_ADDRESS"`
	Acl4_number	string	`json:"ACL4_NUMBER"`
	Acl6_number	string	`json:"ACL6_NUMBER"`
}

var DisplayTelnetServerStatus_Template = `Value TELNET_IPV4_SERVER (\S+)
Value TELNET_IPV6_SERVER (\S+)
Value TELNET_SERVICE_PORT (\d+)
Value TELNET_SERVER_SOURCE_ADDRESS (\S+)
Value ACL4_NUMBER (\d+)
Value ACL6_NUMBER (\d+)

Start
  ^\s*TELNET\s*IPv4\s*server\s*:\s*${TELNET_IPV4_SERVER}
  ^\s*TELNET\s*IPv6\s*server\s*:\s*${TELNET_IPV6_SERVER}
  ^\s*TELNET\s*server\s*port\s*:\s*${TELNET_SERVICE_PORT}
  ^\s*TELNET\s*server\s*source\s*address\s*:\s*${TELNET_SERVER_SOURCE_ADDRESS}
  ^\s*ACL4\s*number\s*:\s*${ACL4_NUMBER}
  ^\s*ACL6\s*number\s*:\s*${ACL6_NUMBER}
  ^\s*$$
  ^. -> Error
`