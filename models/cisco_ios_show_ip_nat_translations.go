package models

type CiscoIosShowIpNatTranslations struct {
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