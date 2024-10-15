package cisco_nxos

type ShowVrf struct {
	Id     string `json:"ID"`
	State  string `json:"STATE"`
	Reason string `json:"REASON"`
	Name   string `json:"NAME"`
}

var ShowVrf_Template string = `Value Required NAME (\S+)
Value Required ID (\S+)
Value Required STATE (\S+)
Value Required REASON (\S+)

Start
  ^VRF-Name\s+VRF-ID\s+State\s+Reason -> Start_record

Start_record
  ^${NAME}\s+${ID}\s+${STATE}\s+${REASON} -> Record
  ^\s+$$
  ^$$
  ^.* -> Error "LINE NOT FOUND"
`
