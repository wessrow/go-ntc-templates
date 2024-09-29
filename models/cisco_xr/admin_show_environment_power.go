package cisco_xr 

type AdminShowEnvironmentPower struct {
	Power_group	string	`json:"POWER_GROUP"`
	Power_module	string	`json:"POWER_MODULE"`
	Supply_type	string	`json:"SUPPLY_TYPE"`
	Status	string	`json:"STATUS"`
}

var AdminShowEnvironmentPower_Template = `Value Filldown POWER_GROUP (\d+)
Value Required POWER_MODULE (\d+\/PM\d+)
Value Required SUPPLY_TYPE (\S+)
Value Required STATUS (OK|Failed)

Start
  ^.*\=+ -> Parse

Parse
  ^Power\sGroup+\s+${POWER_GROUP} -> Record
  ^\s+${POWER_MODULE}\s+${SUPPLY_TYPE}+\s+\S+\s+\S+\s+\S+\s+\S+\s+\S+\s+${STATUS} -> Record
`