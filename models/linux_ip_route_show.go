package models

type LinuxIpRouteShow struct {
	Protocol	string	`json:"PROTOCOL"`
	Metric	string	`json:"METRIC"`
	Network	string	`json:"NETWORK"`
	Nexthop_ip	string	`json:"NEXTHOP_IP"`
	Nexthop_if	string	`json:"NEXTHOP_IF"`
	Scope	string	`json:"SCOPE"`
	Src	string	`json:"SRC"`
}

var LinuxIpRouteShow_Template = `Value PROTOCOL (\S+)
Value METRIC (\d+)
Value NETWORK (\S+)
Value NEXTHOP_IP (\d{1,3}.\d{1,3}.\d{1,3}.\d{1,3})
Value NEXTHOP_IF (\S+)
Value SCOPE (\S+)
Value SRC (\d{1,3}.\d{1,3}.\d{1,3}.\d{1,3})

Start
  # dhcp (discarded)
  ^.*proto dhcp.*
  ^\S+ dev \S+ proto dhcp .*
  # unreachable, broadcast, local (discarded)
  ^unreachable
  ^broadcast*
  ^local
  # default
  ^${NETWORK} via ${NEXTHOP_IP} dev ${NEXTHOP_IF} -> Record
  # local (network)
  ^${NETWORK} dev ${NEXTHOP_IF} proto ${PROTOCOL} scope ${SCOPE} src ${SRC} -> Record
  # static
  ^${NETWORK} via ${NEXTHOP_IP} dev ${NEXTHOP_IF} -> Record
  ^. -> Error

`