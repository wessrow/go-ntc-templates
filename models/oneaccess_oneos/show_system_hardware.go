package oneaccess_oneos 

type ShowSystemHardware struct {
	Device	string	`json:"DEVICE"`
	Cpu	string	`json:"CPU"`
	Core_fequency	string	`json:"CORE_FEQUENCY"`
	Ddr_frequency	string	`json:"DDR_FREQUENCY"`
	Core_complex_bus_frequency	string	`json:"CORE_COMPLEX_BUS_FREQUENCY"`
	Platform_frequency	string	`json:"PLATFORM_FREQUENCY"`
	Physical_ram_size	string	`json:"PHYSICAL_RAM_SIZE"`
	One_os_ram_size	string	`json:"ONE_OS_RAM_SIZE"`
	Nand_flash_size	string	`json:"NAND_FLASH_SIZE"`
	Ram_disk	string	`json:"RAM_DISK"`
	Flash_disk	string	`json:"FLASH_DISK"`
	Uplink	string	`json:"UPLINK"`
	Isdn	string	`json:"ISDN"`
	Radio	string	`json:"RADIO"`
	Usb	string	`json:"USB"`
	Usb0	string	`json:"USB0"`
	Usb1	string	`json:"USB1"`
	Dsp	string	`json:"DSP"`
	Wlan	string	`json:"WLAN"`
	Ports	string	`json:"PORTS"`
	Uplink_ports	string	`json:"UPLINK_PORTS"`
	Secure_boot_protection	string	`json:"SECURE_BOOT_PROTECTION"`
	Wlan0	string	`json:"WLAN0"`
	Wlan1	string	`json:"WLAN1"`
}

var ShowSystemHardware_Template = `Value DEVICE (\S+)
Value CPU (.*)
Value CORE_FEQUENCY (\d+)
Value DDR_FREQUENCY (\d+)
Value CORE_COMPLEX_BUS_FREQUENCY (\d+)
Value PLATFORM_FREQUENCY (\d+)
Value PHYSICAL_RAM_SIZE (\d+\S+)
Value ONE_OS_RAM_SIZE (\d+\S+)
Value NAND_FLASH_SIZE (\d+\S+)
Value RAM_DISK (\d+\S+)
Value FLASH_DISK (\d+\S+)
Value UPLINK ([xX]|)
Value ISDN ([xX]|)
Value RADIO ([xX]|)
Value USB ([xX]|)
Value USB0 ([xX]|)
Value USB1 ([xX]|)
Value DSP (\d+)
Value WLAN (.*)
Value PORTS (.*\s+\d+\s+ports)
Value UPLINK_PORTS (.*\S)
Value SECURE_BOOT_PROTECTION (\S+)
Value WLAN0 (.*\S)
Value WLAN1 (.*\S)

Start
  ^\s*HARDWARE DESCRIPTION\s*$$
  ^\s*Device\s*:\s*${DEVICE}\s*$$
  ^\s*CPU\s*:\s*${CPU}\s*$$
  ^\s*Core\sFreq\s*:\s*${CORE_FEQUENCY}MHz\s+DDR\sFreq\s*:\s*${DDR_FREQUENCY}MHz\s*
  ^\s*Core Complex Bus Freq\s*:\s*${CORE_COMPLEX_BUS_FREQUENCY}MHz\s*Platform Freq\s*:\s*${PLATFORM_FREQUENCY}MHz\s*$$
  ^\s*Physical\sRam\ssize\s*:\s*${PHYSICAL_RAM_SIZE}\s*OneOS Ram size\s*:\s*${ONE_OS_RAM_SIZE}\s*$$
  ^\s*Physical\sRam\ssize\s*:\s*${PHYSICAL_RAM_SIZE}\s*$$
  ^\s*Nand Flash size\s*:\s*${NAND_FLASH_SIZE}\s*$$
  ^\s*Ram disk\s*:\s*${RAM_DISK}\s*Flash disk\s*:\s*${FLASH_DISK}\s*$$
  ^\s*Local\s*:\s*.*\s+Uplink\s+:\s+${UPLINK} -> Continue
  ^\s*Local\s*:\s*.*\s+ISDN\s+:\s+${ISDN} -> Continue
  ^\s*Local\s*:\s*.*\s+Radio\s+:\s+${RADIO} -> Continue
  ^\s*Local\s*:\s*.*\s+Usb0\s+:\s+${USB0} -> Continue
  ^\s*Local\s*:\s*.*\s+Usb1\s+:\s+${USB1} -> Continue
  ^\s*Local\s*:\s*.*\s+Usb\s+:\s+${USB} -> Continue
  ^\s*Secure\sBoot\sprotection\s*:\s*${SECURE_BOOT_PROTECTION}
  ^\s*Dsp\s*:\s*${DSP}\s*$$
  ^\s*Wlan\s*:\s*${WLAN}\s*$$
  ^\s*Local\s+:\s+${PORTS}\s*
  ^\s*Uplink\s+:\s+${UPLINK_PORTS}\s
  ^\s*Wlan\s+0\s+:\s+${WLAN0}\s
  ^\s*Wlan\s+1\s+:\s+${WLAN1}\s
  ^\s*
  ^. -> Error
`