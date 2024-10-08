package huawei_smartax 

type DisplaySysmanServiceState struct {
	Network_service	string	`json:"NETWORK_SERVICE"`
	Port	string	`json:"PORT"`
	State	string	`json:"STATE"`
}

var DisplaySysmanServiceState_Template = `Value Key NETWORK_SERVICE (\S+)
Value PORT (\d+)
Value STATE (enable|disable)

Start
  ^\s*--+
  ^\s*Network\s*service\s*Port\s*State
  ^\s*--+
  ^\s*${NETWORK_SERVICE}\s+${PORT}\s+${STATE} -> Record
  ^\s*${NETWORK_SERVICE}\s+-+\s+${STATE} -> Record
  ^\s*${NETWORK_SERVICE}\s+${PORT}+\s-+ -> Record
  ^\s*${NETWORK_SERVICE}\s+-+\s+-+ -> Record
  ^\s*$$ -> EOF
  ^. -> Error`