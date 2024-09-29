package cisco_xr 

type AdminShowPlatform struct {
	Node	string	`json:"NODE"`
	Type	string	`json:"TYPE"`
	State	string	`json:"STATE"`
	Config_state	string	`json:"CONFIG_STATE"`
}

var AdminShowPlatform_Template = `Value NODE (\S+)
Value TYPE (\S+)
Value STATE ((\S+(\s\S+)*))
Value CONFIG_STATE (\S+)

Start
  ^(\S+\s+){2}PLIM(\s+\S+){2,3}\s*$$ -> XRvChassis
  ^\-+ -> Normal
  ^Location
  ^Node
  ^(Mon|Tue|Wed|Thu|Fri|Sat|Sun)\s\S{3}\s+\d
  ^. -> Error

Normal
  # show platform
  ^${NODE}\s+${TYPE}\s+${STATE}\s+(${CONFIG_STATE})? -> Record
  # admin show platform
  ^${NODE}\s+${TYPE}\s+\S+\s+${STATE}\s+(${CONFIG_STATE})? -> Record
  ^. -> Error

XRvChassis
  ^${NODE}\s+${TYPE}\s+\S+\s+${STATE}\s+(${CONFIG_STATE})? -> Record
  ^\-+
  ^. -> Error
`