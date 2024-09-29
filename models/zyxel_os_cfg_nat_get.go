package models

type ZyxelOsCfgNatGet struct {
	Index	string	`json:"INDEX"`
	Name	string	`json:"NAME"`
	Enabled	string	`json:"ENABLED"`
	Originating_ip	string	`json:"ORIGINATING_IP"`
	Server_ip	string	`json:"SERVER_IP"`
	Wan_interface	string	`json:"WAN_INTERFACE"`
	Start_port	string	`json:"START_PORT"`
	End_port	string	`json:"END_PORT"`
	Transl_start_port	string	`json:"TRANSL_START_PORT"`
	Transl_end_port	string	`json:"TRANSL_END_PORT"`
	Protocol	string	`json:"PROTOCOL"`
}

var ZyxelOsCfgNatGet_Template = `Value INDEX (\d+)
Value NAME (.+?)
Value ENABLED (0|1)
Value ORIGINATING_IP (\d+\.\d+\.\d+\.\d+|N\/A)
Value SERVER_IP (.+?)
Value WAN_INTERFACE (.+?)
Value START_PORT (\d+)
Value END_PORT (\d+)
Value TRANSL_START_PORT (\d+)
Value TRANSL_END_PORT (\d+)
Value PROTOCOL (.+?)

Start
  ^Index\s+Description\s+Enable\s+Originating\sIP\s+Server\sIP\sAddress\s+WAN\sInterface\s+Start\sPort\s+End\sPort\s+Translation\sStart\sPort\s+Translation\sEnd\sPort\s+Protocol\s*$$ -> NATTable
  ^\s*$$
  ^. -> Error

NATTable
  ^${INDEX}\s+${NAME}\s+${ENABLED}\s+${ORIGINATING_IP}\s+${SERVER_IP}\s+${WAN_INTERFACE}\s+${START_PORT}\s+${END_PORT}\s+${TRANSL_START_PORT}\s+${TRANSL_END_PORT}\s+${PROTOCOL}\s*$$ -> Record
  ^Command Successful.\s*$$
  ^\s*$$
  ^. -> Error

`