package models

type MikrotikRouterosInterfaceEthernetPoePrintWithoutPaging struct {
	Index	string	`json:"INDEX"`
	Interface	string	`json:"INTERFACE"`
	Poe_out	string	`json:"POE_OUT"`
	Poe_voltage	string	`json:"POE_VOLTAGE"`
	Poe_priority	string	`json:"POE_PRIORITY"`
	Power_cycle_ping_enabled	string	`json:"POWER_CYCLE_PING_ENABLED"`
	Power_cycle_interval	string	`json:"POWER_CYCLE_INTERVAL"`
}

var MikrotikRouterosInterfaceEthernetPoePrintWithoutPaging_Template = `Value INDEX (\d+)
Value INTERFACE (\S+)
Value POE_OUT (\S+)
Value POE_VOLTAGE (\S+)
Value POE_PRIORITY (\S+)
Value POWER_CYCLE_PING_ENABLED (\S+)
Value POWER_CYCLE_INTERVAL (\S+)

Start
  ^\s*\#\s+NAME\s+POE\-OUT\s+POE\-PRIORITY\s+POWER\-CYCLE\-PING\-ENABLED\s+POWER\-CYCLE\-INTERVAL\s*$$
  ^\s*${INDEX}\s+${INTERFACE}\s+(${POE_OUT}\s+)?(${POE_PRIORITY}\s+)?(${POWER_CYCLE_PING_ENABLED}\s+)?(${POWER_CYCLE_INTERVAL})?\s*$$ -> Record
  ^\s*\#\s+NAME\s+POE\-OUT\s+POE\-VOLTAGE\s+POE\-PRIORITY\s+POWER\-CYCLE\-PING\-ENABLED\s+POWER\-CYCLE\-INTERVAL\s*$$
  ^\s*${INDEX}\s+${INTERFACE}\s+(${POE_OUT}\s+)?(${POE_VOLTAGE}\s+)?(${POE_PRIORITY}\s+)?(${POWER_CYCLE_PING_ENABLED}\s+)?(${POWER_CYCLE_INTERVAL})?\s*$$ -> Record
  ^\s*bad\s+command\s+name\s+po(e)?\s+\(line\s+\d+\s+column\s+\d+\)\s*$$
  ^\s*$$
  ^. -> Error
`