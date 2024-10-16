package cisco_ios

type ShowIpv6AccessLists struct {
	Routing              string `json:"ROUTING"`
	Src_port_range_start string `json:"SRC_PORT_RANGE_START"`
	Src_port_range_end   string `json:"SRC_PORT_RANGE_END"`
	Action               string `json:"ACTION"`
	Ipv6_destination     string `json:"IPV6_DESTINATION"`
	Sequence             string `json:"SEQUENCE"`
	Log                  string `json:"LOG"`
	Authentification     string `json:"AUTHENTIFICATION"`
	Acl_name             string `json:"ACL_NAME"`
	Protocol             string `json:"PROTOCOL"`
	Ipv6_source          string `json:"IPV6_SOURCE"`
	Src_port_match       string `json:"SRC_PORT_MATCH"`
	Port_number          string `json:"PORT_NUMBER"`
}

var ShowIpv6AccessLists_Template string = `Value Required,Filldown ACL_NAME (\S+)
Value ACTION (permit|deny)
Value PROTOCOL (\S+)
Value IPV6_SOURCE ([A-Za-z0-9:]+(\/(?:\d{1,3})|))
Value IPV6_DESTINATION ([A-Za-z0-9:]+(\/(?:\d{1,3})|))
Value SRC_PORT_MATCH (eq|neq|precedence|range|tos|lt|gt|established)
Value SEQUENCE (\d+)
Value PORT_NUMBER (\d+)
Value LOG (log)
Value AUTHENTIFICATION (auth)
Value ROUTING (routing)
Value SRC_PORT_RANGE_START (\S+)
Value SRC_PORT_RANGE_END (\S+)


Start
  ^IPv6\s+access\s+list\s+${ACL_NAME}\s* -> Continue
  ^\s+${ACTION}\s+${PROTOCOL}\s+${IPV6_SOURCE}\s+${IPV6_DESTINATION}\s+(${SRC_PORT_MATCH}\s|)sequence\s${SEQUENCE}\s*$$ -> Record
  ^\s+${ACTION}\s+${PROTOCOL}\s+(host\s+|)${IPV6_SOURCE}\s+(host\s+|)(${IPV6_DESTINATION}\s+|)(${SRC_PORT_MATCH}\s|range\s${SRC_PORT_RANGE_START}\s${SRC_PORT_RANGE_END}\s|)(${LOG}\s+|)(${AUTHENTIFICATION}\s+|)(${ROUTING}\s+|)(${PORT_NUMBER}\s+|)sequence\s${SEQUENCE}\s*$$ -> Record

EOF`
