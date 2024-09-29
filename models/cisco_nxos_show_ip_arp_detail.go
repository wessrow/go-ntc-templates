package models

type CiscoNxosShowIpArpDetail struct {
	Ip_address	string	`json:"IP_ADDRESS"`
	Age	string	`json:"AGE"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Interface	string	`json:"INTERFACE"`
	Physical_interface	string	`json:"PHYSICAL_INTERFACE"`
}

var CiscoNxosShowIpArpDetail_Template = `Value IP_ADDRESS (\d+\.\d+\.\d+\.\d+)
Value AGE (\S+)
Value MAC_ADDRESS (\S+)
Value INTERFACE (\S+)
Value PHYSICAL_INTERFACE (\S+)

Start
  #Ignore junk
  ^.+\s-\sAdjacencies\s
  ^\s*IP\sARP\sTable\s
  ^\s*Total\snumber\sof\sentries:
  ^\s*Address\s+Age\s+MAC\s+Address\s+Interface\s+Physical\sInterface\s*$$ -> Data
  ^. -> Error

Data
  ^${IP_ADDRESS}\s+${AGE}\s+${MAC_ADDRESS}\s+${INTERFACE}\s+${PHYSICAL_INTERFACE}\s*$$ -> Record
  ^. -> Error

`