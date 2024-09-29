package juniper_junos 

type ShowChassisHardware struct {
	Chassis_serial_number	string	`json:"CHASSIS_SERIAL_NUMBER"`
	Chassis_description	string	`json:"CHASSIS_DESCRIPTION"`
	Mid_version	string	`json:"MID_VERSION"`
	Mid_part	string	`json:"MID_PART"`
	Mid_serial_number	string	`json:"MID_SERIAL_NUMBER"`
	Mid_description	string	`json:"MID_DESCRIPTION"`
	Fpm_version	string	`json:"FPM_VERSION"`
	Fpm_part	string	`json:"FPM_PART"`
	Fpm_serial_number	string	`json:"FPM_SERIAL_NUMBER"`
	Fpm_description	string	`json:"FPM_DESCRIPTION"`
	Ps_number	string	`json:"PS_NUMBER"`
	Ps_version	string	`json:"PS_VERSION"`
	Ps_part	string	`json:"PS_PART"`
	Ps_serial_number	string	`json:"PS_SERIAL_NUMBER"`
	Ps_description	string	`json:"PS_DESCRIPTION"`
	Re_number	string	`json:"RE_NUMBER"`
	Re_version	string	`json:"RE_VERSION"`
	Re_part	string	`json:"RE_PART"`
	Re_serial_number	string	`json:"RE_SERIAL_NUMBER"`
	Re_description	string	`json:"RE_DESCRIPTION"`
	Tfeb_number	string	`json:"TFEB_NUMBER"`
	Tfeb_version	string	`json:"TFEB_VERSION"`
	Tfeb_part	string	`json:"TFEB_PART"`
	Tfeb_serial_number	string	`json:"TFEB_SERIAL_NUMBER"`
	Tfeb_description	string	`json:"TFEB_DESCRIPTION"`
	Qxm_number	string	`json:"QXM_NUMBER"`
	Qxm_version	string	`json:"QXM_VERSION"`
	Qxm_part	string	`json:"QXM_PART"`
	Qxm_serial_number	string	`json:"QXM_SERIAL_NUMBER"`
	Qxm_description	string	`json:"QXM_DESCRIPTION"`
	Afeb_number	string	`json:"AFEB_NUMBER"`
	Afeb_version	string	`json:"AFEB_VERSION"`
	Afeb_part	string	`json:"AFEB_PART"`
	Afeb_serial_number	string	`json:"AFEB_SERIAL_NUMBER"`
	Afeb_description	string	`json:"AFEB_DESCRIPTION"`
	Cb_number	string	`json:"CB_NUMBER"`
	Cb_version	string	`json:"CB_VERSION"`
	Cb_part	string	`json:"CB_PART"`
	Cb_serial_number	string	`json:"CB_SERIAL_NUMBER"`
	Cb_description	string	`json:"CB_DESCRIPTION"`
	Fpc_number	string	`json:"FPC_NUMBER"`
	Fpc_version	string	`json:"FPC_VERSION"`
	Fpc_part	string	`json:"FPC_PART"`
	Fpc_serial_number	string	`json:"FPC_SERIAL_NUMBER"`
	Fpc_description	string	`json:"FPC_DESCRIPTION"`
	Cpu_version	string	`json:"CPU_VERSION"`
	Cpu_part	string	`json:"CPU_PART"`
	Cpu_serial_number	string	`json:"CPU_SERIAL_NUMBER"`
	Cpu_description	string	`json:"CPU_DESCRIPTION"`
	Mic_number	string	`json:"MIC_NUMBER"`
	Mic_version	string	`json:"MIC_VERSION"`
	Mic_part	string	`json:"MIC_PART"`
	Mic_serial_number	string	`json:"MIC_SERIAL_NUMBER"`
	Mic_description	string	`json:"MIC_DESCRIPTION"`
	Pic_number	string	`json:"PIC_NUMBER"`
	Pic_version	string	`json:"PIC_VERSION"`
	Pic_part	string	`json:"PIC_PART"`
	Pic_serial_number	string	`json:"PIC_SERIAL_NUMBER"`
	Pic_description	string	`json:"PIC_DESCRIPTION"`
	Xcvr_number	string	`json:"XCVR_NUMBER"`
	Xcvr_version	string	`json:"XCVR_VERSION"`
	Xcvr_part	string	`json:"XCVR_PART"`
	Xcvr_serial_number	string	`json:"XCVR_SERIAL_NUMBER"`
	Xcvr_description	string	`json:"XCVR_DESCRIPTION"`
	Fan_number	string	`json:"FAN_NUMBER"`
	Fan_version	string	`json:"FAN_VERSION"`
	Fan_part	string	`json:"FAN_PART"`
	Fan_serial_number	string	`json:"FAN_SERIAL_NUMBER"`
	Fan_description	string	`json:"FAN_DESCRIPTION"`
}

var ShowChassisHardware_Template = `Value CHASSIS_SERIAL_NUMBER (\w+)
Value CHASSIS_DESCRIPTION (.*)
Value MID_VERSION ([RrEeVv]+\s+\d+|]|[RrEeVv]+|)
Value MID_PART (\S+)
Value MID_SERIAL_NUMBER (\w+)
Value MID_DESCRIPTION (.*)
Value FPM_VERSION ([RrEeVv]+\s+\d+|]|[RrEeVv]+|)
Value FPM_PART (\S+)
Value FPM_SERIAL_NUMBER (\w+)
Value FPM_DESCRIPTION (.*)
Value PS_NUMBER (\d+|)
Value PS_VERSION ([RrEeVv]+\s+\d+|]|[RrEeVv]+|)
Value PS_PART (\S+)
Value PS_SERIAL_NUMBER (\w+)
Value PS_DESCRIPTION (.*)
Value RE_NUMBER (\d+|)
Value RE_VERSION ([RrEeVv]+\s+\d+|]|[RrEeVv]+|)
Value RE_PART (\S+)
Value RE_SERIAL_NUMBER (\w+)
Value RE_DESCRIPTION (.*)
Value TFEB_NUMBER (\d+|)
Value TFEB_VERSION ([RrEeVv]+\s+\d+|]|[RrEeVv]+|)
Value TFEB_PART (\S+)
Value TFEB_SERIAL_NUMBER (\w+)
Value TFEB_DESCRIPTION (.*)
Value QXM_NUMBER (\d+|)
Value QXM_VERSION ([RrEeVv]+\s+\d+|]|[RrEeVv]+|)
Value QXM_PART (\S+)
Value QXM_SERIAL_NUMBER (\w+)
Value QXM_DESCRIPTION (.*)
Value AFEB_NUMBER (\d+|)
Value AFEB_VERSION ([RrEeVv]+\s+\d+|]|[RrEeVv]+|)
Value AFEB_PART (\S+)
Value AFEB_SERIAL_NUMBER (\w+)
Value AFEB_DESCRIPTION (.*)
Value CB_NUMBER (\d+|)
Value CB_VERSION ([RrEeVv]+\s+\d+|]|[RrEeVv]+|)
Value CB_PART (\S+)
Value CB_SERIAL_NUMBER (\w+)
Value CB_DESCRIPTION (.*)
Value Filldown FPC_NUMBER (\d+|)
Value Filldown FPC_VERSION ([RrEeVv]+\s+\d+|]|[RrEeVv]+|)
Value Filldown FPC_PART (\S+)
Value Filldown FPC_SERIAL_NUMBER (\w+)
Value Filldown FPC_DESCRIPTION (.*)
Value Filldown CPU_VERSION ([RrEeVv]+\s+\d+|]|[RrEeVv]+|)
Value Filldown CPU_PART (\S+)
Value Filldown CPU_SERIAL_NUMBER (\w+)
Value Filldown CPU_DESCRIPTION (.*)
Value Filldown MIC_NUMBER (\d+|)
Value Filldown MIC_VERSION ([RrEeVv]+\s+\d+|]|[RrEeVv]+|)
Value Filldown MIC_PART (\S+)
Value Filldown MIC_SERIAL_NUMBER (\w+)
Value Filldown MIC_DESCRIPTION (.*)
Value Filldown PIC_NUMBER (\d+|)
Value Filldown PIC_VERSION ([RrEeVv]+\s+\d+|]|[RrEeVv]+|)
Value Filldown PIC_PART (\S+)
Value Filldown PIC_SERIAL_NUMBER (\w+)
Value Filldown PIC_DESCRIPTION (.*)
Value XCVR_NUMBER (\d+|)
Value XCVR_VERSION ([RrEeVv]+\s+\d+|]|[RrEeVv]+|)
Value XCVR_PART (\S+)
Value XCVR_SERIAL_NUMBER (\S+)
Value XCVR_DESCRIPTION (.*)
Value FAN_NUMBER (\d+|)
Value FAN_VERSION ([RrEeVv]+\s+\d+|)
Value FAN_PART (\w+-\w+|BUILTIN)
Value FAN_SERIAL_NUMBER (\w+)
Value FAN_DESCRIPTION (.*)

Start
  ^Hardware.*
  ^Item\s+Version\s+Part\s+number\s+Serial\s+number\s+Description\s*$$
  ^Chassis.*\s+${CHASSIS_SERIAL_NUMBER}\s+${CHASSIS_DESCRIPTION}$$ -> Record
  ^\s*Midplane\s+${MID_VERSION}\s+${MID_PART}\s+${MID_SERIAL_NUMBER}\s+${MID_DESCRIPTION}$$ -> Record
  ^\s*FPM\sBoard\s+${FPM_VERSION}\s+${FPM_PART}\s+${FPM_SERIAL_NUMBER}\s+${FPM_DESCRIPTION}$$ -> Record
  ^\s*Routing\s+Engine\s+${RE_NUMBER}\s+${RE_VERSION}\s+${RE_PART}\s+${RE_SERIAL_NUMBER}\s+${RE_DESCRIPTION}$$ -> Record
  ^\s*TFEB\s+${TFEB_NUMBER}\s+${TFEB_VERSION}\s+${TFEB_PART}\s+${TFEB_SERIAL_NUMBER}\s+${TFEB_DESCRIPTION}$$ -> Continue
  ^\s*QXM\s+${QXM_NUMBER}\s+${QXM_VERSION}\s+${QXM_PART}\s+${QXM_SERIAL_NUMBER}\s+${QXM_DESCRIPTION}$$ -> Record
  ^\s*AFEB\s+${AFEB_NUMBER}\s+${AFEB_VERSION}\s+${AFEB_PART}\s+${AFEB_SERIAL_NUMBER}\s+${AFEB_DESCRIPTION}$$ -> Record
  ^\s*CB\s+${CB_NUMBER}\s+${CB_VERSION}\s+${CB_PART}\s+${CB_SERIAL_NUMBER}\s+${CB_DESCRIPTION}$$ -> Record
  ^\s*FPC\s+${FPC_NUMBER}\s+${FPC_VERSION}\s+${FPC_PART}\s+${FPC_SERIAL_NUMBER}\s+${FPC_DESCRIPTION}$$
  ^\s*CPU\s+${CPU_VERSION}\s+${CPU_PART}\s+${CPU_SERIAL_NUMBER}\s+${CPU_DESCRIPTION}$$
  ^\s*MIC\s+${MIC_NUMBER}\s+${MIC_VERSION}\s+${MIC_PART}\s+${MIC_SERIAL_NUMBER}\s+${MIC_DESCRIPTION}$$
  ^\s*PIC\s+${PIC_NUMBER}\s+${PIC_VERSION}\s+${PIC_PART}\s+${PIC_SERIAL_NUMBER}\s+${PIC_DESCRIPTION}$$
  ^\s*Xcvr\s+${XCVR_NUMBER}\s+(?:${XCVR_VERSION}\s+)?${XCVR_PART}\s+${XCVR_SERIAL_NUMBER}\s+${XCVR_DESCRIPTION} -> Record
  ^\s*(PEM|Power\s+Supply) -> Continue.Clearall
  ^\s*(PEM|Power\s+Supply)\s+${PS_NUMBER}\s+${PS_VERSION}\s+${PS_PART}\s+${PS_SERIAL_NUMBER}\s+${PS_DESCRIPTION}$$ -> Record
  ^\s*(Fan|Fan\sTray) -> Continue.Clearall
  ^\s*(Fan|Fan\sTray)\s+${FAN_NUMBER}\s+${FAN_VERSION}\s+${FAN_PART}\s+${FAN_SERIAL_NUMBER}\s+${FAN_DESCRIPTION}$$ -> Record
  ^\s*(Fan|Fan\sTray)\s+${FAN_NUMBER}\s+${FAN_DESCRIPTION}$$ -> Record
  ^\s*(Pseudo|AFEB|TFEB|CPU|QXM|Fan\s+Tray).*
  ^{master:\d+}
  ^\s*$$ 
  ^. -> Error

EOF
`