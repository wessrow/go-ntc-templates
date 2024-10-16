package cisco_asa

type ShowXlate struct {
	Destination_intf string   `json:"DESTINATION_INTF"`
	Destination      []string `json:"DESTINATION"`
	Source_intf      string   `json:"SOURCE_INTF"`
	Source           []string `json:"SOURCE"`
}

var ShowXlate_Template string = `Value SOURCE_INTF (\S+)
Value List SOURCE ([0-9a-fA-F:\./]{7,})
Value DESTINATION_INTF (\S+)
Value List DESTINATION ([0-9a-fA-F:\./]{7,})

Start
  ^\d+\s+in\s+use
  ^Flags:
  ^\S+\s+-\s+\S+
  ^NAT\s+from\s+${SOURCE_INTF}:${SOURCE} -> Continue
  ^NAT\s+from\s+\S+:[0-9a-fA-F:\./]+,\s+${SOURCE} -> Continue
  ^NAT\s+from\s+\S+:(?:[0-9a-fA-F:\./]+,\s+){2}${SOURCE},\s*$$
  ^NAT\s+from\s+\S+:(?:[0-9a-fA-F:\./]+,\s+){2}${SOURCE} -> Continue
  ^${SOURCE} -> Continue
  ^[0-9a-fA-F:\./]+,\s+${SOURCE} -> Continue
  ^(?:[0-9a-fA-F:\./]+,\s+){2}${SOURCE}
  ^.+to\s+${DESTINATION_INTF}:${DESTINATION}\s*$$ -> Dest
  ^.+to\s+${DESTINATION_INTF}:${DESTINATION} -> Continue
  ^.+to\s+\S+:[0-9a-fA-F:\./]+\s+${DESTINATION}\s*$$ -> Dest
  ^.+to\s+\S+:[0-9a-fA-F:\./]+\s+${DESTINATION} -> Continue
  ^.+to\s+\S+:(?:[0-9a-fA-F:\./]+\s+){2}${DESTINATION}\s*$$ -> Dest
  ^.*flags -> Record
  ^\s*$$
  ^. -> Error

Dest
  ^${DESTINATION} -> Continue
  ^[0-9a-fA-F:\./]+,\s+${DESTINATION} -> Continue
  ^(?:[0-9a-fA-F:\./]+,\s+){2}${DESTINATION}\s*$$
  ^(?:[0-9a-fA-F:\./]+,\s+){2}${DESTINATION}
  ^.*flags -> Record Start
  ^\s*$$
  ^. -> Error
`
