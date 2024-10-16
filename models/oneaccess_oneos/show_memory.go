package oneaccess_oneos

type ShowMemory struct {
	Ram                  string `json:"RAM"`
	Boot_size            string `json:"BOOT_SIZE"`
	Flash_size           string `json:"FLASH_SIZE"`
	User_partition_total string `json:"USER_PARTITION_TOTAL"`
	User_partition_used  string `json:"USER_PARTITION_USED"`
	User_partition_free  string `json:"USER_PARTITION_FREE"`
}

var ShowMemory_Template string = `Value RAM (\S+|\d+\s?\d+)
Value BOOT_SIZE (\S+|\d+\s?\d+)
Value FLASH_SIZE (\S+|\d+\s?\d+)
Value USER_PARTITION_TOTAL (\S+|\d+\s?\d+)
Value USER_PARTITION_USED (\S+|\d+\s?\d+)
Value USER_PARTITION_FREE (\S+|\d+\s?\d+)

			   
Start
  # total system memory
  ^\s*\|\sRam\ssize\s+\|\s+${RAM}\s+\|
  ^\s*\|\sMemory Total\s+\|\s+${RAM}\s+\|
  # total boot size
  ^\s*\|\s+:\.\.Boot\s+\|\s+${BOOT_SIZE}\s+\|
  ^\s*\|\sBoot\sPartition\s+\|\s+${BOOT_SIZE}\s+\|
  # total flash disk size
  ^\s*\|\s+Extended\sFlash\ssize\s+\|\s+${FLASH_SIZE}\s+\|
  ^\s*\|\sFlash\sTotal\s+\|\s+${FLASH_SIZE}\s+\|
  # total user partition size and usage, part of flash (/BSA)
  ^\s*\|\s+:\.\.Flash\sdisk\stotal\s+\|\s+${USER_PARTITION_TOTAL}\s+\|
  ^\s*\|\s+\-\suser\s+\|\s+${USER_PARTITION_TOTAL}\s+\|\s+${USER_PARTITION_FREE}\s+\|\s+${USER_PARTITION_USED}\s+\|
  ^\s*\|\s+used\s+\|\s+${USER_PARTITION_USED}\s+\|
  ^\s*\|\s+free\s+\|\s+${USER_PARTITION_FREE}\s+\|
  ^\s*=+\s*$$
  ^\s*\|\s*Memory\s+status\s+report\s+\|\s+Total\s+\|\s+Free\s+\|\s+Use\s*%\s+\|\s*$$
  ^\s*\|\s*Memory\s+status\s+report\s+\|\s+Kbytes\s+\|\s+\|\s*$$
  ^\s*\|\-+\|
  ^\s*\|\s*(Shared|Control|Forwarding)\s+Partition\s+\|
  ^\s*\|\s+\-\s*((Shared|Core\s+0)\s+Pool|Binary)\s+\|
  ^\s*\|\s+\|
  ^\s*\|\s+\-\s*(Linux\s+(RAM|cached|buffers|File\s+Systems)|system\s+free)
  ^\s*\|\s+\-\s*(root|tmp)\s+\|
  ^\s*\|\s*File\s+systems\s+\|
  ^\s*\|\s*Removable\s+disks\s+\|
  ^\s*\|\s*:(\s+:)?\.\.
  ^\s*\|\s*:(\s+:)?\s+(used|free)\s+\|
  ^\s*\|\s*Flash\s+size\s+\|
  ^\s*$$
  ^. -> Error
`
