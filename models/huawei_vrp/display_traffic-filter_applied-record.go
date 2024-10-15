package huawei_vrp

type DisplayTrafficFilterAppliedRecord struct {
	Instance_type string `json:"INSTANCE_TYPE"`
	Instance      string `json:"INSTANCE"`
	Direction     string `json:"DIRECTION"`
	Ip_version    string `json:"IP_VERSION"`
	Acl_nb        string `json:"ACL_NB"`
	Acl_name      string `json:"ACL_NAME"`
}

var DisplayTrafficFilterAppliedRecord_Template string = `Value Filldown INSTANCE_TYPE (Interface|Traffic profile)
Value Required INSTANCE (\S+)
Value Required DIRECTION (inbound|outbound)
Value IP_VERSION ((ip|IP)v(4|6))
Value ACL_NB (\d+)
Value ACL_NAME (\S+)

Start
  ^-+ -> Next
  ^\s*${INSTANCE_TYPE}\s+Direction\s+AppliedRecord\s*$$
  ^\s*${INSTANCE}\s+${DIRECTION}\s+(${IP_VERSION}\s+)?(ACL|acl)\s+name\s+${ACL_NAME} -> Record
  ^\s*${INSTANCE}\s+${DIRECTION}\s+(${IP_VERSION}\s+)?(ACL|acl)\s+${ACL_NB} -> Record
  ^\s*Total:\d+\s*$$ -> Next
  ^.*$$ -> Error
`
