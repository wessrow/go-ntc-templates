package models

type CiscoIosShowCryptoPkiCertificates struct {
	Certificate_type	string	`json:"CERTIFICATE_TYPE"`
	Status	string	`json:"STATUS"`
	Serial_number	string	`json:"SERIAL_NUMBER"`
	Certificate_usage	string	`json:"CERTIFICATE_USAGE"`
	Issuer_e	string	`json:"ISSUER_E"`
	Issuer_cn	string	`json:"ISSUER_CN"`
	Issuer_ou	string	`json:"ISSUER_OU"`
	Issuer_o	string	`json:"ISSUER_O"`
	Issuer_l	string	`json:"ISSUER_L"`
	Issuer_st	string	`json:"ISSUER_ST"`
	Issuer_c	string	`json:"ISSUER_C"`
	Subject_e	string	`json:"SUBJECT_E"`
	Subject_cn	string	`json:"SUBJECT_CN"`
	Subject_ou	string	`json:"SUBJECT_OU"`
	Subject_o	string	`json:"SUBJECT_O"`
	Subject_l	string	`json:"SUBJECT_L"`
	Subject_st	string	`json:"SUBJECT_ST"`
	Subject_c	string	`json:"SUBJECT_C"`
	Subject_name	string	`json:"SUBJECT_NAME"`
	Subject_serial_number	string	`json:"SUBJECT_SERIAL_NUMBER"`
	Subject_hostname	string	`json:"SUBJECT_HOSTNAME"`
	Crl_distribution_point	string	`json:"CRL_DISTRIBUTION_POINT"`
	Start_date	string	`json:"START_DATE"`
	End_date	string	`json:"END_DATE"`
	Renew_date	string	`json:"RENEW_DATE"`
	Associated_trustpoints	string	`json:"ASSOCIATED_TRUSTPOINTS"`
	Storage	string	`json:"STORAGE"`
}