package oneaccess_oneos

type ShowIpAccessLists struct {
	Acl_name             string `json:"ACL_NAME"`
	Protocol             string `json:"PROTOCOL"`
	Src_port_match       string `json:"SRC_PORT_MATCH"`
	Src_port_range_start string `json:"SRC_PORT_RANGE_START"`
	Src_port_range_end   string `json:"SRC_PORT_RANGE_END"`
	Dst_network          string `json:"DST_NETWORK"`
	Matches              string `json:"MATCHES"`
	Src_host             string `json:"SRC_HOST"`
	Dst_host             string `json:"DST_HOST"`
	Dst_port             string `json:"DST_PORT"`
	Dst_port_range_start string `json:"DST_PORT_RANGE_START"`
	Dst_port_match       string `json:"DST_PORT_MATCH"`
	Line_num             string `json:"LINE_NUM"`
	Protocol_nr          string `json:"PROTOCOL_NR"`
	Src_any              string `json:"SRC_ANY"`
	Src_network          string `json:"SRC_NETWORK"`
	Src_port             string `json:"SRC_PORT"`
	Dst_any              string `json:"DST_ANY"`
	Dst_wildcard         string `json:"DST_WILDCARD"`
	Acl_type             string `json:"ACL_TYPE"`
	Action               string `json:"ACTION"`
	Src_wildcard         string `json:"SRC_WILDCARD"`
	Dst_port_range_end   string `json:"DST_PORT_RANGE_END"`
	Log                  string `json:"LOG"`
}

var ShowIpAccessLists_Template string = `Value Required,Filldown ACL_TYPE (standard|extended)
Value Required,Filldown ACL_NAME (\S+)
Value LINE_NUM (\d+)
Value ACTION (permit|deny)
Value PROTOCOL (\S+)
Value PROTOCOL_NR (\d+)
Value SRC_HOST (\d+\.\d+\.\d+\.\d+)
Value SRC_ANY (any)
Value SRC_NETWORK (\d+\.\d+\.\d+\.\d+)
Value SRC_WILDCARD (\d+\.\d+\.\d+\.\d+)
#Value SRC_NETWORK_OBJECT_GROUP_NAME (\S+)
Value SRC_PORT_MATCH (eq|neq|precedence|range|tos|lt|gt)
Value SRC_PORT (\d+|[a-z]+)
Value SRC_PORT_RANGE_START ((?<=range\s)\S+)
Value SRC_PORT_RANGE_END (\S+)
Value DST_HOST (\d+\.\d+\.\d+\.\d+)
Value DST_ANY (any)
Value DST_NETWORK (\d+\.\d+\.\d+\.\d+)
Value DST_WILDCARD (\d+\.\d+\.\d+\.\d+)
#Value DST_NETWORK_OBJECT_GROUP_NAME (\S+)
Value DST_PORT_MATCH (eq|neq|precedence|range|tos|lt|gt)
Value DST_PORT (\d+|[a-z]+)
Value DST_PORT_RANGE_START ((?<=range\s)\S+)
Value DST_PORT_RANGE_END (\S+)
#Value SERVICE_OBJECT_GROUP_NAME (\S+)
#Value FLAGS_MATCH (match-all|match-any)
#Value TCP_FLAG (((\+|-|)ack(\s*?)|(\+|-|)established(\s*?)|(\+|-|)fin(\s*?)|(\+|-|)fragments(\s*?)|(\+|-|)psh(\s*?)|(\+|-|)rst(\s*?)|(\+|-|)syn(\s*?)|urg(\s*?))+)
Value LOG (log)
#Value LOG_TAG (\S+)
#Value ICMP_TYPE (administratively-prohibited|echo|echo-reply|mask-request|packet-too-big|parameter-problem|port-unreachable|redirect|router-advertisement|router-solicitation|time-exceeded|ttl-exceeded|unreachable)
#Value TIME (\S+)
#Value STATE (inactive|active)
Value MATCHES (\d+)

Start
  ^ip\saccess-list\s(standard|extended) -> Continue.Clearall
  ^ip\saccess-list\s${ACL_TYPE}\s${ACL_NAME}\s* -> Record ONEOS6
  ^IP\saccess\slist\s(standard|extended) -> Continue.Clearall
  ^IP\saccess\slist\s${ACL_TYPE}\s${ACL_NAME}\s* -> Record ONEOS5
  ^interface.*
  ^For info on.*
  ^\s*$$
  ^.* -> Error "Could not parse line:"

ONEOS6
  ^\s+${LINE_NUM}\s+${ACTION}(?:\s+${PROTOCOL})?\s+(?:${PROTOCOL_NR}\s)?${SRC_NETWORK}\s+${SRC_WILDCARD}(${SRC_PORT}|)(\s+${DST_NETWORK}\s+${DST_WILDCARD}|)(\s+${DST_PORT}|)(\s+${LOG}|)(?:\s+\(${MATCHES}\s+matches\)|)\s* -> Record  
  ^\s*$$ -> Start
  ^.* -> Error "Could not parse line:"

ONEOS5
  ^${ACTION}\s+(${PROTOCOL}\s+|)(host\s+${SRC_HOST}|${SRC_ANY}|${SRC_NETWORK}\s+${SRC_WILDCARD})(\s+${SRC_PORT_MATCH}\s+|)(${SRC_PORT_RANGE_START}\s+${SRC_PORT_RANGE_END}|${SRC_PORT}|)\s+(host\s+${DST_HOST}|${DST_ANY}|${DST_NETWORK}\s+${DST_WILDCARD})(\s+${DST_PORT_MATCH}\s+(${DST_PORT_RANGE_START}\s+${DST_PORT_RANGE_END}|${DST_PORT}|)|)(\s+${LOG}|)(?:\s+\(${MATCHES}\s+matches\)|) -> Record
  ^\s*$$ -> Start
  ^.* -> Error "Could not parse line:"

EOF
`
