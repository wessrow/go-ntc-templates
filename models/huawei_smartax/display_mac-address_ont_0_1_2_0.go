package huawei_smartax

type DisplayMacAddressOnt0120 struct {
	Ont_id        string `json:"ONT_ID"`
	Ont_port_type string `json:"ONT_PORT_TYPE"`
	Ont_port_id   string `json:"ONT_PORT_ID"`
	Mac           string `json:"MAC"`
	Vlan          string `json:"VLAN"`
	Fsp           string `json:"FSP"`
}

var DisplayMacAddressOnt0120_Template string = `Value FSP (\d+\/\s*\d+\/\d+)
Value ONT_ID (\d+)
Value ONT_PORT_TYPE (\S+)
Value ONT_PORT_ID (\d+)
Value MAC (\w+-\w+-\w+)
Value VLAN (\d+)

Start
  ^\s*-
  ^\s*F\/S\/P\s+ONTID\s+ONT\s+ONT\s+MAC-ADDRESS\s+VLAN
  ^\s*port-type\s+port-ID -> Macs
  ^\s*$$
  ^. -> Error

Macs
  ^\s*${FSP}\s+${ONT_ID}\s+${ONT_PORT_TYPE}\s+${ONT_PORT_ID}\s+${MAC}\s+(-|${VLAN}) -> Record
  ^\s*-
  ^\s*Total:\s*\d+\s*$$ -> EOF
  ^\s*$$
  ^. -> Error
`
