package zte_zxros 

type ShowVersion struct {
	Hardware	string	`json:"HARDWARE"`
	Version	string	`json:"VERSION"`
	Uptime	string	`json:"UPTIME"`
}

var ShowVersion_Template = `Value HARDWARE (.+)
Value VERSION (.+)
Value UPTIME (.+?)


Start
  ^ZXCTN${HARDWARE}
  ^ZTE.+Software,\s+Version:\s+${VERSION},\s+Release\s+software
  ^.+uptime\s+is\s+${UPTIME}\s*$$ -> EOF
  ^Copyright
  ^System\s+image
  ^\s*$$
  ^. -> Error
`