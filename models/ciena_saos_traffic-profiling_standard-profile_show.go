package models

type CienaSaosTrafficProfilingStandardProfileShow struct {
	Port	string	`json:"PORT"`
	Id	string	`json:"ID"`
	Profile_name	string	`json:"PROFILE_NAME"`
	Role	string	`json:"ROLE"`
	Parent_child	string	`json:"PARENT_CHILD"`
	Cir	string	`json:"CIR"`
	Eir	string	`json:"EIR"`
	Cbs	string	`json:"CBS"`
	Ebs	string	`json:"EBS"`
	Classifier_a	string	`json:"CLASSIFIER_A"`
	Classifier_b	string	`json:"CLASSIFIER_B"`
}