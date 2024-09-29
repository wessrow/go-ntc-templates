package models

type CheckpointGaiaFwStat struct {
	Policy	string	`json:"POLICY"`
	Policy_install_date	string	`json:"POLICY_INSTALL_DATE"`
}

var CheckpointGaiaFwStat_Template = `Value POLICY (\S+)
Value POLICY_INSTALL_DATE ([0-9]{1,2}[a-zA-Z]{3}[0-9]{4}\s\d{1,2}\:\d{1,2}\:\d{1,2})

Start
  ^\w+\s${POLICY}\s+${POLICY_INSTALL_DATE}\s+ -> Record

`