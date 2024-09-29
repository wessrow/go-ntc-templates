package mikrotik_routeros 

type IpRoutePrintDetail struct {
	Number	string	`json:"NUMBER"`
	Comment	string	`json:"COMMENT"`
	Active	string	`json:"ACTIVE"`
	Dynamic	string	`json:"DYNAMIC"`
	Protocol	string	`json:"PROTOCOL"`
	Discard_packet	string	`json:"DISCARD_PACKET"`
	Network	string	`json:"NETWORK"`
	Mask	string	`json:"MASK"`
	Pref_src	string	`json:"PREF_SRC"`
	Gateway	string	`json:"GATEWAY"`
	Gateway_status	string	`json:"GATEWAY_STATUS"`
	Check_gateway	string	`json:"CHECK_GATEWAY"`
	Distance	string	`json:"DISTANCE"`
	Type	string	`json:"TYPE"`
}

var IpRoutePrintDetail_Template = `Value NUMBER (\d+)
Value COMMENT (.*?)
Value ACTIVE (.)
Value DYNAMIC (.)
Value PROTOCOL (.)
Value DISCARD_PACKET (.)
Value Required NETWORK (\S+)
Value MASK (\d{1,2})
Value PREF_SRC (\S+)
Value GATEWAY (\S+)
Value GATEWAY_STATUS (.+?(?=[^\=]\S+\=))
Value CHECK_GATEWAY (\S+)
Value DISTANCE (\d+)
Value TYPE (\S+)

Start
  ^\s*Flags:\s+X\s+-\s+disabled,\s+A\s+-\s+active,\s+D\s+-\s+dynamic,\s+C\s+-\s+connect,\s+S\s+-\s+static,\s+r\s+-\s+rip,\s+b\s+-\s+bgp,\s+o\s+-\s+ospf,\s+m\s+-\s+mme,\s+B\s+-\s+blackhole,\s+U\s+-\s+unreachable,\s+P\s+-\s+prohibit\s*$$
  ^\s*${NUMBER}\s+${ACTIVE}${DYNAMIC}${PROTOCOL}${DISCARD_PACKET}\s+dst-address=${NETWORK}/${MASK}(?:\s+pref-src=${PREF_SRC})?(?:\s+gateway\=${GATEWAY})?(?:\s+gateway-status=${GATEWAY_STATUS})?.*?(?:\s+check-gateway=${CHECK_GATEWAY})?\s+distance=${DISTANCE}.*?(\s+ospf-type=${TYPE})?.*$$ -> Record
  ^\s*${NUMBER}\s+${ACTIVE}${DYNAMIC}${PROTOCOL}${DISCARD_PACKET}\s+;{3}\s*${COMMENT}\s*$$
  ^\s*dst-address=${NETWORK}/${MASK}(?:\s+pref-src=${PREF_SRC})?(?:\s+gateway\=${GATEWAY})?(?:\s+gateway-status=${GATEWAY_STATUS})?.*?(?:\s+check-gateway=${CHECK_GATEWAY})?\s+distance=${DISTANCE}.*?(\s+ospf-type=${TYPE})?.*$$ -> Record
  ^\s*$$
  ^. -> Error
`