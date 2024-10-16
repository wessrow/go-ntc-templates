package cisco_ios

type ShowEnvironmentPowerAll struct {
	Ps_module string `json:"PS_MODULE"`
	Pid       string `json:"PID"`
	Serial    string `json:"SERIAL"`
	Status    string `json:"STATUS"`
	Sys_pwr   string `json:"SYS_PWR"`
	Poe_pwr   string `json:"POE_PWR"`
	Watts     string `json:"WATTS"`
}

var ShowEnvironmentPowerAll_Template string = `####
# For 3850 switches, does not work on 3750 switches
###
Value PS_MODULE (\w+)
Value PID (\S+)
Value SERIAL (\w+)
Value STATUS (.+?)
Value SYS_PWR (Good|Bad)
Value POE_PWR (\w+)
Value WATTS (\w+)

Start
  ^${PS_MODULE}\s+${PID}\s+${SERIAL}\s+${STATUS}\s+${SYS_PWR}\s+${POE_PWR}\s+${WATTS} -> Record
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is
`
