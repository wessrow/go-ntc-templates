package huawei_vrp

type DisplayDevice struct {
	Slot            string `json:"SLOT"`
	Card            string `json:"CARD"`
	Pid             string `json:"PID"`
	Online_status   string `json:"ONLINE_STATUS"`
	Power_status    string `json:"POWER_STATUS"`
	Register_status string `json:"REGISTER_STATUS"`
	Alarm_status    string `json:"ALARM_STATUS"`
	Ha_status       string `json:"HA_STATUS"`
}

var DisplayDevice_Template string = `Value Filldown SLOT (\w+)
Value CARD (\S+)
Value Required PID (\S+(\s+\S+)?)
Value ONLINE_STATUS (\w+)
Value POWER_STATUS (\S+)
Value REGISTER_STATUS (\w+)
Value ALARM_STATUS (\S+)
Value HA_STATUS (\w+)

Start
  ^\*.*$$
  ^\-+$$
  ^(\S+\s+)?Device\s+status\S+$$ -> DeviceStatus
  ^Slot\/Card.+$$ -> Card
  ^. -> Error

# device card
Card
  ^${SLOT}/${CARD}\s+${PID}\s+${ONLINE_STATUS}\s+${POWER_STATUS}\s+${REGISTER_STATUS}\s+${ALARM_STATUS}\s*$$ -> Record
  ^\-+$$
  ^. -> Error

# device and device all
DeviceStatus
  ^\-+$$
  ^Slot.+$$
  ^\s+${CARD}\s+${PID}\s+${ONLINE_STATUS}\s+${POWER_STATUS}\s+${REGISTER_STATUS}\s+${ALARM_STATUS}\s+${HA_STATUS}\s*$$ -> Record
  ^${SLOT}\s+${CARD}\s+\*?${PID}\s+${ONLINE_STATUS}\s+${POWER_STATUS}\s+${REGISTER_STATUS}\s+${ALARM_STATUS}\s+${HA_STATUS}\s*$$ -> Record
  ^. -> Error
`
