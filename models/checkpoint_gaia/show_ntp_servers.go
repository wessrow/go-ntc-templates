package checkpoint_gaia

type ShowNtpServers struct {
	Ip_address string `json:"IP_ADDRESS"`
	Type       string `json:"TYPE"`
	Version    string `json:"VERSION"`
}

var ShowNtpServers_Template string = `Value IP_ADDRESS ((?:[0-9]{1,3}\.){3}[0-9]{1,3})
Value TYPE (\S+)
Value VERSION (\S+)

Start
  ^IP\sAddress\s+Type\s+Version\s*$$
  ^${IP_ADDRESS}\s+${TYPE}\s+${VERSION}\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`
