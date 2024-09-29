package cisco_ios 

type ShowAccessList struct {
	Name	string	`json:"NAME"`
	Type	string	`json:"TYPE"`
	Sn	string	`json:"SN"`
	Action	string	`json:"ACTION"`
	Protocol	string	`json:"PROTOCOL"`
	Source	string	`json:"SOURCE"`
	Operator_source_port	string	`json:"OPERATOR_SOURCE_PORT"`
	Operator_destination_port	string	`json:"OPERATOR_DESTINATION_PORT"`
	Source_port	string	`json:"SOURCE_PORT"`
	Destination_port	string	`json:"DESTINATION_PORT"`
	Destination	string	`json:"DESTINATION"`
	Modifier	string	`json:"MODIFIER"`
	Matches	string	`json:"MATCHES"`
}

var ShowAccessList_Template = `# IOS ACL structure is quite complex.
# 
# Extended IP ACL:
# 	SEQUENCE_NUMBER ACTION PROTOCOL SOURCE [PORT id] [RANGE start finish] DESTINATION [MODIFIER]
#
# Standard ACL:
#	SEQUENCE_NUMBER ACTION SOURCE DESTINATION [MODIFIER]
#	 
#
Value Filldown NAME (\S+)
Value Filldown TYPE (\S+)
Value Required SN (\d+)
Value ACTION (\w+)
Value PROTOCOL (\w+)
#
# SOURCE RegEx must be able to catch every possible source combination including masks. 'any' and 'host' and 'wildcard bits' are possible too.
#
Value SOURCE (host\s\d+\.\d+\.\d+\.\d+|any|\d+\.\d+\.\d+\.\d+\s\d+\.\d+\.\d+\.\d+|\d+\.\d+\.\d+\.\d+,\s+wildcard bits\s\d+\.\d+\.\d+\.\d+|\d+\.\d+\.\d+\.\d+)
#
# We can specify protocols to match. 'eq', 'gt', 'lt', 'range' and 'neq' are supported.
#
Value OPERATOR_SOURCE_PORT (eq|gt|lt|neq|range)
Value OPERATOR_DESTINATION_PORT (eq|gt|lt|neq|range)
#
# This can be either a single port or a list.
#
Value SOURCE_PORT ([\s\S]+)
Value DESTINATION_PORT ([\s\S]+)
#
# DESTINATION RegEx must be able to catch every possible source combination including masks. 'any' and 'host' and 'wildcard bits' are possible too.
# 
Value DESTINATION (host\s\d+\.\d+\.\d+\.\d+|any|\d+\.\d+\.\d+\.\d+\s\d+\.\d+\.\d+\.\d+|\d+\.\d+\.\d+\.\d+,\s+wildcard bits\s\d+\.\d+\.\d+\.\d+)
Value MODIFIER (log|tos normal|eq ftp|gt 1024)
Value MATCHES (\d+)

Start
  ^${TYPE}.*list\s+${NAME}
  ^\s*${SN}\s+${ACTION}\s+${PROTOCOL}\s+${SOURCE}\s+${OPERATOR_SOURCE_PORT}\s${SOURCE_PORT}\s${DESTINATION}\s${OPERATOR_DESTINATION_PORT}\s${DESTINATION_PORT} -> Record
  ^\s*${SN}\s+${ACTION}\s+${PROTOCOL}\s+${SOURCE}\s+${DESTINATION}(\s+${OPERATOR_SOURCE_PORT})*(\s${SOURCE_PORT})* -> Record
  ^\s*${SN}\s+${ACTION}\s+${PROTOCOL}\s+${SOURCE}(\s+${OPERATOR_SOURCE_PORT})*(\s${SOURCE_PORT})*\s+${DESTINATION}(\s+${MODIFIER})*(\s\(*(\s\(${MATCHES}\smatches\))*\smatches\))* -> Record
  ^\s*${SN}\s+${ACTION}\s+${SOURCE}(\s+${MODIFIER})*(\s\(${MATCHES}\smatches\))* -> Record
  ^Load\s+for\s+
  ^Time\s+source\s+is
`