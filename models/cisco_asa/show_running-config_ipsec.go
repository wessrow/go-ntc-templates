package cisco_asa

type ShowRunningConfigIpsec struct {
	Ike_version string   `json:"IKE_VERSION"`
	Name        string   `json:"NAME"`
	Encryption  []string `json:"ENCRYPTION"`
	Auth        []string `json:"AUTH"`
}

var ShowRunningConfigIpsec_Template string = `Value Required IKE_VERSION ([1-2])
Value NAME (\S+)
Value List ENCRYPTION (\S+)
Value List AUTH (\S+)


Start
  ^crypto\s+ipsec\s+ikev -> Continue.Record
  ^crypto\s+ipsec\s+ikev${IKE_VERSION}\s+transform-set\s+${NAME}\s+${ENCRYPTION}\s+${AUTH}\s+ -> Record
  ^crypto\s+ipsec\s+ikev${IKE_VERSION}\s+ipsec-proposal\s+${NAME}\s*
  ^\s+protocol\s+esp\s+encryption\s+\S+\s+\S+\s+\S+\s+\S+\s+\S+\s+\S+\s+\S+\s+\S+\s+\S+\s+\S+\s+${ENCRYPTION}\s* -> Continue
  ^\s+protocol\s+esp\s+encryption\s+\S+\s+\S+\s+\S+\s+\S+\s+\S+\s+\S+\s+\S+\s+\S+\s+\S+\s+${ENCRYPTION}\s* -> Continue
  ^\s+protocol\s+esp\s+encryption\s+\S+\s+\S+\s+\S+\s+\S+\s+\S+\s+\S+\s+\S+\s+\S+\s+${ENCRYPTION}\s* -> Continue
  ^\s+protocol\s+esp\s+encryption\s+\S+\s+\S+\s+\S+\s+\S+\s+\S+\s+\S+\s+\S+\s+${ENCRYPTION}\s* -> Continue
  ^\s+protocol\s+esp\s+encryption\s+\S+\s+\S+\s+\S+\s+\S+\s+\S+\s+\S+\s+${ENCRYPTION}\s* -> Continue
  ^\s+protocol\s+esp\s+encryption\s+\S+\s+\S+\s+\S+\s+\S+\s+\S+\s${ENCRYPTION}\s* -> Continue
  ^\s+protocol\s+esp\s+encryption\s+\S+\s+\S+\s+\S+\s+\S+\s+${ENCRYPTION}\s* -> Continue
  ^\s+protocol\s+esp\s+encryption\s+\S+\s+\S+\s+\S+\s+${ENCRYPTION}\s* -> Continue
  ^\s+protocol\s+esp\s+encryption\s+\S+\s+\S+\s+${ENCRYPTION}\s* -> Continue
  ^\s+protocol\s+esp\s+encryption\s+\S+\s+${ENCRYPTION}\s* -> Continue
  ^\s+protocol\s+esp\s+encryption\s+${ENCRYPTION}\s* -> Continue
  ^\s+protocol\s+esp\s+integrity\s+\S+\s+\S+\s+\S+\s+\S+\s+${AUTH}\s* -> Continue
  ^\s+protocol\s+esp\s+integrity\s+\S+\s+\S+\s+\S+\s+${AUTH}\s* -> Continue
  ^\s+protocol\s+esp\s+integrity\s+\S+\s+\S+\s+${AUTH}\s* -> Continue
  ^\s+protocol\s+esp\s+integrity\s+\S+\s+${AUTH}\s* -> Continue
  ^\s+protocol\s+esp\s+integrity\s+${AUTH}\s* -> Record

EOF


`
