package paloalto_panos

type TestSecurityPolicyMatch struct {
	To                  string `json:"TO"`
	Destination_region  string `json:"DESTINATION_REGION"`
	Icmp_unreachable    string `json:"ICMP_UNREACHABLE"`
	Name                string `json:"NAME"`
	Source_device       string `json:"SOURCE_DEVICE"`
	Source_region       string `json:"SOURCE_REGION"`
	Destination         string `json:"DESTINATION"`
	User                string `json:"USER"`
	Category            string `json:"CATEGORY"`
	Application_service string `json:"APPLICATION_SERVICE"`
	Terminal            string `json:"TERMINAL"`
	Index               string `json:"INDEX"`
	Source              string `json:"SOURCE"`
	Destination_device  string `json:"DESTINATION_DEVICE"`
	Action              string `json:"ACTION"`
	From                string `json:"FROM"`
}

var TestSecurityPolicyMatch_Template string = `Value Required NAME (.*?)
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
  ^. -> Error`
