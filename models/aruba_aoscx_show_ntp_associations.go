package models

type ArubaAoscxShowNtpAssociations struct {
	Id	string	`json:"ID"`
	Name	string	`json:"NAME"`
	Remote	string	`json:"REMOTE"`
	Ref_id	string	`json:"REF_ID"`
	Stratum	string	`json:"STRATUM"`
	Last	string	`json:"LAST"`
	Poll	string	`json:"POLL"`
	Reach	string	`json:"REACH"`
}