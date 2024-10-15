package cisco_xr

type ShowHsrp struct {
	Prio         string `json:"PRIO"`
	Preempt      string `json:"PREEMPT"`
	State        string `json:"STATE"`
	Active_addr  string `json:"ACTIVE_ADDR"`
	Standby_addr string `json:"STANDBY_ADDR"`
	Group_addr   string `json:"GROUP_ADDR"`
	Interface    string `json:"INTERFACE"`
	Group        string `json:"GROUP"`
}

var ShowHsrp_Template string = `Value INTERFACE (\S+)
Value GROUP (\d+)
Value PRIO (\d+)
Value PREEMPT (P)
Value STATE (\w+)
Value ACTIVE_ADDR (\S+)
Value STANDBY_ADDR (\S+)
Value GROUP_ADDR ((?:\d+\.?){4}|(?:[a-f0-9]*:){1,7}[a-f0-9]*)

Start
  ^${INTERFACE}\s+${GROUP}\s+${PRIO}\s+(?:${PREEMPT}\s+)?${STATE}\s+${ACTIVE_ADDR}\s+${STANDBY_ADDR}\s+${GROUP_ADDR}\s*$$ -> Record
  ^IPv6\s+Groups: -> IPv6Groups
  ^Interface\s+Grp\s+Pri\s+P\s+State\s+Active\s+addr\s+Standby\s+addr\s+\s+Group\s+addr\s*$$
  ^IPv4\s+Groups:
  ^\s+P\s+indicates
  ^\s+|\s*$$
  ^\S+\s+\S+\s+\d+\s+\d+:\d+:\d+\.\d+\s+\S+\s*$$
  ^\s*$$
  ^. -> Error

IPv6Groups
  # multiline IPv6 HSRP group
  ^${INTERFACE}\s+${GROUP}\s+${PRIO}\s+(?:${PREEMPT}\s+)?${STATE}\s+${ACTIVE_ADDR}(?:\s+${STANDBY_ADDR})?\s*$$
  ^(?:\s){45,50}${STANDBY_ADDR}\s*$$
  ^(?:\s){62,}${GROUP_ADDR}\s*$$ -> Record
  # single line IPv6 HSRP group
  ^${INTERFACE}\s+${GROUP}\s+${PRIO}\s+(?:${PREEMPT}\s+)?${STATE}\s+${ACTIVE_ADDR}\s+${STANDBY_ADDR}(?:\s+${GROUP_ADDR})?\s*$$ -> Record
  ^Interface\s+Grp\s+Pri\s+P\s+State\s+Active\s+addr\s+Standby\s+addr\s+\s+Group\s+addr\s*$$
  ^\s+P\s+indicates
  ^\s+|\s*$$
  ^. -> Error
`
