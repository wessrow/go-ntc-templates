package huawei_smartax

type DisplayOntInfo012 struct {
	Config_state  string `json:"CONFIG_STATE"`
	Match_state   string `json:"MATCH_STATE"`
	Protect_side  string `json:"PROTECT_SIDE"`
	Fsp           string `json:"FSP"`
	Ont_id        string `json:"ONT_ID"`
	Serial_number string `json:"SERIAL_NUMBER"`
	Control_flag  string `json:"CONTROL_FLAG"`
	Run_state     string `json:"RUN_STATE"`
}

var DisplayOntInfo012_Template string = `Value Required FSP (\w+\/\s\w+\/\w+)
Value Key ONT_ID (\d+)
Value SERIAL_NUMBER (\w+)
Value CONTROL_FLAG (\w+)
Value RUN_STATE (\w*)
Value CONFIG_STATE (\w*)
Value MATCH_STATE (\w*)
Value PROTECT_SIDE (no|yes)

Start
  ^\s+-
  ^\s+F\/S\/P\s+ONT\s+SN\s+Control\s+Run\s+Config\s+Match\s+Protect
  ^\s+ID\s+flag\s+state\s+state\s+state\s+side -> SNs

SNs
  ^\s*${FSP}\s*${ONT_ID}\s*${SERIAL_NUMBER}\s*${CONTROL_FLAG}\s*(${RUN_STATE}|-)\s*(${CONFIG_STATE}|-)\s*(${MATCH_STATE}|-)\s*${PROTECT_SIDE}\s* -> Record
  ^\s*-
  ^\s*F\/S\/P\s+(ONT|ONT-ID)\s+Description -> EOF
  ^. -> Error`
