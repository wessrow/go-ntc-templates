package cisco_ios

type ShowVrrpBrief struct {
	Master_ip_address  string `json:"MASTER_IP_ADDRESS"`
	Virtual_ip_address string `json:"VIRTUAL_IP_ADDRESS"`
	Addr_family        string `json:"ADDR_FAMILY"`
	Time               string `json:"TIME"`
	Own                string `json:"OWN"`
	Preempt            string `json:"PREEMPT"`
	State              string `json:"STATE"`
	Interface          string `json:"INTERFACE"`
	Group              string `json:"GROUP"`
	Priority           string `json:"PRIORITY"`
}

var ShowVrrpBrief_Template string = `Value INTERFACE (\S+)
Value GROUP (\d+)
Value ADDR_FAMILY (\S+)
Value PRIORITY (\d+)
Value TIME (\d+)
Value OWN (.?)
Value PREEMPT (.)
Value STATE (\w+)
Value MASTER_IP_ADDRESS ([a-fA-F\d\.\:]+)
Value VIRTUAL_IP_ADDRESS ([a-fA-F\d\.\:]+)

Start
  ^Interface\s+Grp\s+Pri\s+Time\s+Own\s+Pre\s+State\s+Master addr\s+Group addr\s*$$ -> Vrrp
  ^\s*Interface\s+Grp\s+A-F\s+Pri\s+Time\s+Own\s+Pre\s+State\s+Master addr\/Group addr\s*$$ -> Vrrpv3
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is

Vrrp
  # VRRP Legacy (VRRPv2 for IPv4)
  ^${INTERFACE}\s+${GROUP}\s*${PRIORITY}\s+${TIME}\s+${OWN}\s+${PREEMPT}\s+${STATE}\s+${MASTER_IP_ADDRESS}\s+${VIRTUAL_IP_ADDRESS} -> Record
  ^${INTERFACE}\s*$$
  ^\s*${GROUP}\s+${PRIORITY}\s+${TIME}\s+${OWN}\s+${PREEMPT}\s+${STATE}\s+${MASTER_IP_ADDRESS}\s+${VIRTUAL_IP_ADDRESS} -> Record
  ^\s*$$
  ^. -> Error

Vrrpv3
  # VRRP Unified (VRRPv3 for IPv4 and IPv6)
  # - ignoring the "local" designation since we can use the STATE value to determine who the master is
  ^\s*${INTERFACE}\s+${GROUP}\s+${ADDR_FAMILY}\s+${PRIORITY}\s+${TIME}\s+${OWN}\s+${PREEMPT}\s+${STATE}\s+${MASTER_IP_ADDRESS}\(local\)\s+${VIRTUAL_IP_ADDRESS} -> Record
  ^\s*${INTERFACE}\s+${GROUP}\s+${ADDR_FAMILY}\s+${PRIORITY}\s+${TIME}\s+${OWN}\s+${PREEMPT}\s+${STATE}\s+${MASTER_IP_ADDRESS}\s+${VIRTUAL_IP_ADDRESS} -> Record
  # rows that roll over to a second line
  ^\s*${INTERFACE}\s*$$
  ^\s*${GROUP}\s+${ADDR_FAMILY}\s+${PRIORITY}\s+${TIME}\s+${OWN}\s+${PREEMPT}\s+${STATE}\s+${MASTER_IP_ADDRESS}\(local\)\s+${VIRTUAL_IP_ADDRESS} -> Record
  ^\s*${INTERFACE}\s*$$
  ^\s*${GROUP}\s+${ADDR_FAMILY}\s+${PRIORITY}\s+${TIME}\s+${OWN}\s+${PREEMPT}\s+${STATE}\s+${MASTER_IP_ADDRESS}\s+${VIRTUAL_IP_ADDRESS} -> Record
  ^\s*$$
  ^. -> Error
`
