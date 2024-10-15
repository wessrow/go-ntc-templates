package cisco_ios

type ShowCryptoPkiCertificates struct {
	Subject_e              string `json:"SUBJECT_E"`
	Subject_name           string `json:"SUBJECT_NAME"`
	Crl_distribution_point string `json:"CRL_DISTRIBUTION_POINT"`
	Issuer_o               string `json:"ISSUER_O"`
	Issuer_l               string `json:"ISSUER_L"`
	Certificate_type       string `json:"CERTIFICATE_TYPE"`
	Issuer_e               string `json:"ISSUER_E"`
	Subject_hostname       string `json:"SUBJECT_HOSTNAME"`
	Subject_l              string `json:"SUBJECT_L"`
	Subject_st             string `json:"SUBJECT_ST"`
	Subject_c              string `json:"SUBJECT_C"`
	Associated_trustpoints string `json:"ASSOCIATED_TRUSTPOINTS"`
	Issuer_cn              string `json:"ISSUER_CN"`
	Issuer_st              string `json:"ISSUER_ST"`
	Subject_serial_number  string `json:"SUBJECT_SERIAL_NUMBER"`
	Serial_number          string `json:"SERIAL_NUMBER"`
	Certificate_usage      string `json:"CERTIFICATE_USAGE"`
	Subject_o              string `json:"SUBJECT_O"`
	Renew_date             string `json:"RENEW_DATE"`
	Issuer_ou              string `json:"ISSUER_OU"`
	Issuer_c               string `json:"ISSUER_C"`
	Subject_ou             string `json:"SUBJECT_OU"`
	End_date               string `json:"END_DATE"`
	Storage                string `json:"STORAGE"`
	Status                 string `json:"STATUS"`
	Subject_cn             string `json:"SUBJECT_CN"`
	Start_date             string `json:"START_DATE"`
}

var ShowCryptoPkiCertificates_Template string = `Value Required CERTIFICATE_TYPE ([CA]*\s*Certificate)
Value STATUS ([\s\S]+?)
Value SERIAL_NUMBER (\S+)
Value CERTIFICATE_USAGE ([\s\S]+?)
Value ISSUER_E (\S+)
Value ISSUER_CN ([\s\S]+)
Value ISSUER_OU ([\s\S]+)
Value ISSUER_O ([\s\S]+)
Value ISSUER_L ([\s\S]+)
Value ISSUER_ST ([\s\S]+)
Value ISSUER_C ([\s\S]+)
Value SUBJECT_E (\S+)
Value SUBJECT_CN ([\s\S]+)
Value SUBJECT_OU ([\s\S]+)
Value SUBJECT_O ([\s\S]+)
Value SUBJECT_L ([\s\S]+)
Value SUBJECT_ST ([\s\S]+)
Value SUBJECT_C ([\s\S]+)
Value SUBJECT_NAME (\S+)
Value SUBJECT_SERIAL_NUMBER (\S+)
Value SUBJECT_HOSTNAME (\S+)
Value CRL_DISTRIBUTION_POINT (\S+)
Value START_DATE ([\s\S]+?)
Value END_DATE ([\s\S]+?)
Value RENEW_DATE ([\s\S]+?)
Value ASSOCIATED_TRUSTPOINTS (\S+)
Value STORAGE (\S+)

Start
  ^\s*[CA]*\s*Certificate\s*$$ -> Continue.Record
  ^\s*${CERTIFICATE_TYPE}\s*$$
  ^\s*Status:\s*${STATUS}\s*$$
  ^\s*Certificate\s*Serial\s*Number\s*\(hex\):\s*${SERIAL_NUMBER}\s*$$
  ^\s*Certificate\s*Usage:\s*${CERTIFICATE_USAGE}\s*$$
  ^\s*Issuer:\s*$$ -> Issuer
  ^\s*start\s*date:\s*${START_DATE}\s*$$
  ^\s*end\s*date:\s*${END_DATE}\s*$$
  ^\s*renew\s*date:\s*${RENEW_DATE}\s*$$
  ^\s*Associated\s*Trustpoints:\s*${ASSOCIATED_TRUSTPOINTS}\s*$$
  ^\s*Storage:\s*${STORAGE}\s*$$
  ^\s*$$
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is
  ^. -> Error

Issuer
  ^\s*e=${ISSUER_E}\s*$$
  ^\s*cn=${ISSUER_CN}\s*$$
  ^\s*ou=${ISSUER_OU}\s*$$
  ^\s*o=${ISSUER_O}\s*$$
  ^\s*l=${ISSUER_L}\s*$$
  ^\s*st=${ISSUER_ST}\s*$$
  ^\s*c=${ISSUER_C}\s*$$
  ^\s*Subject:\s*$$ -> Subject
  ^. -> Error

Subject
  ^\s*e=${SUBJECT_E}\s*$$
  ^\s*cn=${SUBJECT_CN}\s*$$
  ^\s*ou=${SUBJECT_OU}\s*$$
  ^\s*o=${SUBJECT_O}\s*$$
  ^\s*l=${SUBJECT_L}\s*$$
  ^\s*st=${SUBJECT_ST}\s*$$
  ^\s*c=${SUBJECT_C}\s*$$
  ^\s*Name:\s*${SUBJECT_NAME}\s*$$
  ^\s*Serial\s*Number:\s*${SUBJECT_SERIAL_NUMBER}\s*$$
  ^\s*serialNumber=\S+\s*$$
  ^\s*hostname=${SUBJECT_HOSTNAME}\s*$$
  ^\s*CRL Distribution Points:\s*$$ -> CRLDistPoints
  ^\s*Validity\s*Date:\s*$$ -> Start
  ^. -> Error

CRLDistPoints
  ^\s*${CRL_DISTRIBUTION_POINT}\s*$$
  ^\s*Validity\s*Date:\s*$$ -> Start
  ^. -> Error
`
