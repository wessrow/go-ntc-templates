package models

type OneaccessOneosShowSoftFileInfo struct {
	Status	string	`json:"STATUS"`
	Filename	string	`json:"FILENAME"`
	Filetype	string	`json:"FILETYPE"`
	Software_version	string	`json:"SOFTWARE_VERSION"`
	Software_date	string	`json:"SOFTWARE_DATE"`
	Filesize	string	`json:"FILESIZE"`
	Header_checksum	string	`json:"HEADER_CHECKSUM"`
	Computed_checksum	string	`json:"COMPUTED_CHECKSUM"`
	Target_device	string	`json:"TARGET_DEVICE"`
}

var OneaccessOneosShowSoftFileInfo_Template = `Value STATUS (.*)
Value FILENAME (\S+)
Value FILETYPE (.+)
Value SOFTWARE_VERSION (\S+)
Value SOFTWARE_DATE (.*)
Value FILESIZE (\d+)
Value HEADER_CHECKSUM (\S+)
Value COMPUTED_CHECKSUM (\S+)
Value TARGET_DEVICE (\S+)

Start
  ^Binary\sfile
  ^\s+file\sname\s+=\s+${FILENAME}
  ^\s+file\stype\s+=\s+${FILETYPE}
  ^\s+software\sversion\s+=\s+${SOFTWARE_VERSION}
  ^\s+software\screation\sdate\s+=\s+${SOFTWARE_DATE}
  ^\s+file\ssize\s+=\s+${FILESIZE}
  ^\s+header\schecksum\s+=\s+${HEADER_CHECKSUM}
  ^\s+computed\schecksum\s+=\s+${COMPUTED_CHECKSUM}
  ^\s+target\sdevice\s+=\s+${TARGET_DEVICE}
  ^\s*---
  ^[fF]ile\sis\s${STATUS} -> Record
  ^\s*$$
  ^. -> Error

`