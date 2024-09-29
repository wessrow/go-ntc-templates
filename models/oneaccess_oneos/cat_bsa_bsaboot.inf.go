package oneaccess_oneos 

type CatBsaBsabootInf struct {
	Bootfile	string	`json:"BOOTFILE"`
	Bootfolder	string	`json:"BOOTFOLDER"`
	Configfile	string	`json:"CONFIGFILE"`
	Configfolder	string	`json:"CONFIGFOLDER"`
}

var CatBsaBsabootInf_Template = `Value BOOTFILE (\S+)
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