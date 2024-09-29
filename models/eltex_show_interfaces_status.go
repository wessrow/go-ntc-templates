package models

type EltexShowInterfacesStatus struct {
	Interface	string	`json:"INTERFACE"`
	Type	string	`json:"TYPE"`
	Duplex	string	`json:"DUPLEX"`
	Speed	string	`json:"SPEED"`
	Bw	string	`json:"BW"`
	Neg	string	`json:"NEG"`
	Flow_ctrl	string	`json:"FLOW_CTRL"`
	Link_state	string	`json:"LINK_STATE"`
	Up_time	string	`json:"UP_TIME"`
	Back_pressure	string	`json:"BACK_PRESSURE"`
	Mdix_mode	string	`json:"MDIX_MODE"`
	Interface_mode	string	`json:"INTERFACE_MODE"`
}

var EltexShowInterfacesStatus_Template = `Value INTERFACE (\S+)
Value TYPE (\S+)
Value DUPLEX (\S+)
Value SPEED (\d+|--)
Value BW (\S+)
Value NEG (Enabled|Disabled|--)
Value FLOW_CTRL (On|Off|--)
Value LINK_STATE ((Up|Down(?:\s+\(\S+\))?|Not Present))
Value UP_TIME (\d+,\d{2}:\d{2}:\d{2}|--)
Value BACK_PRESSURE (\S+)
Value MDIX_MODE (On|Off|--)
Value INTERFACE_MODE (.*?)

Start
  ^\s*(?:Flow\s+)?Link(?:(?:\s+Up(\s+)?[Tt]ime)?\s+Back\s+Mdix)?\s*$$
  ^\s*Port\s+Type\s+Speed\s+control\s+State\s+Port\s+Mode\s*$$ -> Port6
  ^\s*Port\s+Type\s+Duplex\s+Speed\s+Neg\s+ctrl\s+State\s+\(d,h:m:s\)\s+Pressure\s+Mode\s+Port\s+Mode.*$$ -> Port11
  ^\s*Port\s+Type\s+Duplex\s+Speed\s+Neg\s+ctrl\s+State\s+Pressure\s+Mode\s+Port\s+Mode\s*$$ -> Port10
  ^\s*Ch\s+BW\s+control\s+State\s+Port\s+Mode\s*$$ -> Ch5
  ^\s*Ch\s+Duplex\s+BW\s+Neg\s+control\s+State\s+Port\s+Mode\s*$$ -> Ch7
  ^\s*Ch\s+Type\s+Duplex\s+Speed\s+Neg\s+control\s+State\s*$$ -> Ch7_v2
  ^\s*Ch\s+Type\s+Duplex\s+Speed\s+Neg\s+control\s+State\s+Port\s+Mode.*$$ -> Ch8
  ^\s*Oob\s+ -> Oob
  ^\s*nc\s+\(not\s+connected\)\s*:\s+The\s+interface\s+is\s+not\s+connected\.\s*$$
  ^\s*err\s+\(error-disabled\)\s*:\s+The\s+interface\s+was\s+suspended\s+by\s+the\s+system\.\s*$$
  ^\s*adm\s+\(admin\.shutdown\)\s*:\s+The\s+interface\s+was\s+suspended\s+by\s+administrator\.\s*$$
  ^\s*$$
  ^. -> Error

Port6
  ^(?:\s*-+)+\s*$$
  ^${INTERFACE}\s+${TYPE}\s+${SPEED}\s+${FLOW_CTRL}\s+${LINK_STATE}\s+${INTERFACE_MODE}\s*$$ -> Record
  ^\s*$$ -> Start
  ^. -> Error

Port11
  ^(?:\s*-+)+\s*$$
  ^\s*${INTERFACE}\s+${TYPE}\s+${DUPLEX}\s+${SPEED}\s+${NEG}\s+${FLOW_CTRL}\s+${LINK_STATE}\s+${UP_TIME}\s*${BACK_PRESSURE}\s+${MDIX_MODE}\s+${INTERFACE_MODE}\s*$$ -> Record
  ^\s*$$ -> Start
  ^. -> Error

Port10
  ^(?:\s*-+)+\s*$$
  ^${INTERFACE}\s+${TYPE}\s+${DUPLEX}\s+${SPEED}\s+${NEG}\s+${FLOW_CTRL}\s+${LINK_STATE}\s+${BACK_PRESSURE}\s+${MDIX_MODE}\s+${INTERFACE_MODE}\s*$$ -> Record
  ^\s*$$ -> Start
  ^. -> Error

Ch5
  ^(?:\s*-+)+\s*$$
  ^\s*${INTERFACE}\s+${BW}\s+${FLOW_CTRL}\s+${LINK_STATE}\s+${INTERFACE_MODE}\s*$$ -> Record
  ^\s*$$ -> Start
  ^. -> Error

Ch7
  ^(?:\s*-+)+\s*$$
  ^\s*${INTERFACE}\s+${DUPLEX}\s+${BW}\s+${NEG}\s+${FLOW_CTRL}\s+${LINK_STATE}\s+${INTERFACE_MODE}\s*$$ -> Record
  ^\s*$$ -> Start
  ^. -> Error

Ch7_v2
  ^(?:\s*-+)+\s*$$
  ^\s*${INTERFACE}\s+${TYPE}\s+${DUPLEX}\s+${SPEED}\s+${NEG}\s+${FLOW_CTRL}\s+${LINK_STATE}\s*$$ -> Record
  ^\s*$$ -> Start
  ^. -> Error

Ch8
  ^(?:\s*-+)+\s*$$
  ^\s*${INTERFACE}\s+${TYPE}\s+${DUPLEX}\s+${SPEED}\s+${NEG}\s+${FLOW_CTRL}\s+${LINK_STATE}\s+${INTERFACE_MODE}\s*$$ -> Record
  ^\s*$$ -> Start
  ^. -> Error

Oob
  ^(?:\s*-+)+\s*$$
  ^\s*${INTERFACE}\s+${TYPE}\s+${DUPLEX}\s+${SPEED}\s+${NEG}\s+${LINK_STATE}\s*$$ -> Record
  ^\s*$$ -> Start
  ^. -> Error

`