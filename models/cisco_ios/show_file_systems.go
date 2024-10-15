package cisco_ios

type ShowFileSystems struct {
	Free     string `json:"FREE"`
	Type     string `json:"TYPE"`
	Flags    string `json:"FLAGS"`
	Prefixes string `json:"PREFIXES"`
	Size     string `json:"SIZE"`
}

var ShowFileSystems_Template string = `Value SIZE (\S+)
Value FREE (\S+)
Value TYPE (\S+)
Value FLAGS (\S+)
Value PREFIXES (\S+)

Start
  ^\s*Size -> FILE_SYSTEM

FILE_SYSTEM
  ^\s*${SIZE}\s+${FREE}\s+${TYPE}\s+${FLAGS}\s+${PREFIXES}\s*$$ -> Record
  ^\s*$$

EOF
`
