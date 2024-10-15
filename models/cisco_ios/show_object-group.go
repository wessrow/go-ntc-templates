package cisco_ios

type ShowObjectGroup struct {
	Host_range_start string `json:"HOST_RANGE_START"`
	Host_range_end   string `json:"HOST_RANGE_END"`
	Protocol         string `json:"PROTOCOL"`
	Type             string `json:"TYPE"`
	Any              string `json:"ANY"`
	Port_range_end   string `json:"PORT_RANGE_END"`
	Network          string `json:"NETWORK"`
	Netmask          string `json:"NETMASK"`
	Nested_groups    string `json:"NESTED_GROUPS"`
	Host             string `json:"HOST"`
	Port             string `json:"PORT"`
	Port_range_start string `json:"PORT_RANGE_START"`
	Name             string `json:"NAME"`
	Description      string `json:"DESCRIPTION"`
	Port_match       string `json:"PORT_MATCH"`
	Icmp_type        string `json:"ICMP_TYPE"`
}

var ShowObjectGroup_Template string = `Value Required,Filldown TYPE (Service|Network)
Value Required,Filldown NAME (\S+)
Value DESCRIPTION (.+)
Value NESTED_GROUPS (\S+)
Value ANY (any)
Value HOST (\d+\.\d+\.\d+\.\d+)
Value HOST_RANGE_START (\d+\.\d+\.\d+\.\d+)
Value HOST_RANGE_END (\d+\.\d+\.\d+\.\d+)
Value NETWORK (\d+\.\d+\.\d+\.\d+)
Value NETMASK (\d+\.\d+\.\d+\.\d+)
Value PROTOCOL (\S+)
Value PORT_MATCH (eq|neq|range|lt|gt)
Value PORT ((?<!range\s)\S+)
Value PORT_RANGE_START ((?<=range\s)\S+)
Value PORT_RANGE_END (\S+)
Value ICMP_TYPE (echo|echo-reply|administratively-prohibited|unreachable|port-unreachable|redirect|router-advertisement|router-solicitation|packet-too-big|time-exceeded|ttl-exceeded|parameter-problem)

Start
  ^(Service|Network) -> Continue.Clearall
  ^${TYPE}\s+object\s+group\s+${NAME}\s*$$ -> Record
  ^\s+Description\s+${DESCRIPTION}$$ -> Record
  ^\s+group-object\s+${NESTED_GROUPS}\s*$$ -> Record
  ^\s+(host\s+${HOST}|range\s+${HOST_RANGE_START}\s+${HOST_RANGE_END}|${ANY}|${NETWORK}\s+${NETMASK})\s*$$ -> Record
  ^\s+icmp\s+${ICMP_TYPE}\s*$$ -> Record
  ^\s+${PROTOCOL}\s+${PORT_MATCH}\s+(${PORT_RANGE_START}\s+${PORT_RANGE_END}|${PORT})\s*$$ -> Record
  ^\s+${PROTOCOL}\s*$$ -> Record
  ^\s*$$
  ^. -> Error

EOF

`
