package cisco_ios

type ShowIpAccessLists struct {
	Action                        string `json:"ACTION"`
	Src_host                      string `json:"SRC_HOST"`
	Src_port_range_end            string `json:"SRC_PORT_RANGE_END"`
	Dst_port_range_start          string `json:"DST_PORT_RANGE_START"`
	Time                          string `json:"TIME"`
	Src_port_match                string `json:"SRC_PORT_MATCH"`
	Icmp_type                     string `json:"ICMP_TYPE"`
	Src_any                       string `json:"SRC_ANY"`
	Dst_any                       string `json:"DST_ANY"`
	Dst_network                   string `json:"DST_NETWORK"`
	Dst_port_range_end            string `json:"DST_PORT_RANGE_END"`
	Service_object_group_name     string `json:"SERVICE_OBJECT_GROUP_NAME"`
	Log                           string `json:"LOG"`
	State                         string `json:"STATE"`
	Src_port                      string `json:"SRC_PORT"`
	Acl_name                      string `json:"ACL_NAME"`
	Src_wildcard                  string `json:"SRC_WILDCARD"`
	Dst_host                      string `json:"DST_HOST"`
	Dst_wildcard                  string `json:"DST_WILDCARD"`
	Dst_port                      string `json:"DST_PORT"`
	Src_network                   string `json:"SRC_NETWORK"`
	Src_port_range_start          string `json:"SRC_PORT_RANGE_START"`
	Dst_network_object_group_name string `json:"DST_NETWORK_OBJECT_GROUP_NAME"`
	Tcp_flag                      string `json:"TCP_FLAG"`
	Log_tag                       string `json:"LOG_TAG"`
	Protocol                      string `json:"PROTOCOL"`
	Src_network_object_group_name string `json:"SRC_NETWORK_OBJECT_GROUP_NAME"`
	Flags_match                   string `json:"FLAGS_MATCH"`
	Acl_type                      string `json:"ACL_TYPE"`
	Line_num                      string `json:"LINE_NUM"`
	Dst_port_match                string `json:"DST_PORT_MATCH"`
	Matches                       string `json:"MATCHES"`
}

var ShowIpAccessLists_Template string = `Value Required,Filldown ACL_TYPE (Standard|Extended)
Value Required,Filldown ACL_NAME (\S+)
Value LINE_NUM (\d+)
Value ACTION (permit|deny)
Value PROTOCOL (\S+)
Value SRC_HOST (\d+\.\d+\.\d+\.\d+)
Value SRC_ANY (any)
Value SRC_NETWORK (\d+\.\d+\.\d+\.\d+)
Value SRC_WILDCARD (\d+\.\d+\.\d+\.\d+)
Value SRC_NETWORK_OBJECT_GROUP_NAME (\S+)
Value SRC_PORT_MATCH (eq|neq|precedence|range|tos|lt|gt)
Value SRC_PORT ((?<!range\s).+?)
Value SRC_PORT_RANGE_START ((?<=range\s)\S+)
Value SRC_PORT_RANGE_END (\S+)
Value DST_HOST (\d+\.\d+\.\d+\.\d+)
Value DST_ANY (any)
Value DST_NETWORK (\d+\.\d+\.\d+\.\d+)
Value DST_WILDCARD (\d+\.\d+\.\d+\.\d+)
Value DST_NETWORK_OBJECT_GROUP_NAME (\S+)
Value DST_PORT_MATCH (eq|neq|precedence|range|tos|lt|gt)
Value DST_PORT ((?<!range\s).+?)
Value DST_PORT_RANGE_START ((?<=range\s)\S+)
Value DST_PORT_RANGE_END (\S+)
Value SERVICE_OBJECT_GROUP_NAME (\S+)
Value FLAGS_MATCH (match-all|match-any)
Value TCP_FLAG (((\+|-|)ack(\s*?)|(\+|-|)established(\s*?)|(\+|-|)fin(\s*?)|(\+|-|)fragments(\s*?)|(\+|-|)psh(\s*?)|(\+|-|)rst(\s*?)|(\+|-|)syn(\s*?)|urg(\s*?))+)
Value LOG (log-input|log)
Value LOG_TAG (\S+)
Value ICMP_TYPE (administratively-prohibited|echo|echo-reply|mask-request|packet-too-big|parameter-problem|port-unreachable|redirect|router-advertisement|router-solicitation|time-exceeded|ttl-exceeded|unreachable)
Value TIME (\S+)
Value STATE (inactive|active)
Value MATCHES (\d+)

Start
  ^(Standard|Extended) -> Continue.Clearall
  ^${ACL_TYPE}\s+IP\s+access\s+list\s+${ACL_NAME}\s* -> Record
  ^\s+${LINE_NUM}\s+${ACTION}\s+(${PROTOCOL}|${SERVICE_OBJECT_GROUP_NAME})\s+(host\s+${SRC_HOST}|${SRC_ANY}|${SRC_NETWORK}\s+${SRC_WILDCARD}|object-group\s+${SRC_NETWORK_OBJECT_GROUP_NAME})(\s+${SRC_PORT_MATCH}\s+|)(${SRC_PORT_RANGE_START}\s+${SRC_PORT_RANGE_END}|${SRC_PORT}|)\s+(host\s+${DST_HOST}|${DST_ANY}|${DST_NETWORK}\s+${DST_WILDCARD}|object-group\s+${DST_NETWORK_OBJECT_GROUP_NAME})(\s+${DST_PORT_MATCH}\s+(${DST_PORT_RANGE_START}\s+${DST_PORT_RANGE_END}|${DST_PORT}|)|\s+(${FLAGS_MATCH}\s+|)${TCP_FLAG}|)(\s+${ICMP_TYPE}|)(\s+${LOG}|)(\s+time-range\s+${TIME}\s+\(${STATE}\)|)(?:\s+\(${MATCHES}\s+\S+\)|)(\s+\(tag\s+=\s+${LOG_TAG}\)|)\s*$$ -> Record
  ^\s+${LINE_NUM}\s+${ACTION}\s+(${SRC_NETWORK},\s+wildcard\s+bits\s+${SRC_WILDCARD}|${SRC_HOST}|${SRC_ANY})(\s+${LOG}|)(\s+time-range\s+${TIME}\s+\(${STATE}\)|)(?:\s+\(${MATCHES}\s+\S+\)|)(\s+\(tag\s+=\s+${LOG_TAG}\)|)\s*$$ -> Record
  ^\s*$$
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is
  ^.* -> Error "Could not parse line:"

EOF
`
