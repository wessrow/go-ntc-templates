package models

type CiscoNxosShowEnvironmentTemperature struct {
	Current_value	string	`json:"CURRENT_VALUE"`
	Major_threshold	string	`json:"MAJOR_THRESHOLD"`
	Minor_threshold	string	`json:"MINOR_THRESHOLD"`
	Status	string	`json:"STATUS"`
	Module	string	`json:"MODULE"`
	Sensor	string	`json:"SENSOR"`
}

var CiscoNxosShowEnvironmentTemperature_Template = `Value CURRENT_VALUE (\d+)
Value MAJOR_THRESHOLD (\d+)
Value MINOR_THRESHOLD (\d+)
Value STATUS (\S+)
Value MODULE (\S+)
Value SENSOR ((\s|\S)+?)

Start
  ^Temperature
  ^-+
  ^Module\s+Sensor\s+MajorThresh\s+
  ^(\s+\(Celsius\))+
  ^${MODULE}\s+${SENSOR}\s+${MAJOR_THRESHOLD}\s+${MINOR_THRESHOLD}\s+${CURRENT_VALUE}\s+${STATUS} -> Record
  ^\s+$$
  ^$$
  ^.* -> Error "LINE NOT FOUND"
`