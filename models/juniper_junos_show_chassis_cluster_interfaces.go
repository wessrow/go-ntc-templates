package models

type JuniperJunosShowChassisClusterInterfaces struct {
	Control_link_status	string	`json:"CONTROL_LINK_STATUS"`
	Fabric_link_status	string	`json:"FABRIC_LINK_STATUS"`
	Link_type	string	`json:"LINK_TYPE"`
	Index	string	`json:"INDEX"`
	Interface	string	`json:"INTERFACE"`
	Status	string	`json:"STATUS"`
	Security	string	`json:"SECURITY"`
	Child_interface	string	`json:"CHILD_INTERFACE"`
	Redundancy_group	string	`json:"REDUNDANCY_GROUP"`
	Weight	string	`json:"WEIGHT"`
}

var JuniperJunosShowChassisClusterInterfaces_Template = `Value Filldown CONTROL_LINK_STATUS (\S+)
Value Filldown FABRIC_LINK_STATUS (\S+)
Value Filldown LINK_TYPE (\S+)
Value INDEX (\d+)
Value INTERFACE (\S+)
Value STATUS ((\S+\s*\/\s*\S+|\S+))
Value SECURITY (\S+)
Value CHILD_INTERFACE (\S+)
Value REDUNDANCY_GROUP (\d+|[Nn]ot\s+[Cc]onfigured)
Value WEIGHT (\d+)


Start
  ^[Cc]ontrol\s+[Ll]ink\s+[Ss]tatus:\s+${CONTROL_LINK_STATUS}\s*$$ -> Control
  ^[Ff]abric\s+[Ll]ink\s+[Ss]tatus:\s+${FABRIC_LINK_STATUS}\s*$$ -> Fabric
  ^${LINK_TYPE}\s+Information:\s*$$ -> Continue
  ^[Rr]edundant-ethernet -> RedEth
  ^[Rr]edundant-pseudo -> RedPse
  ^[Ii]nterface\s+[Mm]onitoring -> IntMon
  ^\s*$$
  ^. -> Error

Control
  ^${LINK_TYPE}\s+interfaces:\s*$$
  ^\s*Index\s+Interface\s+Monitored-Status\s+Security\s*$$
  ^\s*Index\s+Interface\s+Monitored-Status\s+Internal-SA\s*$$
  ^\s*Index\s+Interface\s+Status\s+Internal-SA\s*$$
  ^\s*${INDEX}\s+${INTERFACE}\s+${STATUS}\s+${SECURITY}\s*$$ -> Record
  ^\S+ -> Continue.Clearall
  ^[Ff]abric\s+[Ll]ink\s+[Ss]tatus:\s+${FABRIC_LINK_STATUS}\s*$$ -> Fabric
  ^${LINK_TYPE}\s+Information:\s*$$ -> Continue
  ^[Rr]edundant-ethernet -> RedEth
  ^[Rr]edundant-pseudo -> RedPse
  ^Interface\s+[Mm]onitoring -> IntMon
  ^\s*$$
  ^. -> Error

Fabric
  ^${LINK_TYPE}\s+interfaces:\s*$$
  ^\s*Name\s+Child-interface\s+Status\s+Security\s*$$
  ^\s*Name\s+Child-interface\s+Status\s*$$
  ^\s*\(Physical\/Monitored\)\s*$$
  ^\s*${INTERFACE}\s+${CHILD_INTERFACE}\s+${STATUS}\s+${SECURITY}\s*$$ -> Record
  ^\s*${INTERFACE}\s+${CHILD_INTERFACE}\s+${STATUS}\s*$$ -> Record
  ^\s*\S+\s*$$
  ^\s+\w+\d+\s*$$
  ^\S+ -> Continue.Clearall
  ^${LINK_TYPE}\s+Information:\s*$$ -> Continue
  ^[Rr]edundant-ethernet -> RedEth
  ^[Rr]edundant-pseudo -> RedPse
  ^Interface\s+[Mm]onitoring -> IntMon
  ^\s*$$
  ^Control\s+[Ll]ink\s+[Ss]tatus:\s+${CONTROL_LINK_STATUS}\s*$$ - Control
  ^. -> Error

RedEth
  ^\s*Name\s+Status\s+Redundancy-group\s*$$
  ^\s*${INTERFACE}\s+${STATUS}\s+${REDUNDANCY_GROUP}\s*$$ -> Record
  ^\S+ -> Continue.Clearall
  ^${LINK_TYPE}\s+Information -> Continue
  ^[Rr]edundant-pseudo -> RedPse
  ^Interface\s+[Mm]onitoring -> IntMon
  ^\s*$$
  ^Control\s+[Ll]ink\s+[Ss]tatus:\s+${CONTROL_LINK_STATUS}\s*$$ - Control
  ^[Ff]abric\s+[Ll]ink\s+[Ss]tatus:\s+${FABRIC_LINK_STATUS}\s*$$ -> Fabric
  ^. -> Error

RedPse
  ^\s*Name\s+Status\s+Redundancy-group\s*$$
  ^\s*${INTERFACE}\s+${STATUS}\s+${REDUNDANCY_GROUP}\s*$$ -> Record
  ^\S+ -> Continue.Clearall
  ^${LINK_TYPE}\s+Information -> Continue
  ^[Ii]nterface\s+[Mm]onitoring -> IntMon
  ^\s*$$
  ^Control\s+[Ll]ink\s+[Ss]tatus:\s+${CONTROL_LINK_STATUS}\s*$$ - Control
  ^[Ff]abric\s+[Ll]ink\s+[Ss]tatus:\s+${FABRIC_LINK_STATUS}\s*$$ -> Fabric
  ^[Rr]edundant-ethernet -> RedEth
  ^. -> Error

IntMon
  ^\s*Interface\s+Weight\s+Status\s+Redundancy-group\s*$$
  ^\s*${INTERFACE}\s+${WEIGHT}\s+${STATUS}(?:\s+${REDUNDANCY_GROUP}|)\s*$$ -> Record
  ^\s*$$
  ^\S+ -> Continue.Clearall
  ^${LINK_TYPE}\s+Information:\s*$$ -> Continue
  ^Control\s+[Ll]ink\s+[Ss]tatus:\s+${CONTROL_LINK_STATUS}\s*$$ - Control
  ^[Ff]abric\s+[Ll]ink\s+[Ss]tatus:\s+${FABRIC_LINK_STATUS}\s*$$ -> Fabric
  ^[Rr]edundant-ethernet -> RedEth
  ^[Rr]edundant-pseudo -> RedPse
  ^. -> Error

`