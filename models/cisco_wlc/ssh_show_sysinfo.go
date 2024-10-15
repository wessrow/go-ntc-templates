package cisco_wlc

type SshShowSysinfo struct {
	System_up_time               string `json:"SYSTEM_UP_TIME"`
	Configured_country           string `json:"CONFIGURED_COUNTRY"`
	Number_wlans                 string `json:"NUMBER_WLANS"`
	Number_active_clients        string `json:"NUMBER_ACTIVE_CLIENTS"`
	Mac_address                  string `json:"MAC_ADDRESS"`
	Maximum_aps                  string `json:"MAXIMUM_APS"`
	System_contact               string `json:"SYSTEM_CONTACT"`
	System_name                  string `json:"SYSTEM_NAME"`
	Ip_address                   string `json:"IP_ADDRESS"`
	Firmware_version             string `json:"FIRMWARE_VERSION"`
	Bootloader_version           string `json:"BOOTLOADER_VERSION"`
	System_location              string `json:"SYSTEM_LOCATION"`
	System_timezone_location     string `json:"SYSTEM_TIMEZONE_LOCATION"`
	Product_version              string `json:"PRODUCT_VERSION"`
	Build_type                   string `json:"BUILD_TYPE"`
	Last_reset                   string `json:"LAST_RESET"`
	Field_recovery_image_version string `json:"FIELD_RECOVERY_IMAGE_VERSION"`
}

var SshShowSysinfo_Template string = `Value PRODUCT_VERSION (.*)
Value BOOTLOADER_VERSION (.*)
Value FIELD_RECOVERY_IMAGE_VERSION (.*)
Value FIRMWARE_VERSION (.*)
Value BUILD_TYPE (.*)
Value SYSTEM_NAME (.*)
Value SYSTEM_LOCATION (.*)
Value SYSTEM_CONTACT (.*)
Value IP_ADDRESS (.*)
Value LAST_RESET (.*)
Value SYSTEM_UP_TIME (.*)
Value SYSTEM_TIMEZONE_LOCATION (.*)
Value CONFIGURED_COUNTRY (.*)
Value NUMBER_WLANS (\d+)
Value NUMBER_ACTIVE_CLIENTS (\d+)
Value MAC_ADDRESS (.*)
Value MAXIMUM_APS (\d+)

Start
  ^Product Version\.+\s${PRODUCT_VERSION}
  ^Bootloader Version\.+\s${BOOTLOADER_VERSION}
  ^Field Recovery Image Version\.+\s${FIELD_RECOVERY_IMAGE_VERSION}
  ^Firmware Version\.+\s${FIRMWARE_VERSION}
  ^Build Type\.+\s${BUILD_TYPE}
  ^System Name\.+\s${SYSTEM_NAME}
  ^System Location\.+\s${SYSTEM_LOCATION}
  ^System Contact\.+\s${SYSTEM_CONTACT}
  ^IP Address\.+\s${IP_ADDRESS}
  ^Last Reset\.+\s${LAST_RESET}
  ^System Up Time\.+\s${SYSTEM_UP_TIME}
  ^System Timezone Location\.+\s${SYSTEM_TIMEZONE_LOCATION}
  ^Configured Country\.+\s${CONFIGURED_COUNTRY}
  ^Number of WLANs\.+\s${NUMBER_WLANS}
  ^Number of Active Clients\.+\s${NUMBER_ACTIVE_CLIENTS}
  ^Burned-in MAC Address\.+\s${MAC_ADDRESS}
  ^Maximum number of APs supported\.+\s${MAXIMUM_APS} -> Record
`
