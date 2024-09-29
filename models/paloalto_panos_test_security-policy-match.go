package models

type PaloaltoPanosTestSecurityPolicyMatch struct {
	Name	string	`json:"NAME"`
	Index	string	`json:"INDEX"`
	From	string	`json:"FROM"`
	Source	string	`json:"SOURCE"`
	Source_region	string	`json:"SOURCE_REGION"`
	To	string	`json:"TO"`
	Destination	string	`json:"DESTINATION"`
	Destination_region	string	`json:"DESTINATION_REGION"`
	User	string	`json:"USER"`
	Source_device	string	`json:"SOURCE_DEVICE"`
	Destination_device	string	`json:"DESTINATION_DEVICE"`
	Category	string	`json:"CATEGORY"`
	Application_service	string	`json:"APPLICATION_SERVICE"`
	Action	string	`json:"ACTION"`
	Icmp_unreachable	string	`json:"ICMP_UNREACHABLE"`
	Terminal	string	`json:"TERMINAL"`
}

var PaloaltoPanosTestSecurityPolicyMatch_Template = `Value Required NAME (.*?)
Value Required INDEX (\d+)
Value FROM (\S+)
Value SOURCE (\S+|\[.*\])
Value SOURCE_REGION (\S+)
Value TO (\S+)
Value DESTINATION (\S+|\[.*\])
Value DESTINATION_REGION (\S+)
Value USER (\S+)
Value SOURCE_DEVICE (\S+)
Value DESTINATION_DEVICE (\S+)
Value CATEGORY (\S+)
Value APPLICATION_SERVICE (\S+|\[.*\])
Value ACTION (\S+)
Value ICMP_UNREACHABLE (\S+)
Value TERMINAL (\S+)

Start
  ^\"${NAME};\s+index:\s+${INDEX}.\s+{
  ^\s+from\s+${FROM};
  ^\s+source\s+${SOURCE};
  ^\s+source-region\s+${SOURCE_REGION}; 
  ^\s+to\s+${TO};
  ^\s+destination\s+${DESTINATION};
  ^\s+destination-region\s+${DESTINATION_REGION};
  ^\s+user\s+${USER};
  ^\s+source-device\s+${SOURCE_DEVICE};
  ^\s+destinataion-device\s+${DESTINATION_DEVICE};
  ^\s+category\s+${CATEGORY};
  ^\s+application/service\s+${APPLICATION_SERVICE};
  ^\s+action\s+${ACTION};
  ^\s+icmp-unreachable:\s+${ICMP_UNREACHABLE}
  ^\s+terminal\s+${TERMINAL};
  ^} -> Record
  ^\s*$$
  ^. -> Error
`