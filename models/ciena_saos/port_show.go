package ciena_saos

type PortShow struct {
	Name          string `json:"NAME"`
	Duration      string `json:"DURATION"`
	Xcvr          string `json:"XCVR"`
	Admin_link    string `json:"ADMIN_LINK"`
	Admin_mode    string `json:"ADMIN_MODE"`
	Admin_autoneg string `json:"ADMIN_AUTONEG"`
	Type          string `json:"TYPE"`
	Link          string `json:"LINK"`
	Stp           string `json:"STP"`
	Mode          string `json:"MODE"`
	Autoneg       string `json:"AUTONEG"`
}

var PortShow_Template string = `Value NAME ([0-9A-Za-z\.]+)
Value TYPE (\S*)
Value LINK (Up|Down)
Value DURATION (\w+\s*\wh\s*\w+m\s*\w+)
Value XCVR ((\s*)|(\w+))
Value STP (\w+)
Value MODE ((\w+\/\w+)|(\s*))
Value AUTONEG ((\w+)|(\s*))
Value ADMIN_LINK (\w+)
Value ADMIN_MODE ((\w+\/\w+)|(\s*))
Value ADMIN_AUTONEG ((\w+)|(\s*))


Start
  ^\+-+\s*PORT\s*GLOBAL\s*CONFIGURATION\s*-+\+ 
  ^\+-+
  ^\|\s*P\w+\s*\|\s*V\w+\s*\|
  ^\+-+
  ^\|\s*Rx\s*\w+\s*\s*\w+\s*\w+\s*\w+\s*\w+\s*\|\s*\w+\s*\|
  ^\+-+
  ^\+-+
  ^\|\s*Por\w+\s*Table\s*\|\s*Operational\s*Status\s*\|\s*Admin\s*Config\s*\|
  ^\|-+
  ^\|\s*Por\w+\s*\|\s*Port\s*\s*\|\s*\|\s*Link\s*\State\s*\|\s*\|\s*\|\s*\|Auto\|\s*\|\s*\|Auto\|
  ^\|\s*Name\s*\|\s*Type\s*\|Link\|\s*Duration\s*\|XCVR\|STP\|\s*Mode\s*\|Neg\s*\|Link\|\s*Mode\s*\|Neg\s*\|
  ^\|-+
  ^\|\s*${NAME}\s*\|\s*${TYPE}\s*\|\s*${LINK}\s*\|\s*${DURATION}\|\s*${XCVR}\s*\|${STP}\|\s*${MODE}\s*\|\s*${AUTONEG}\s*\|\s*${ADMIN_LINK}\s*\|\s*${ADMIN_MODE}\s*\|\s*${ADMIN_AUTONEG}\s*\| -> Record
  ^\+-+
  ^\s*$$
  ^. -> Error
  
`
