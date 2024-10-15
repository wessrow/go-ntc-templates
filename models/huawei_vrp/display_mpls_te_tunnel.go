package huawei_vrp

type DisplayMplsTeTunnel struct {
	In_label    string `json:"IN_LABEL"`
	Out_label   string `json:"OUT_LABEL"`
	Role        string `json:"ROLE"`
	Name        string `json:"NAME"`
	Source      string `json:"SOURCE"`
	Destination string `json:"DESTINATION"`
	Id          string `json:"ID"`
}

var DisplayMplsTeTunnel_Template string = `Value SOURCE (\d+\.\d+\.\d+\.\d+)
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
