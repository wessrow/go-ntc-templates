package eltex

type ShowSystem struct {
	Sensor_temp                   []string `json:"SENSOR_TEMP"`
	Unit_redundant_power_supply   []string `json:"UNIT_REDUNDANT_POWER_SUPPLY"`
	Contact                       string   `json:"CONTACT"`
	Fan_id                        []string `json:"FAN_ID"`
	Unit_temp_status              []string `json:"UNIT_TEMP_STATUS"`
	Reset_btn_status              string   `json:"RESET_BTN_STATUS"`
	Fans_status                   []string `json:"FANS_STATUS"`
	Location                      []string `json:"LOCATION"`
	Address                       string   `json:"ADDRESS"`
	Oid                           string   `json:"OID"`
	Main_power_supply_status      string   `json:"MAIN_POWER_SUPPLY_STATUS"`
	Redundant_power_supply_status string   `json:"REDUNDANT_POWER_SUPPLY_STATUS"`
	Sensor_status                 []string `json:"SENSOR_STATUS"`
	Description                   string   `json:"DESCRIPTION"`
	Hostname                      string   `json:"HOSTNAME"`
	Unit_main_power_supply        []string `json:"UNIT_MAIN_POWER_SUPPLY"`
	Unit_fans_status              []string `json:"UNIT_FANS_STATUS"`
	Unit_descriptions             []string `json:"UNIT_DESCRIPTIONS"`
	Unit_temp                     []string `json:"UNIT_TEMP"`
	Uptime                        string   `json:"UPTIME"`
	Sensor_id                     []string `json:"SENSOR_ID"`
}

var ShowSystem_Template string = `Value DESCRIPTION (.*)
Value UPTIME (.*)
Value CONTACT (.*)
Value HOSTNAME (.*)
Value List LOCATION (.*)
Value ADDRESS (.*)
Value OID (.*)
Value MAIN_POWER_SUPPLY_STATUS (.*)
Value REDUNDANT_POWER_SUPPLY_STATUS (.*)
Value RESET_BTN_STATUS (.*)
Value List FAN_ID (\d+)
Value List FANS_STATUS (.*)
Value List SENSOR_ID (\d+)
Value List SENSOR_TEMP (\d+)
Value List SENSOR_STATUS (.*)
Value List UNIT_DESCRIPTIONS (\S+(?:\s+\S+)*)
Value List UNIT_MAIN_POWER_SUPPLY (OK|NOT\s+PRESENT|\S+)
Value List UNIT_REDUNDANT_POWER_SUPPLY (OK|NOT\s+PRESENT|\S+)
Value List UNIT_FANS_STATUS (.*)
Value List UNIT_TEMP (\d+)
Value List UNIT_TEMP_STATUS (.*)

Start
  ^\s*System\s+Description:\s*${DESCRIPTION}\s*$$
  ^\s*System\s+Up\s+Time\s+\(days,hour:min:sec\):\s*${UPTIME}\s*$$
  ^\s*System\s+Contact:\s*${CONTACT}\s*$$
  ^\s*System\s+Name:\s*${HOSTNAME}\s*$$
  ^\s*System\s+Location:\s*${LOCATION}\s*$$ -> SystemLocation
  ^\s*System\s+MAC\s+Address:\s*${ADDRESS}\s*$$
  ^\s*System\s+Object\s+ID:\s*${OID}\s*$$
  ^\s*Reset-Button:\s*${RESET_BTN_STATUS}\s*$$
  ^\s*Main\s+Power\s+Supply\s+Status(?:\s+\[AC\])?:\s*${MAIN_POWER_SUPPLY_STATUS}\s*$$
  ^\s*Redundant\s+Power\s+Supply\s+Status(?:\s+\[[AD]C\])?:\s*${REDUNDANT_POWER_SUPPLY_STATUS}\s*$$
  ^\s*Fan(?:\s+Module)?\s+${FAN_ID}\s+Status:\s*${FANS_STATUS}\s*$$
  ^\s*Sensor\s+Index\s+Temperature\s+\(Celsius\)\s+Status\s*$$ -> Sensors
  ^\s*Unit\s+Device\s*$$ -> UnitsDescription
  ^\s*$$
  ^. -> Error

SystemLocation
  ^\s*System\s+MAC\s+Address:\s*${ADDRESS}\s*$$ -> Start
  ^\s*${LOCATION}\s*$$
  ^\s*$$
  ^. -> Error

Sensors
  ^(\s*-+)*\s*$$
  ^\s*${SENSOR_ID}\s+${SENSOR_TEMP}\s+${SENSOR_STATUS}\s*$$
  ^\s*$$
  ^. -> Error

UnitsDescription
  ^(\s*-+)*\s*$$
  ^\s*\d+\s+${UNIT_DESCRIPTIONS}\s*$$
  ^\s*Unit\s+Fans\s+Status\s*$$ -> UnitsFansStatuses
  ^\s*Unit\s+Main\s+Power\s+Supply\s+Redundant\s+Power\s+Supply -> UnitsPowerSupplies
  ^\s*${UNIT_DESCRIPTIONS}\s*$$
  ^\s*$$
  ^. -> Error

UnitsPowerSupplies
  ^(\s*-+)*\s*$$
  ^\s*\d+\s+${UNIT_MAIN_POWER_SUPPLY}\s+${UNIT_REDUNDANT_POWER_SUPPLY}\s*$$
  ^\s*Unit\s+Fans\s+Status\s*$$ -> UnitsFansStatuses
  ^\s*$$
  ^. -> Error

UnitsFansStatuses
  ^(\s*-+)*\s*$$
  ^\s*\d+\s+${UNIT_FANS_STATUS}\s*$$
  ^\s*Unit\s+Temperature\s+\(Celsius\)\s+Status\s*$$ -> UnitsTempStatuses
  ^\s*$$
  ^. -> Error

UnitsTempStatuses
  ^(\s*-+)*\s*$$
  ^\s*\d+\s+${UNIT_TEMP}\s+${UNIT_TEMP_STATUS}\s*$$
  ^\s*$$
  ^. -> Error
`
