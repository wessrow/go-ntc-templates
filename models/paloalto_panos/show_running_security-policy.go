package paloalto_panos

type ShowRunningSecurityPolicy struct {
	From                string `json:"FROM"`
	To                  string `json:"TO"`
	Destination         string `json:"DESTINATION"`
	Application_service string `json:"APPLICATION_SERVICE"`
	Name                string `json:"NAME"`
	Destination_region  string `json:"DESTINATION_REGION"`
	User                string `json:"USER"`
	Terminal            string `json:"TERMINAL"`
	Action              string `json:"ACTION"`
	Type                string `json:"TYPE"`
	Source              string `json:"SOURCE"`
	Source_region       string `json:"SOURCE_REGION"`
	Category            string `json:"CATEGORY"`
	Icmp_unreachable    string `json:"ICMP_UNREACHABLE"`
}

var ShowRunningSecurityPolicy_Template string = `Value Filldown NAME (.*?)
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
