package cisco_wlc

type SshShowPortSummary struct {
	Port            string `json:"PORT"`
	Admin_mode      string `json:"ADMIN_MODE"`
	Link_status     string `json:"LINK_STATUS"`
	Poe             string `json:"POE"`
	Type            string `json:"TYPE"`
	Stp_status      string `json:"STP_STATUS"`
	Physical_mode   string `json:"PHYSICAL_MODE"`
	Physical_status string `json:"PHYSICAL_STATUS"`
	Link_trap       string `json:"LINK_TRAP"`
}

var SshShowPortSummary_Template string = `Value PORT (\S+)
Value TYPE (\S+)
Value STP_STATUS (\S+)
Value ADMIN_MODE (\S+)
Value PHYSICAL_MODE (\w+\s?\w+)
Value PHYSICAL_STATUS (\w+\s?\w+)
Value LINK_STATUS (\S+)
Value LINK_TRAP (\S+)
Value POE (\S+(\s+\S+\s+\S+)?)

Start
  ^\s+STP\s+Admin\s+(Physical\s+){2}\s*(Link\s*){2}\s*$$
  ^\s*Pr\s+Type\s+Stat\s+(Mode\s+){2}\s*(Status\s+){2}\s*Trap\s+POE\s*$$
  #
  # Port Status and Modes (after the dashed line)
  ^\s*[-\s]+$$ -> Port_Status
  ^\s*$$
  ^. -> Error

Port_Status
  ^\s*${PORT}\s+${TYPE}\s+${STP_STATUS}\s+${ADMIN_MODE}\s+${PHYSICAL_MODE}\s+${PHYSICAL_STATUS}\s+${LINK_STATUS}\s+${LINK_TRAP}\s+${POE}\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`
