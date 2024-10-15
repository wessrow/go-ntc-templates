package oneaccess_oneos

type ShowSnmpCommunity struct {
	Name          string `json:"NAME"`
	Security_name string `json:"SECURITY_NAME"`
	Acl           string `json:"ACL"`
	Group         string `json:"GROUP"`
}

var ShowSnmpCommunity_Template string = `Value NAME (\S+)
Value SECURITY_NAME (\S+)
Value ACL (\S+)
Value GROUP (\S+)
 
Start
  ^SNMP.*community -> Continue.Record
  ^SNMP\s${SECURITY_NAME}\s+community\s*:\s+${NAME}.*
  ^SNMP\s\S+\s+community\s*:\s+\S+\s+v2group\s+${GROUP}
  ^SNMP\saccess\scontrol\slists\s*:\s*${ACL}
  ^\s*$$
  ^. -> Error
`
