//go:generate bash -c "./get_latest_templates.sh && go run ."

package main

import (
	"github.com/sirupsen/logrus"
	"github.com/wessrow/go-ntc-templates/models/cisco_ios"
	"github.com/wessrow/go-ntc-templates/parse"
)

func main() {
	// generate.GenerateFSMStructs("")

	showVersion := `
	Cisco IOS Software, vios_l2 Software (vios_l2-ADVENTERPRISEK9-M), Version 15.2(CML_NIGHTLY_20190423)FLO_DSGS7, EARLY DEPLOYMENT DEVELOPMENT BUILD, synced to  V152_6_0_81_E
Technical Support: http://www.cisco.com/techsupport
Copyright (c) 1986-2019 by Cisco Systems, Inc.
Compiled Tue 23-Apr-19 04:48 by mmen


ROM: Bootstrap program is IOSv

SW2 uptime is 1 week, 4 days, 13 hours, 46 minutes
System returned to ROM by reload
System image file is "flash0:/vios_l2-adventerprisek9-m"
Last reload reason: Unknown reason



This product contains cryptographic features and is subject to United
States and local country laws governing import, export, transfer and
use. Delivery of Cisco cryptographic products does not imply
third-party authority to import, export, distribute or use encryption.
Importers, exporters, distributors and users are responsible for
compliance with U.S. and local country laws. By using this product you
agree to comply with applicable laws and regulations. If you are unable
to comply with U.S. and local laws, return this product immediately.

A summary of U.S. laws governing Cisco cryptographic products may be found at:
http://www.cisco.com/wwl/export/crypto/tool/stqrg.html

If you require further assistance please contact us by sending email to
export@cisco.com.

Cisco IOSv () processor (revision 1.0) with 935153K/111616K bytes of memory.
Processor board ID 9694JRIOL9S
8 Gigabit Ethernet interfaces
DRAM configuration is 72 bits wide with parity disabled.
256K bytes of non-volatile configuration memory.
2097152K bytes of ATA System CompactFlash 0 (Read/Write)
0K bytes of ATA CompactFlash 1 (Read/Write)
9572K bytes of ATA CompactFlash 2 (Read/Write)
0K bytes of ATA CompactFlash 3 (Read/Write)

Configuration register is 0x101
	`

	test, _ := parse.ParseCommand[cisco_ios.ShowVersion](showVersion, cisco_ios.ShowVersion_Template)

	logrus.Info(test[0].Hostname)
}
