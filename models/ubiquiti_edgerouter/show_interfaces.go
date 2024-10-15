package ubiquiti_edgerouter

type ShowInterfaces struct {
	Link_status string   `json:"LINK_STATUS"`
	Description string   `json:"DESCRIPTION"`
	Interface   string   `json:"INTERFACE"`
	Ip_address  []string `json:"IP_ADDRESS"`
	State       string   `json:"STATE"`
}

var ShowInterfaces_Template string = `Value Required INTERFACE (\w+(\.\d+)?)
Value List IP_ADDRESS (\S+)
Value STATE (\w)
Value LINK_STATUS (\w)
Value DESCRIPTION ([\w\d\-\s]+)

Start
  ^[-\s]+ -> Interfaces

Interfaces
  ^(\w+) -> Continue.Record
  ^${INTERFACE}\s+${IP_ADDRESS}\s+(${STATE}/${LINK_STATUS})(\s+${DESCRIPTION})?
  ^\s+${IP_ADDRESS}
  ^. -> Error
`
