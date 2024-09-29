package models

type CiscoIosShowPowerStatus struct {
	Ps	string	`json:"PS"`
	Input_type	string	`json:"INPUT_TYPE"`
	Input_status	string	`json:"INPUT_STATUS"`
	Model	string	`json:"MODEL"`
	Type	string	`json:"TYPE"`
	Status	string	`json:"STATUS"`
	Fan_sensor	string	`json:"FAN_SENSOR"`
	Inline_status	string	`json:"INLINE_STATUS"`
}

var CiscoIosShowPowerStatus_Template = `####
# For 4500 switches
###
Value Required PS (\S+)
Value INPUT_TYPE (\S+)
Value INPUT_STATUS (off|good|err-disable)
Value Filldown MODEL (\S+)
Value TYPE (.+?)
Value STATUS (off|good|err-disable)
Value Filldown FAN_SENSOR (\w+)
Value Filldown INLINE_STATUS (good|n.a.)

Start
  ^Supply.*Model -> POWER
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is


POWER
  ^${PS}\s+${MODEL}\s+${TYPE}\s+${STATUS}\s+${FAN_SENSOR}\s+${INLINE_STATUS} -> Record
  ^${PS}\s{25}${INPUT_TYPE}\s+${INPUT_STATUS} -> Record
  # Capture just INPUT_STATUS when INPUT_TYPE field is blank
  ^${PS}\s{32}${INPUT_STATUS} -> Record
  ^$$ -> Clearall

`