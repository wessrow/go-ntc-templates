package models

type OneaccessOneosCatBsaBsabootInf struct {
	Bootfile	string	`json:"BOOTFILE"`
	Bootfolder	string	`json:"BOOTFOLDER"`
	Configfile	string	`json:"CONFIGFILE"`
	Configfolder	string	`json:"CONFIGFOLDER"`
}

var OneaccessOneosCatBsaBsabootInf_Template = `Value BOOTFILE (\S+)
Value BOOTFOLDER (/BSA/binaries)
Value CONFIGFILE (\S+)
Value CONFIGFOLDER (/BSA/config)

Start
  ^.*${BOOTFOLDER} -> Continue
  ^.*/BSA/binaries/${BOOTFILE}
  ^.*${CONFIGFOLDER} -> Continue
  ^.*/BSA/config/${CONFIGFILE}
  ^.*$$
  ^. -> Error

`