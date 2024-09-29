package cisco_ios 

type ShowFileSystems struct {
	Size	string	`json:"SIZE"`
	Free	string	`json:"FREE"`
	Type	string	`json:"TYPE"`
	Flags	string	`json:"FLAGS"`
	Prefixes	string	`json:"PREFIXES"`
}

var ShowFileSystems_Template = `Value SIZE (\S+)
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