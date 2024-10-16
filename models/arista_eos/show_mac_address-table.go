package arista_eos

type ShowMacAddressTable struct {
	Moves            string   `json:"MOVES"`
	Last_move        string   `json:"LAST_MOVE"`
	Mac_address      string   `json:"MAC_ADDRESS"`
	Type             string   `json:"TYPE"`
	Vlan_id          string   `json:"VLAN_ID"`
	Destination_port []string `json:"DESTINATION_PORT"`
}

var ShowMacAddressTable_Template string = `Value MAC_ADDRESS (\w+\.\w+\.\w+|\w+\:\w+\:\w+\:\w+\:\w+\:\w+)
Value TYPE (\w+)
Value VLAN_ID (\d+)
Value List DESTINATION_PORT ((\w+\/\w+)|\w+)
Value MOVES (\d+)
Value LAST_MOVE (.+)

Start
  ^\s*Vlan\s+(M|m)ac\sAddress\s+Type\s+Ports\s+Moves\s+Last\sMove\s*$$ -> MAT
  ^\s*Vlan\s+(M|m)ac\sAddress\s+Type\s+Ports\s*$$ -> MMAT
  ^\s*$$
  ^\s*(M|m)ac\sAddress\sTable\s*$$
  ^\s*Multicast\s(M|m)ac\sAddress\sTable\s*$$
  ^-*$$
  # ^(-*\s*)*$$
  ^. -> Error

MAT
  # capture the mac address table
  ^\s*${VLAN_ID}\s+${MAC_ADDRESS}\s+${TYPE}\s+${DESTINATION_PORT}\s*$$ -> Record
  ^\s*${VLAN_ID}\s+${MAC_ADDRESS}\s+${TYPE}\s+${DESTINATION_PORT}\s+${MOVES}\s+${LAST_MOVE}\sago$$ -> Record
  ^\s*Total\s+[M|m]ac -> Start
  ^\s*-+\s+.*-*$$
  ^\s*$$
  ^. -> Error

MMAT
  # capture the multicast mac address table
  # key on the record
  ^\s*\d+\s+(\w+\.\w+\.\w+|\w+\:\w+\:\w+\:\w+\:\w+\:\w+)\s+\w+\s*.*$$ -> Continue.Record
  # match the record with one to five destination ports more.
  ^\s*${VLAN_ID}\s+${MAC_ADDRESS}\s+${TYPE}\s+${DESTINATION_PORT}$$ -> Record
  ^\s*${VLAN_ID}\s+${MAC_ADDRESS}\s+${TYPE}\s+${DESTINATION_PORT}\s -> Continue
  ^\s*\d+\s+(\w+\.\w+\.\w+|\w+\:\w+\:\w+\:\w+\:\w+\:\w+)\s+\w+\s+((\w+\/\w+|\w+)\s){1}${DESTINATION_PORT}\s* -> Continue
  ^\s*\d+\s+(\w+\.\w+\.\w+|\w+\:\w+\:\w+\:\w+\:\w+\:\w+)\s+\w+\s+((\w+\/\w+|\w+)\s){2}${DESTINATION_PORT}\s* -> Continue
  ^\s*\d+\s+(\w+\.\w+\.\w+|\w+\:\w+\:\w+\:\w+\:\w+\:\w+)\s+\w+\s+((\w+\/\w+|\w+)\s){3}${DESTINATION_PORT}\s* -> Continue
  ^\s*\d+\s+(\w+\.\w+\.\w+|\w+\:\w+\:\w+\:\w+\:\w+\:\w+)\s+\w+\s+((\w+\/\w+|\w+)\s){4}${DESTINATION_PORT}$$ -> Continue
  ^\s*Total\s+[M|m]ac -> Start
  ^\s*-+\s+.*-*$$
  ^\s*$$`
