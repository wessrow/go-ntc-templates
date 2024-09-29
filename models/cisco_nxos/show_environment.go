package cisco_nxos 

type ShowEnvironment struct {
	Power_supply	string	`json:"POWER_SUPPLY"`
	Power_supply_model	string	`json:"POWER_SUPPLY_MODEL"`
	Power_supply_output	string	`json:"POWER_SUPPLY_OUTPUT"`
	Power_supply_input	string	`json:"POWER_SUPPLY_INPUT"`
	Power_supply_capacity	string	`json:"POWER_SUPPLY_CAPACITY"`
	Power_supply_status	string	`json:"POWER_SUPPLY_STATUS"`
	Fan	string	`json:"FAN"`
	Fan_status	string	`json:"FAN_STATUS"`
	Temperature_module	string	`json:"TEMPERATURE_MODULE"`
	Temperature_sensor	string	`json:"TEMPERATURE_SENSOR"`
	Temperature_major_thresh	string	`json:"TEMPERATURE_MAJOR_THRESH"`
	Temperature_minor_thresh	string	`json:"TEMPERATURE_MINOR_THRESH"`
	Temperature_current	string	`json:"TEMPERATURE_CURRENT"`
	Temperature_status	string	`json:"TEMPERATURE_STATUS"`
}

var ShowEnvironment_Template = `Value POWER_SUPPLY (\d+)
Value POWER_SUPPLY_MODEL (\S+)
Value POWER_SUPPLY_OUTPUT (\d+)
Value POWER_SUPPLY_INPUT (\d+)
Value POWER_SUPPLY_CAPACITY (\d+)
Value POWER_SUPPLY_STATUS (\w+)
Value FAN (Fan\S+)
Value FAN_STATUS (\S+)
Value TEMPERATURE_MODULE (\d+)
Value TEMPERATURE_SENSOR ([^\s\(\)]+)
Value TEMPERATURE_MAJOR_THRESH (\d+)
Value TEMPERATURE_MINOR_THRESH (\d+)
Value TEMPERATURE_CURRENT (\d+)
Value TEMPERATURE_STATUS (\S+)

Start
  # Note: 2020-12 this template is broken, mashing 3 tables into one.
  ^Power Supply:\s*$$ -> Power
  ^Fan:\s*$$ -> Fan
  ^Temperature:\s*$$ -> Temperature
  #^Fan\s+Model\s+Hw\s+(Direction\s+)?Status -> Fan

Power
  #Capture Power with only Out
  ^Power\s+Actual\s+Total
  ^${POWER_SUPPLY}\s+${POWER_SUPPLY_MODEL}\s+${POWER_SUPPLY_OUTPUT}\s+\w+\s+${POWER_SUPPLY_CAPACITY}\s+\w+\s+${POWER_SUPPLY_STATUS}\s*$$ -> Record
  #
  # Capture Power with Out and In
  ^Power\s+Actual\s+Actual\s+Total
  ^${POWER_SUPPLY}\s+${POWER_SUPPLY_MODEL}\s+${POWER_SUPPLY_OUTPUT}\s+\w+\s+${POWER_SUPPLY_INPUT}\s+\w+\s+${POWER_SUPPLY_CAPACITY}\s+\w+\s+${POWER_SUPPLY_STATUS}\s*$$ -> Record
  #
  # Done with Power section back to Start
  ^Module\s+Model\s+Draw\s+Allocated\s+Status\s* -> Start
  ^Power\s+Usage\s+Summary: -> Start
  #
  # Skip junk in Power section
  ^Voltage:\s+\d+\s+Volts
  ^\s*$$ -> Start
  ^Supply\s+Model\s+Output\s+Capacity\s+Status
  ^Supply\s+Model\s+Output\s+Input\s+Capacity\s+Status
  ^\s+\(Watts\s\)\s+
  ^-+\s+-+
  ^. -> Error

Fan
  ^\s*Fan\s+Model\s+Hw\s+Status
  ^${FAN}\s+\S+\s+\S+\s+${FAN_STATUS}\s*$$ -> Record
  ^\s*Fan\s+Model\s+Hw\s+Direction\s+Status
  ^${FAN}\s+\S+\s+\S+\s+\S+\s+${FAN_STATUS}\s*$$ -> Record
  #
  # Blank line back to Start
  ^\s*$$ -> Start
  #
  # Fan Ignore
  ^Fan\s+Zone\s+Speed\s*:
  ^Fan\sAir\sFilter\s:
  ^-+\s*$$
  ^. -> Error

Temperature
  ^Module\s+Sensor\s+MajorThresh\s+MinorThres\s+CurTemp\s+Status
  ^${TEMPERATURE_MODULE}\s+${TEMPERATURE_SENSOR}(\s*\(\S+\))?\s+${TEMPERATURE_MAJOR_THRESH}\s+${TEMPERATURE_MINOR_THRESH}\s+${TEMPERATURE_CURRENT}\s+${TEMPERATURE_STATUS}\s* -> Record
  ^\s*\(Celsius\)\s+\(Celsius\)
  ^-+\s*$$
  ^\s*$$
  ^. -> Error
`