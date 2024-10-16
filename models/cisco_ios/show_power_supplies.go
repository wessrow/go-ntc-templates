package cisco_ios

type ShowPowerSupplies struct {
	Ps_avail  string `json:"PS_AVAIL"`
	Ps_needed string `json:"PS_NEEDED"`
}

var ShowPowerSupplies_Template string = `####
# For 4500 switches
###
Value PS_NEEDED (\d)
Value PS_AVAIL (\d)

Start
  ^Power\ssupplies\sneeded\sby\ssystem\s+:\s${PS_NEEDED} -> Continue
  ^Power\ssupplies\scurrently\savailable\s+:\s${PS_AVAIL}
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is
`
