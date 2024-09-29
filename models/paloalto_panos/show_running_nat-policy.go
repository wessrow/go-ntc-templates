package paloalto_panos 

type ShowRunningNatPolicy struct {
	Name	string	`json:"NAME"`
	Nat_type	string	`json:"NAT_TYPE"`
	From	string	`json:"FROM"`
	Source	string	`json:"SOURCE"`
	To	string	`json:"TO"`
	To_interface	string	`json:"TO_INTERFACE"`
	Destination	string	`json:"DESTINATION"`
	Service	string	`json:"SERVICE"`
	Translate_to	string	`json:"TRANSLATE_TO"`
	Terminal	string	`json:"TERMINAL"`
}

var ShowRunningNatPolicy_Template = `Value Filldown NAME (.*?)
Value Required NAT_TYPE (\S+)
Value FROM (\S+|\[(\s\S+)+\s\])
# any  |  ip (optional mask)  |  [ multiIP ]
Value SOURCE (any|([A-Fa-f0-9:\.]+(\/\d+)?)|\[(\s[A-Fa-f0-9:\.]+(\/\d+)?)+\s\])
Value TO (\S+|\[(\s\S+)+\s\])
Value TO_INTERFACE (\S*)
Value DESTINATION (\S+|\[(\s\S+)+\s\])
Value SERVICE ([\S+\s+]+)
Value TRANSLATE_TO ([\S+\s+]+)
Value TERMINAL (\S+)

Start
  ^${NAME}\s+\{$$
  ^\s+nat-type\s+${NAT_TYPE};$$
  ^\s+from\s+${FROM};$$
  ^\s+source\s+${SOURCE};$$ 
  ^\s+to\s+${TO};$$
  ^.+to-interface\s${TO_INTERFACE}\s?;$$
  ^\s+destination\s+${DESTINATION};$$
  ^\s+service\s+${SERVICE};$$
  ^\s+translate-to\s+"${TRANSLATE_TO}";$$
  ^\s+terminal\s+${TERMINAL};$$
  ^}$$ -> Record
  ^. -> Error
`