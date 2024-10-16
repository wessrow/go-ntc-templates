package cisco_nxos

type ShowVrfInterface struct {
	Interface string `json:"INTERFACE"`
	Name      string `json:"NAME"`
	Id        string `json:"ID"`
	Origin    string `json:"ORIGIN"`
}

var ShowVrfInterface_Template string = `Value INTERFACE (\S+)
Value NAME (\S+)
Value ID (\S+)
Value ORIGIN (\S+)

Start
  ^Interface\s+VRF-Name\s+VRF-ID\s+Site-of-Origin -> Start_record

Start_record
  ^${INTERFACE}\s+${NAME}\s+${ID}\s+${ORIGIN} -> Record
  ^\s*$$
  ^. -> Error
`
