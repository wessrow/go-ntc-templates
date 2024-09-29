package models

type PaloaltoPanosShowRunningSecurityPolicy struct {
	Name	string	`json:"NAME"`
	From	string	`json:"FROM"`
	Source	string	`json:"SOURCE"`
	Source_region	string	`json:"SOURCE_REGION"`
	To	string	`json:"TO"`
	Destination	string	`json:"DESTINATION"`
	Destination_region	string	`json:"DESTINATION_REGION"`
	User	string	`json:"USER"`
	Category	string	`json:"CATEGORY"`
	Application_service	string	`json:"APPLICATION_SERVICE"`
	Action	string	`json:"ACTION"`
	Icmp_unreachable	string	`json:"ICMP_UNREACHABLE"`
	Terminal	string	`json:"TERMINAL"`
	Type	string	`json:"TYPE"`
}

var PaloaltoPanosShowRunningSecurityPolicy_Template = `Value Filldown NAME (.*?)
Value Required FROM (\S+)
Value SOURCE (\S+)
Value SOURCE_REGION (\S+)
Value TO (\S+)
Value DESTINATION ([\S+\s+]+)
Value DESTINATION_REGION (\S+)
Value USER (\S+)
Value CATEGORY (\S+)
Value APPLICATION_SERVICE ([\S+\s+]+)
Value ACTION (\S+)
Value ICMP_UNREACHABLE (\S+)
Value TERMINAL (\S+)
Value TYPE (\S+)

Start
  ^${NAME}\s+\{
  ^\s+from\s+${FROM};
  ^\s+source\s+${SOURCE};
  ^\s+source-region\s+${SOURCE_REGION}; 
  ^\s+to\s+${TO};
  ^\s+destination\s+${DESTINATION};
  ^\s+destination-region\s+${DESTINATION_REGION};
  ^\s+user\s+${USER};
  ^\s+category\s+${CATEGORY};
  ^\s+application/service\s+${APPLICATION_SERVICE};
  ^\s+action\s+${ACTION};
  ^\s+icmp-unreachable:\s+${ICMP_UNREACHABLE}
  ^\s+terminal\s+${TERMINAL};
  ^\s+type\s+${TYPE};
  ^} -> Record

`