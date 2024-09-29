package arista_eos 

type ShowSnmpCommunity struct {
	Name	string	`json:"NAME"`
	Access	string	`json:"ACCESS"`
	View	string	`json:"VIEW"`
	Acl	string	`json:"ACL"`
}

var ShowSnmpCommunity_Template = `Value NAME (\S+)
Value ACCESS (\S+)
Value VIEW (\S+)
Value ACL (\S+)

Start
  ^Community\sview:\s+${VIEW}
  ^.*list:\s+${ACL}
  ^Community\sname:\s+${NAME}
  ^Community\saccess:\s+${ACCESS} 
  ^.* -> Record
  #^Community\svies:\s+${VIEW}
  #^Access\slist:\s+${ACL} -> Record
`