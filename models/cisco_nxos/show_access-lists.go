package cisco_nxos

type ShowAccessLists struct {
	Name        string `json:"NAME"`
	Action      string `json:"ACTION"`
	Source      string `json:"SOURCE"`
	Range       string `json:"RANGE"`
	Port        string `json:"PORT"`
	Modifier    string `json:"MODIFIER"`
	Sn          string `json:"SN"`
	Protocol    string `json:"PROTOCOL"`
	Destination string `json:"DESTINATION"`
	Remark      string `json:"REMARK"`
}

var ShowAccessLists_Template string = `# NX-OS ACL structure is quite complex and includes IP, IPv6 and MAC ACL types.
#
# IP/IPv6 ACL:
# 	SEQUENCE_NUMBER ACTION PROTOCOL SOURCE [PORT id] [RANGE start finish] DESTINATION [MODIFIER]
#
# MAC ACL:
#	SEQUENCE_NUMBER ACTION SOURCE DESTINATION [MODIFIER]
#
Value Filldown NAME (.*)
Value Required SN (\d+)
Value ACTION (\w+)
Value PROTOCOL ([^any][\w+\d+]+)
# SOURCE RegEx must be able to catch both MAC and IP addresses including masks. 'any' and 'address group' are possible too.
Value SOURCE ([\w+\d+]+\.[\w+\d+]+\.[\w+\d+]+\s[\w+\d+]+\.[\w+\d+]+\.[\w+\d+]+|any|\d+\.\d+\.\d+\.\d+/\d+|addrgroup\s\S+|\S+)
# RANGE RegEx catches the first and last protocol to match in the interval range.
Value RANGE (\S+\s\S+)
# We can specify protocols to match. 'eq', 'gt', 'lt' and 'neq' are supported.
Value PORT (eq\s\S+|gt\s\S+|lt\s\S+|neq\s\S+)
# SOURCE RegEx must be able to catch both MAC and IP addresses including masks. 'any' and 'address group' are possible too.
Value DESTINATION ([\w+\d+]+\.[\w+\d+]+\.[\w+\d+]+\s[\w+\d+]+\.[\w+\d+]+\.[\w+\d+]+|any|\d+\.\d+\.\d+\.\d+/\d+|addrgroup\s\S+|\S+)
# What it remains is some ACL modifier used to do fancy stuff.
Value MODIFIER (.*)
Value REMARK (.*)

Start
  # ACL name is matched not starting with space. FILLDOWN, other line's inherit last NAME
  ^[^\s].*\slist\s+${NAME}$$
  # Catch remarks as uniq type, capture Action first and then Action=remark
  ^\s+${SN}\s+${ACTION}\s -> Continue
  ^\s+${SN}\s+remark\s${REMARK}$$ -> Record
  # This line matches all the possible combination of fields specifed above.
  ^\s+${SN}\s+${ACTION}(\s+${PROTOCOL})*\s+${SOURCE}(\s+${PORT})*(\s+range\s${RANGE})*\s+${DESTINATION}(\s+${MODIFIER})* -> Record
  # Drop lines we are not interested in
  ^\s+statistics per-entry
  ^\s*$$
  ^. -> Error`
