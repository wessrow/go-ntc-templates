package ericsson_ipos 

type ShowVersion struct {
	Version	string	`json:"VERSION"`
	Uptime	string	`json:"UPTIME"`
}

var ShowVersion_Template = `Value VERSION (\S[^\[]+)
Value UPTIME (\d+\s+\S+\s+\d+\s+\S+\s+\d+\s+\S+)

Start
  ^Ericsson\s+IPOS\s+Version\s+IPOS-v${VERSION}-Release
  ^.+Up\s+Time\s+-\s+${UPTIME}
`