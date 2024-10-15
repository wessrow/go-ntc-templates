package ciena_saos

type TrafficProfilingStandardProfileShow struct {
	Cir          string `json:"CIR"`
	Eir          string `json:"EIR"`
	Cbs          string `json:"CBS"`
	Port         string `json:"PORT"`
	Id           string `json:"ID"`
	Profile_name string `json:"PROFILE_NAME"`
	Role         string `json:"ROLE"`
	Parent_child string `json:"PARENT_CHILD"`
	Ebs          string `json:"EBS"`
	Classifier_a string `json:"CLASSIFIER_A"`
	Classifier_b string `json:"CLASSIFIER_B"`
}

var TrafficProfilingStandardProfileShow_Template string = `Value PORT ([0-9A-Za-z\.]+)
Value ID (\d+)
Value PROFILE_NAME (\S+)
Value ROLE (\S+)
Value PARENT_CHILD (\d+)
Value CIR (\d+)
Value EIR (\d+)
Value CBS (\d+)
Value EBS (\d+)
Value CLASSIFIER_A (\S+)
Value CLASSIFIER_B (\S+)


Start
  ^\+-+\s*STANDARD\s*PROFILE\s*TABLE\s*-+\+
  ^\|\s*Port
  ^\|\s*\|\s*ID
  ^\+-+\+
  ^\|\s*${PORT}\s*\|\s*${ID}\s*\|\s*${PROFILE_NAME}\s*\|\s*${ROLE}\s*\|\s*${PARENT_CHILD}\s*\|\s*${CIR}\s*\|\s*${EIR}\s*\|\s*${CBS}\s*\|\s*${EBS}\s*\|\s*${CLASSIFIER_A}\s*\|\s*${CLASSIFIER_B}\s*\| -> Record
  ^\+-.*
  ^\s*$$
  ^. -> Error
  `
