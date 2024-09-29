package aruba_aoscx 

type ShowVsfDetail struct {
	Member_id	string	`json:"MEMBER_ID"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Status	string	`json:"STATUS"`
	Cpu	string	`json:"CPU"`
	Memory	string	`json:"MEMORY"`
	Vsf_link_1	string	`json:"VSF_LINK_1"`
	Vsf_link_2	string	`json:"VSF_LINK_2"`
}

var ShowVsfDetail_Template = `Value MEMBER_ID (\S+)
Value MAC_ADDRESS (\S+)
Value STATUS (.*)
Value CPU (\d+%)
Value MEMORY (\d+%)
Value VSF_LINK_1 (.*)
Value VSF_LINK_2 (.*)

Start
  ^VSF\s*${MEMBER_ID}
  ^Member.* -> Continue.Record
  ^Member\s+ID\s+:\s+${MEMBER_ID}
  ^\s+MAC\s+Address\s+:\s+${MAC_ADDRESS}                                                                          
  ^\s+Status\s+:\s+${STATUS}                                                                       
  ^\s+CPU\sUtilization\s+:\s+${CPU}                                                                      
  ^\s+Memory\sUtilization\s+:\s+${MEMORY}
  ^\s+VSF\s+Link\s+1\s+:${VSF_LINK_1}
  ^\s+VSF\s+Link\s+2\s+:${VSF_LINK_2}
`