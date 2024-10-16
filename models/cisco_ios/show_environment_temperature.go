package cisco_ios

type ShowEnvironmentTemperature struct {
	Hotspot_yellow_threshold  string `json:"HOTSPOT_YELLOW_THRESHOLD"`
	Inlet_temperature_state   string `json:"INLET_TEMPERATURE_STATE"`
	Inlet_red_threshold       string `json:"INLET_RED_THRESHOLD"`
	Hotspot_temperature_value string `json:"HOTSPOT_TEMPERATURE_VALUE"`
	Hotspot_temperature_state string `json:"HOTSPOT_TEMPERATURE_STATE"`
	Hotspot_red_threshold     string `json:"HOTSPOT_RED_THRESHOLD"`
	Switch                    string `json:"SWITCH"`
	Switch_temperature        string `json:"SWITCH_TEMPERATURE"`
	Inlet_temperature_value   string `json:"INLET_TEMPERATURE_VALUE"`
	Inlet_yellow_threshold    string `json:"INLET_YELLOW_THRESHOLD"`
}

var ShowEnvironmentTemperature_Template string = `Value SWITCH (\d+)
Value SWITCH_TEMPERATURE (\w+)
Value INLET_TEMPERATURE_VALUE (\d+)
Value INLET_TEMPERATURE_STATE (\S+)
Value INLET_YELLOW_THRESHOLD (\d+)
Value INLET_RED_THRESHOLD (\d+)
Value HOTSPOT_TEMPERATURE_VALUE (\d+)
Value HOTSPOT_TEMPERATURE_STATE (\S+)
Value HOTSPOT_YELLOW_THRESHOLD (\d+)
Value HOTSPOT_RED_THRESHOLD (\d+)

Start
  ^Switch\s+\d -> Continue.Record
  ^Switch\s+${SWITCH}:\s+SYSTEM\s+TEMPERATURE\s+is\s+${SWITCH_TEMPERATURE}\s*$$
  ^Inlet\s+Temperature\s+Value:\s+${INLET_TEMPERATURE_VALUE}\s+Degree\s+Celsius\s*$$ -> Inlet
  ^Hotspot\s+Temperature\s+Value:\s+${HOTSPOT_TEMPERATURE_VALUE}\s+Degree\s+Celsius\s*$$ -> Hotspot
  ^\s*$$
  ^. -> Error 
 
Inlet
  ^Temperature\s+State:\s+${INLET_TEMPERATURE_STATE}\s*$$
  ^Yellow\s+Threshold\s+:\s+${INLET_YELLOW_THRESHOLD}\s+Degree\s+Celsius\s*$$
  ^Red\s+Threshold\s+:\s+${INLET_RED_THRESHOLD}\s+Degree\s+Celsius\s*$$
  ^Hotspot\s+Temperature\s+Value:\s+${HOTSPOT_TEMPERATURE_VALUE}\s+Degree\s+Celsius\s*$$ -> Hotspot
  ^Switch\s+\d -> Continue.Record
  ^Switch\s+${SWITCH}:\s+SYSTEM\s+TEMPERATURE\s+is\s+${SWITCH_TEMPERATURE}\s*$$ -> Start
  ^\s*$$
  ^. -> Error 
 
Hotspot
  ^Temperature\s+State:\s+${HOTSPOT_TEMPERATURE_STATE}\s*$$
  ^Yellow\s+Threshold\s+:\s+${HOTSPOT_YELLOW_THRESHOLD}\s+Degree\s+Celsius\s*$$
  ^Red\s+Threshold\s+:\s+${HOTSPOT_RED_THRESHOLD}\s+Degree\s+Celsius\s*$$
  ^Switch\s+\d -> Continue.Record
  ^Switch\s+${SWITCH}:\s+SYSTEM\s+TEMPERATURE\s+is\s+${SWITCH_TEMPERATURE}\s*$$ -> Start
  ^Inlet\s+Temperature\s+Value:\s+${INLET_TEMPERATURE_VALUE}\s+Degree\s+Celsius\s*$$ -> Inlet
  ^\s*$$
  ^. -> Error

`
