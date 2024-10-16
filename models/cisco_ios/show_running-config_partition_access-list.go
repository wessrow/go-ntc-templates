package cisco_ios

type ShowRunningConfigPartitionAccessList struct {
	Acl_name             string   `json:"ACL_NAME"`
	Line_num             string   `json:"LINE_NUM"`
	Src_any              string   `json:"SRC_ANY"`
	Dst_port_range_start string   `json:"DST_PORT_RANGE_START"`
	Dst_host             string   `json:"DST_HOST"`
	Dst_network          string   `json:"DST_NETWORK"`
	Dst_port             string   `json:"DST_PORT"`
	Flags_match          string   `json:"FLAGS_MATCH"`
	Action               string   `json:"ACTION"`
	Protocol             string   `json:"PROTOCOL"`
	Src_network          string   `json:"SRC_NETWORK"`
	Src_port_match       string   `json:"SRC_PORT_MATCH"`
	Dst_wildcard         string   `json:"DST_WILDCARD"`
	Icmp_type            string   `json:"ICMP_TYPE"`
	Src_port             string   `json:"SRC_PORT"`
	Log                  string   `json:"LOG"`
	Acl_type             string   `json:"ACL_TYPE"`
	Src_port_range_start string   `json:"SRC_PORT_RANGE_START"`
	Dst_port_range_end   string   `json:"DST_PORT_RANGE_END"`
	Dst_port_match       string   `json:"DST_PORT_MATCH"`
	Tcp_flag             string   `json:"TCP_FLAG"`
	Time                 string   `json:"TIME"`
	Dst_any              string   `json:"DST_ANY"`
	Comment              []string `json:"COMMENT"`
	Src_host             string   `json:"SRC_HOST"`
	Src_port_range_end   string   `json:"SRC_PORT_RANGE_END"`
	Src_wildcard         string   `json:"SRC_WILDCARD"`
}

var ShowRunningConfigPartitionAccessList_Template string = `Value Required,Filldown ACL_NAME (\S+)
Value Filldown ACL_TYPE (standard|extended)
Value LINE_NUM (\d+)
Value List COMMENT (.*)
Value ACTION (permit|deny)
Value PROTOCOL (\S+)
Value SRC_HOST (\d+\.\d+\.\d+\.\d+)
Value SRC_ANY (any)
Value SRC_NETWORK (\d+\.\d+\.\d+\.\d+)
Value SRC_WILDCARD (\d+\.\d+\.\d+\.\d+)
Value SRC_PORT_MATCH (eq|neq|range|lt|gt)
Value SRC_PORT ((?<!range\s).+?)
Value SRC_PORT_RANGE_START ((?<=range\s)\S+)
Value SRC_PORT_RANGE_END (\S+)
Value DST_HOST (\d+\.\d+\.\d+\.\d+)
Value DST_ANY (any)
Value DST_NETWORK (\d+\.\d+\.\d+\.\d+)
Value DST_WILDCARD (\d+\.\d+\.\d+\.\d+)
Value DST_PORT_MATCH (eq|neq|range|lt|gt)
Value DST_PORT ((?<!range\s)\S+)
Value DST_PORT_RANGE_START ((?<=range\s)\S+)
Value DST_PORT_RANGE_END (\S+)
Value FLAGS_MATCH (match-all|match-any)
Value TCP_FLAG (((\+|-|)ack(\s*?)|(\+|-|)established(\s*?)|(\+|-|)fin(\s*?)|(\+|-|)fragments(\s*?)|(\+|-|)psh(\s*?)|(\+|-|)rst(\s*?)|(\+|-|)syn(\s*?)|urg(\s*?))+)
Value LOG (log-input|log)
Value ICMP_TYPE (echo|echo-reply|administratively-prohibited|unreachable|port-unreachable|redirect|router-advertisement|router-solicitation|packet-too-big|time-exceeded|ttl-exceeded|parameter-problem)
Value TIME (\S+)


Start
  # Clear all data to start new named ACL
  ^(ip\s+|)access-list -> Continue.Clearall
  # Record new named ACL; the record is required if a named ACL does not have any explicit policy entries.
  ^ip\s+access-list\s+${ACL_TYPE}\s+${ACL_NAME}\s* -> Record
  # Record named ACL Extended entry
  ^\s+(${LINE_NUM}\s+|)${ACTION}\s+${PROTOCOL}\s+(host\s+${SRC_HOST}|${SRC_ANY}|${SRC_NETWORK}\s+${SRC_WILDCARD})(\s+${SRC_PORT_MATCH}\s+(${SRC_PORT_RANGE_START}\s+${SRC_PORT_RANGE_END}|${SRC_PORT})|)\s+(host\s+${DST_HOST}|${DST_ANY}|${DST_NETWORK}\s+${DST_WILDCARD})(\s+${DST_PORT_MATCH}\s+(${DST_PORT_RANGE_START}\s+${DST_PORT_RANGE_END}|${DST_PORT})|\s+(${FLAGS_MATCH}\s+|)${TCP_FLAG}|)(\s+${ICMP_TYPE}|)(\s+${LOG}|)(\s+time-range\s+${TIME}|)\s*$$ -> Record
  # Record named ACL Standard entry
  ^\s+(${LINE_NUM}\s+|)${ACTION}\s+(${SRC_NETWORK}\s+${SRC_WILDCARD}|${SRC_ANY}|${SRC_HOST})(\s+${LOG}|)(\s+time-range\s+${TIME}|)\s* -> Record
  # Record named ACL Remark
  ^\s+(${LINE_NUM}\s+|)remark\s+${COMMENT}\s*
  # Record numbered ACL Extended entry
  ^access-list\s+${ACL_NAME}\s+${ACTION}\s+${PROTOCOL}\s+(host\s+${SRC_HOST}|${SRC_ANY}|${SRC_NETWORK}\s+${SRC_WILDCARD})(\s+${SRC_PORT_MATCH}\s+(${SRC_PORT_RANGE_START}\s+${SRC_PORT_RANGE_END}|${SRC_PORT})|)\s+(host\s+${DST_HOST}|${DST_ANY}|${DST_NETWORK}\s+${DST_WILDCARD})(\s+${DST_PORT_MATCH}\s+(${DST_PORT_RANGE_START}\s+${DST_PORT_RANGE_END}|${DST_PORT})|\s+(${FLAGS_MATCH}\s+|)${TCP_FLAG}|)(\s+${ICMP_TYPE}|)(\s+${LOG}|)(\s+time-range\s+${TIME}|)\s* -> Record
  # Record numbered ACL Standard entry
  ^access-list\s+${ACL_NAME}\s+${ACTION}\s+(${SRC_NETWORK}\s+${SRC_WILDCARD}|${SRC_ANY}|${SRC_HOST})(\s+${LOG}|)(\s+time-range\s+${TIME}|)\s* -> Record
  # Record numbered ACL Remark
  ^access-list\s+${ACL_NAME}\s+remark\s+${COMMENT}\s*
  # Catch all unuseful raw data
  ^(!\s*|$$|Building configuration.*|Current configuration.*|Configuration.*|end.*)
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is
  # Error out if raw data does not match any above rules.
  ^.* -> Error "Could not parse line:"

EOF
`
