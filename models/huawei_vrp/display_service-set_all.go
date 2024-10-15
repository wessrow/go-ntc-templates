package huawei_vrp

type DisplayServiceSetAll struct {
	Ssid string `json:"SSID"`
	Id   string `json:"ID"`
	Name string `json:"NAME"`
}

var DisplayServiceSetAll_Template string = `Value ID (\d+)
Value NAME (\w+)
Value SSID (\w+)

Start
  ^\s*-+ -> Next
  ^\s*ID.+$$ -> Next
  ^\s*${ID}\s+${NAME}\s+${SSID}\s*$$ -> Record
  ^\s*Total:.+$$ -> Next
  ^.*$$ -> Error
`
