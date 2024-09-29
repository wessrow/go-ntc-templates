package models

type HuaweiVrpDisplayMplsTeTunnel struct {
	Source	string	`json:"SOURCE"`
	Destination	string	`json:"DESTINATION"`
	Id	string	`json:"ID"`
	In_label	string	`json:"IN_LABEL"`
	Out_label	string	`json:"OUT_LABEL"`
	Role	string	`json:"ROLE"`
	Name	string	`json:"NAME"`
}

var HuaweiVrpDisplayMplsTeTunnel_Template = `Value SOURCE (\d+\.\d+\.\d+\.\d+)
Value DESTINATION (\d+\.\d+\.\d+\.\d+)
Value ID (\d+)
Value IN_LABEL (\S+)
Value OUT_LABEL (\S+)
Value ROLE (\S+)
Value NAME (\S+)

Start
  ^-+
  ^Ingress\s+LsrId\s+Destination\s+LSPID\s+In/OutLabel\s+R\s+Tunnel-name\s*$$
  ^-+
  ^${SOURCE}\s+${DESTINATION}\s+${ID}?\s+${IN_LABEL}/${OUT_LABEL}\s+${ROLE}\s+${NAME}\s*$$ -> Record
  ^-+
  ^\s*\*\s+means
  ^\s*\S:\s+\S+,\s+\S+:\s+\S+,
  ^\s*$$
  ^. -> Error

`