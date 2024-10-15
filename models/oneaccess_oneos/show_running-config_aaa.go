package oneaccess_oneos

type ShowRunningConfigAaa struct {
	Aaa_group    string   `json:"AAA_GROUP"`
	Aaa_protocol string   `json:"AAA_PROTOCOL"`
	Aaa_command  string   `json:"AAA_COMMAND"`
	Aaa_config   []string `json:"AAA_CONFIG"`
	Aaa_servers  []string `json:"AAA_SERVERS"`
}

var ShowRunningConfigAaa_Template string = `Value Filldown AAA_GROUP (\S+)
Value Required AAA_PROTOCOL (authentication|authorization|accounting)
Value AAA_COMMAND (\w+)
Value List AAA_CONFIG (.*)
Value Filldown,List AAA_SERVERS (\S+)

Start
  ^\s*aaa\s+group\s+server\s+tacacs\s+${AAA_GROUP}
  ^\s+server\s+${AAA_SERVERS}\s*$$
  ^\s*aaa\s+${AAA_PROTOCOL}\s+${AAA_COMMAND}\s+${AAA_CONFIG}\s*$$ -> Record
  ^exit
  ^\s*$$
  ^. -> Error
`
