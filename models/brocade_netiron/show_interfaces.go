package brocade_netiron

type ShowInterfaces struct {
	Outputcollisions             string `json:"outputcollisions"`
	Outputnptransmittedpackets   string `json:"outputnptransmittedpackets"`
	Lldpbpduforward              string `json:"lldpbpduforward"`
	Activelagrole                string `json:"activelagrole"`
	Numtaggedvlans               string `json:"numtaggedvlans"`
	Conflagprimaryport           string `json:"conflagprimaryport"`
	Monitorstate                 string `json:"monitorstate"`
	Inputbytestotal              string `json:"inputbytestotal"`
	Inputnpreceivedpackets       string `json:"inputnpreceivedpackets"`
	Stprootguard                 string `json:"stprootguard"`
	Conffiberduplex              string `json:"conffiberduplex"`
	Priority                     string `json:"priority"`
	Inputbroadcasts              string `json:"inputbroadcasts"`
	Outputunderruns              string `json:"outputunderruns"`
	Outputerrors                 string `json:"outputerrors"`
	Confcopperspeed              string `json:"confcopperspeed"`
	Stpconfstate                 string `json:"stpconfstate"`
	Inputcrcerrors               string `json:"inputcrcerrors"`
	Inputignorederrors           string `json:"inputignorederrors"`
	Interface                    string `json:"interface"`
	L2portstate                  string `json:"l2portstate"`
	Inputmulticasts              string `json:"inputmulticasts"`
	Inputerrors                  string `json:"inputerrors"`
	Outputmulticasts             string `json:"outputmulticasts"`
	Ipencapsulation              string `json:"ipencapsulation"`
	Untaggedvlan                 string `json:"untaggedvlan"`
	Inputpacketstotal            string `json:"inputpacketstotal"`
	L2portmode                   string `json:"l2portmode"`
	Priorityforcedstate          string `json:"priorityforcedstate"`
	Mirrorstate                  string `json:"mirrorstate"`
	Openflowindex                string `json:"openflowindex"`
	Loadinterval                 string `json:"loadinterval"`
	Inputpacketssec              string `json:"inputpacketssec"`
	Stpbpduguard                 string `json:"stpbpduguard"`
	Bia                          string `json:"bia"`
	Ipaddress                    string `json:"ipaddress"`
	Outputbitssec                string `json:"outputbitssec"`
	Inputgiants                  string `json:"inputgiants"`
	L2mtubytes                   string `json:"l2mtubytes"`
	Inputbitssec                 string `json:"inputbitssec"`
	Inputpercentutilization      string `json:"inputpercentutilization"`
	Outputunicasts               string `json:"outputunicasts"`
	Hardwaretype                 string `json:"hardwaretype"`
	Actualspeed                  string `json:"actualspeed"`
	Outputbroadcasts             string `json:"outputbroadcasts"`
	Mac                          string `json:"mac"`
	Inputframeerrors             string `json:"inputframeerrors"`
	Portname                     string `json:"portname"`
	Inputrunts                   string `json:"inputrunts"`
	Outputbytestotal             string `json:"outputbytestotal"`
	Protostate                   string `json:"protostate"`
	Actualduplex                 string `json:"actualduplex"`
	L3mtubytes                   string `json:"l3mtubytes"`
	Outputreceivedfromtmpackets  string `json:"outputreceivedfromtmpackets"`
	Outputpacketssec             string `json:"outputpacketssec"`
	Inputnpingressdroppedpackets string `json:"inputnpingressdroppedpackets"`
	Inputnobuffers               string `json:"inputnobuffers"`
	Inputunicasts                string `json:"inputunicasts"`
	Dhcpsnooptrust               string `json:"dhcpsnooptrust"`
	Conflagports                 string `json:"conflagports"`
	Clusterl2protoforwarding     string `json:"clusterl2protoforwarding"`
	Outputpercentutilization     string `json:"outputpercentutilization"`
	Confcopperduplex             string `json:"confcopperduplex"`
	Activelagprimaryport         string `json:"activelagprimaryport"`
	Dropprecedenceforcestate     string `json:"dropprecedenceforcestate"`
	Encapsulation                string `json:"encapsulation"`
	Inputsenttotmpackets         string `json:"inputsenttotmpackets"`
	Outputpacketstotal           string `json:"outputpacketstotal"`
	Conffiberspeed               string `json:"conffiberspeed"`
	Flowcontrolstate             string `json:"flowcontrolstate"`
	Lacpbpduforward              string `json:"lacpbpduforward"`
	Activelagports               string `json:"activelagports"`
	Conflagrole                  string `json:"conflagrole"`
	Openflowopstate              string `json:"openflowopstate"`
	Intstate                     string `json:"intstate"`
	Dropprecedencelevel          string `json:"dropprecedencelevel"`
}

var ShowInterfaces_Template string = `Value interface ([a-zA-Z0-9\/]+)
Value Required intstate (up|down|disabled)
Value protostate (up|down)
Value stprootguard (enabled|disabled)
Value stpbpduguard (enabled|disabled)
Value hardwaretype ([a-zA-Z0-9\s]+)
Value mac ([0-9a-f\.]+)
Value bia ([0-9a-f\.]+)
Value conffiberspeed ([0-9a-zA-Z]+)
Value confcopperspeed ([0-9a-zA-Z]+)
Value conffiberduplex ([0-9a-zA-Z]+)
Value confcopperduplex ([0-9a-zA-Z]+)
Value actualspeed ([0-9a-zA-Z]+)
Value actualduplex ([0-9a-zA-Z]+)
Value l2portstate (Forwarding|Disabled)
Value l2portmode (tagged|untagged|dual)
Value untaggedvlan (\d+)
Value numtaggedvlans (\d+)
Value stpconfstate (ON|OFF)
Value priority ([a-zA-Z0-9]+)
Value flowcontrolstate (enabled|disabled)
Value priorityforcedstate (enabled|disabled)
Value dropprecedencelevel (\d+)
Value dropprecedenceforcestate (enabled|disabled)
Value dhcpsnooptrust (ON|OFF)
Value mirrorstate (enabled|disabled)
Value monitorstate (enabled|disabled)
Value lacpbpduforward (Enabled|Disabled)
Value lldpbpduforward (Enabled|Disabled)
Value activelagports ([0-9\/\-\,]+)
Value activelagrole (primary|secondary)
Value activelagprimaryport ([0-9\/]+)
Value conflagports ([0-9\/\-\,]+)
Value conflagrole (primary|secondary)
Value conflagprimaryport ([0-9\/]+)
Value portname ([a-zA-Z0-9\.\ \"\(\)\#\/\-]+)
Value l2mtubytes (\d+)
Value encapsulation ([a-zA-Z0-9]+)
Value openflowopstate (Disabled|Enabled)
Value openflowindex (\d+)
Value clusterl2protoforwarding (disabled|enabled)
Value loadinterval (\d+)
Value inputbitssec (\d+)
Value inputpacketssec (\d+)
Value inputpercentutilization ([\d\.]+)
Value outputbitssec (\d+)
Value outputpacketssec (\d+)
Value outputpercentutilization ([\d\.]+)
Value inputpacketstotal (\d+)
Value inputbytestotal (\d+)
Value inputnobuffers (\d+)
Value inputbroadcasts (\d+)
Value inputmulticasts (\d+)
Value inputunicasts (\d+)
Value inputerrors (\d+)
Value inputcrcerrors (\d+)
Value inputframeerrors (\d+)
Value inputignorederrors (\d+)
Value inputrunts (\d+)
Value inputgiants (\d+)
Value inputnpreceivedpackets (\d+)
Value inputsenttotmpackets (\d+)
Value inputnpingressdroppedpackets (\d+)
Value outputpacketstotal (\d+)
Value outputbytestotal (\d+)
Value outputunderruns (\d+)
Value outputbroadcasts (\d+)
Value outputmulticasts (\d+)
Value outputunicasts (\d+)
Value outputerrors (\d+)
Value outputcollisions (\d+)
Value outputnptransmittedpackets (\d+)
Value outputreceivedfromtmpackets (\d+)
Value l3mtubytes (\d+)
Value ipaddress ([0-9\.\/]+)
Value ipencapsulation ([a-zA-Z]+)

Start
  ^\w -> Continue.Record
  ^${interface} is ${intstate}, line protocol is ${protostate}
  ^\s+Hardware is ${hardwaretype} -> Continue
  ^\s+Hardware is ${hardwaretype}, address is ${mac} \(bia ${bia}\) -> Continue
  ^\s+STP Root Guard is ${stprootguard}, STP BPDU Guard is ${stpbpduguard} -> Continue
  ^\s+Configured speed ${conffiberspeed}, actual ${actualspeed}, configured duplex ${conffiberduplex}, actual ${actualduplex} -> Continue
  ^\s+Configured fiber speed ${conffiberspeed}, configured copper speed ${confcopperspeed}, actual ${actualspeed}, configured fiber duplex ${conffiberduplex}, configured copper duplex ${confcopperduplex}, actual ${actualduplex} -> Continue
  ^\s+Member of VLAN ${untaggedvlan} \(untagged\), port is in ${l2portmode} mode, port state is ${l2portstate} -> Continue
  ^\s+Member of Control VLAN 4095, ${numtaggedvlans} L2 VLAN\(S\) \(tagged\), port is in ${l2portmode} mode, port state is ${l2portstate} -> Continue
  ^\s+Member of Control VLAN 4095, VLAN ${untaggedvlan} \(untagged\), port is in ${l2portmode} mode, port state is ${l2portstate} -> Continue
  ^\s+Member of Control VLAN 4095, VLAN ${untaggedvlan} \(untagged\), ${numtaggedvlans} L2 VLANS \(tagged\), -> Continue
  ^\s+port is in ${l2portmode} mode, port state is ${l2portstate} -> Continue
  ^\s+Type is Vlan \(Vlan Id: ${untaggedvlan}\) -> Continue
  ^\s+STP configured to ${stpconfstate}, Priority is ${priority}, flow control ${flowcontrolstate} -> Continue
  ^\s+Priority force ${priorityforcedstate}, Drop precedence level ${dropprecedencelevel}, Drop precedence force ${dropprecedenceforcestate} -> Continue
  ^\s+dhcp-snooping-trust configured to ${dhcpsnooptrust} -> Continue
  ^\s+mirror ${mirrorstate}, monitor ${monitorstate} -> Continue
  ^\s+LACP BPDU Forwarding:${lacpbpduforward} -> Continue
  ^\s+LLDP BPDU Forwarding:${lldpbpduforward} -> Continue
  ^\s+Member of active trunk ports ${activelagports}, ${activelagrole} port, primary port is ${activelagprimaryport} -> Continue
  ^\s+Member of configured trunk ports ${conflagports}, ${conflagrole} port, primary port is ${conflagprimaryport} -> Continue
  ^\s+Member of active trunk ports ${activelagports}, ${activelagrole} port -> Continue
  ^\s+Member of configured trunk ports ${conflagports}, ${conflagrole} port -> Continue
  ^\s+Port name is ${portname} -> Continue
  ^\s+MTU ${l2mtubytes} bytes, encapsulation ${encapsulation} -> Continue
  ^\s+Openflow: ${openflowopstate}, Openflow Index ${openflowindex} -> Continue
  ^\s+Cluster L2 protocol forwarding ${clusterl2protoforwarding} -> Continue
  ^\s+${loadinterval} second input rate: ${inputbitssec} bits/sec, ${inputpacketssec} packets/sec, ${inputpercentutilization}% utilization -> Continue
  ^\s+(\d+) second output rate: ${outputbitssec} bits/sec, ${outputpacketssec} packets/sec, ${outputpercentutilization}% utilization -> Continue
  ^\s+${inputpacketstotal} packets input, ${inputbytestotal} bytes, ${inputnobuffers} no buffer -> Continue
  ^\s+Received ${inputbroadcasts} broadcasts, ${inputmulticasts} multicasts, ${inputunicasts} unicasts -> Continue
  ^\s+${inputerrors} input errors, ${inputcrcerrors} CRC, ${inputframeerrors} frame, ${inputignorederrors} ignored -> Continue
  ^\s+${inputrunts} runts, ${inputgiants} giants -> Continue
  ^\s+NP received ${inputnpreceivedpackets} packets, Sent to TM ${inputsenttotmpackets} packets -> Continue
  ^\s+NP Ingress dropped ${inputnpingressdroppedpackets} packets -> Continue
  ^\s+${outputpacketstotal} packets output, ${outputbytestotal} bytes, ${outputunderruns} underruns -> Continue
  ^\s+Transmitted ${outputbroadcasts} broadcasts, ${outputmulticasts} multicasts, ${outputunicasts} unicasts -> Continue
  ^\s+${outputerrors} output errors, ${outputcollisions} collisions -> Continue
  ^\s+NP transmitted ${outputnptransmittedpackets} packets, Received from TM ${outputreceivedfromtmpackets} packets -> Continue
  ^\s+IP MTU ${l3mtubytes} bytes, encapsulation ${ipencapsulation} -> Continue
  ^\s+Internet address is ${ipaddress}, IP MTU ${l3mtubytes} bytes, encapsulation ${ipencapsulation} -> Continue

`
