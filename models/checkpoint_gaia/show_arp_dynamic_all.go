package checkpoint_gaia

type ShowArpDynamicAll struct {
	Ip_address  string `json:"IP_ADDRESS"`
	Mac_address string `json:"MAC_ADDRESS"`
}

var ShowArpDynamicAll_Template string = `Value IP_ADDRESS (\d+\.\d+\.\d+\.\d+)
Value MAC_ADDRESS (\S+)

Start
  ^Dynamic\sArp\sParameters
  ^IP\sAddress\s+Mac\sAddress
  ^${IP_ADDRESS}\s+${MAC_ADDRESS} -> Record
 ^\s*$$
 ^. -> Error
`
