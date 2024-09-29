package cisco_ios 

type ShowIpNatTranslations struct {
	Protocol	string	`json:"PROTOCOL"`
	Inside_global_ip	string	`json:"INSIDE_GLOBAL_IP"`
	Inside_global_port	string	`json:"INSIDE_GLOBAL_PORT"`
	Inside_local_ip	string	`json:"INSIDE_LOCAL_IP"`
	Inside_local_port	string	`json:"INSIDE_LOCAL_PORT"`
	Outside_local_ip	string	`json:"OUTSIDE_LOCAL_IP"`
	Outside_local_port	string	`json:"OUTSIDE_LOCAL_PORT"`
	Outside_global_ip	string	`json:"OUTSIDE_GLOBAL_IP"`
	Outside_global_port	string	`json:"OUTSIDE_GLOBAL_PORT"`
}

var ShowIpNatTranslations_Template = `Value PROTOCOL (tcp|udp|icmp|---)
Value INSIDE_GLOBAL_IP (\d+\.\d+\.\d+\.\d+|---)
Value INSIDE_GLOBAL_PORT (\S+)
Value INSIDE_LOCAL_IP (\d+\.\d+\.\d+\.\d+|---)
Value INSIDE_LOCAL_PORT (\S+)
Value OUTSIDE_LOCAL_IP (\d+\.\d+\.\d+\.\d+|---)
Value OUTSIDE_LOCAL_PORT (\S+)
Value OUTSIDE_GLOBAL_IP (\d+\.\d+\.\d+\.\d+|---)
Value OUTSIDE_GLOBAL_PORT (\S+)

Start
  ^Pro\s+Inside\sglobal\s+Inside\slocal\s+Outside\slocal\s+Outside\sglobal
  ^${PROTOCOL}\s+${INSIDE_GLOBAL_IP}(:${INSIDE_GLOBAL_PORT})?\s+${INSIDE_LOCAL_IP}(:${INSIDE_LOCAL_PORT})?\s+${OUTSIDE_LOCAL_IP}(:${OUTSIDE_GLOBAL_PORT})?\s+${OUTSIDE_GLOBAL_IP}(:${OUTSIDE_LOCAL_PORT})? -> Record
`