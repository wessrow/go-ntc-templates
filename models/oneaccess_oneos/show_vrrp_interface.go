package oneaccess_oneos

type ShowVrrpInterface struct {
	Master_ip          string `json:"MASTER_IP"`
	Master_state       string `json:"MASTER_STATE"`
	Version            string `json:"VERSION"`
	Virtual_mac        string `json:"VIRTUAL_MAC"`
	Preempt_min_delay  string `json:"PREEMPT_MIN_DELAY"`
	Priority           string `json:"PRIORITY"`
	Group              string `json:"GROUP"`
	State              string `json:"STATE"`
	Virtual_netmask    string `json:"VIRTUAL_NETMASK"`
	Interface          string `json:"INTERFACE"`
	Virtual_ip         string `json:"VIRTUAL_IP"`
	Preempt            string `json:"PREEMPT"`
	Advertise_interval string `json:"ADVERTISE_INTERVAL"`
	Master_priority    string `json:"MASTER_PRIORITY"`
}

var ShowVrrpInterface_Template string = `Value Required INTERFACE (\S+(?:\s\S+)?)
Value GROUP (\d+)
Value STATE (\S+)
Value VERSION (\d+)
Value VIRTUAL_IP ([a-fA-F\d\.\:]+)
Value VIRTUAL_NETMASK ([a-fA-F\d\.\:]+)
Value VIRTUAL_MAC ([a-fA-F\d\.\:]+)
Value ADVERTISE_INTERVAL (\S+)
Value PREEMPT (\w+)
Value PREEMPT_MIN_DELAY (\d+)
Value PRIORITY (\d+)
Value MASTER_IP (\S+)
Value MASTER_PRIORITY (\S+)
Value MASTER_STATE (\S+)


Start
  ^\S+.*\s\-\sGroup -> Continue.Record
  ^${INTERFACE}\s\-\sGroup\s${GROUP}
  ^\s+State\sis\s${STATE}
  ^\s+Version\s${VERSION}
  ^\s+Virtual\sIP\saddress\s${VIRTUAL_IP},\sNetmask\s${VIRTUAL_NETMASK}
  ^\s+Virtual\sMAC\saddress\sis\s${VIRTUAL_MAC}
  ^\s+Advertisement\sinterval\sis\s${ADVERTISE_INTERVAL}\ssec
  ^\s+Preemption\sis\s${PREEMPT}(,\smin\sdelay\sis\s${PREEMPT_MIN_DELAY}\ssec)?
  ^\s+Priority\s${PRIORITY}
  ^\s+Master\srouter\sis\s${MASTER_IP}\s\(${MASTER_STATE}\),\spriority\sis\s${MASTER_PRIORITY}
  ^\s*$$
  ^. -> Error
`
