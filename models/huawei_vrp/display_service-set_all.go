package huawei_vrp 

type DisplayServiceSetAll struct {
	Id	string	`json:"ID"`
	Name	string	`json:"NAME"`
	Ssid	string	`json:"SSID"`
}

var DisplayServiceSetAll_Template = `Value ID (\d+)
Value NAME (\w+)
Value SSID (\w+)

Start
  ^\s*-+ -> Next
  ^\s*ID.+$$ -> Next
  ^\s*${ID}\s+${NAME}\s+${SSID}\s*$$ -> Record
  ^\s*Total:.+$$ -> Next
  ^.*$$ -> Error
`