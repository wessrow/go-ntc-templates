package brocade_fastiron

type ShowInterfaces struct {
	Inputrunts               string `json:"inputrunts"`
	Queue1queued             string `json:"queue1queued"`
	Queue7queued             string `json:"queue7queued"`
	Inputunicasts            string `json:"inputunicasts"`
	Outputpacketstotal       string `json:"outputpacketstotal"`
	Outputmulticasts         string `json:"outputmulticasts"`
	Relayagentinfo           string `json:"relayagentinfo"`
	Conflagprimaryport       string `json:"conflagprimaryport"`
	Inputpercentutilization  string `json:"inputpercentutilization"`
	Inputbytestotal          string `json:"inputbytestotal"`
	Inputmulticasts          string `json:"inputmulticasts"`
	Queue1dropped            string `json:"queue1dropped"`
	Outputerrors             string `json:"outputerrors"`
	Queue2dropped            string `json:"queue2dropped"`
	Hardwaretype             string `json:"hardwaretype"`
	Bia                      string `json:"bia"`
	Flowcontrolconfigstate   string `json:"flowcontrolconfigstate"`
	Flowcontrolopstate       string `json:"flowcontrolopstate"`
	Stpbpduguard             string `json:"stpbpduguard"`
	Flowcontrolnegconfig     string `json:"flowcontrolnegconfig"`
	Mac                      string `json:"mac"`
	Priority                 string `json:"priority"`
	Inputpacketssec          string `json:"inputpacketssec"`
	Queue2queued             string `json:"queue2queued"`
	Inputcrcerrors           string `json:"inputcrcerrors"`
	Outputbroadcasts         string `json:"outputbroadcasts"`
	Queue0queued             string `json:"queue0queued"`
	Queue7dropped            string `json:"queue7dropped"`
	Queue5queued             string `json:"queue5queued"`
	Queue6dropped            string `json:"queue6dropped"`
	Stprootguard             string `json:"stprootguard"`
	Inputerrors              string `json:"inputerrors"`
	Queue0dropped            string `json:"queue0dropped"`
	Queue3queued             string `json:"queue3queued"`
	Queue6queued             string `json:"queue6queued"`
	Monitorstate             string `json:"monitorstate"`
	Activelagports           string `json:"activelagports"`
	Activelagrole            string `json:"activelagrole"`
	Outputpercentutilization string `json:"outputpercentutilization"`
	Queue4dropped            string `json:"queue4dropped"`
	Protostate               string `json:"protostate"`
	Loadinterval             string `json:"loadinterval"`
	Inputframeerrors         string `json:"inputframeerrors"`
	Queue3dropped            string `json:"queue3dropped"`
	Inputbroadcasts          string `json:"inputbroadcasts"`
	Inputgiants              string `json:"inputgiants"`
	Outputunderruns          string `json:"outputunderruns"`
	Untaggedvlan             string `json:"untaggedvlan"`
	Mirrorstate              string `json:"mirrorstate"`
	Activelagprimaryport     string `json:"activelagprimaryport"`
	Inputbitssec             string `json:"inputbitssec"`
	Conflagports             string `json:"conflagports"`
	Outputbytestotal         string `json:"outputbytestotal"`
	Intstate                 string `json:"intstate"`
	Actualspeed              string `json:"actualspeed"`
	Actualduplex             string `json:"actualduplex"`
	Conflagrole              string `json:"conflagrole"`
	Outputbitssec            string `json:"outputbitssec"`
	Inputpacketstotal        string `json:"inputpacketstotal"`
	Outputcollisions         string `json:"outputcollisions"`
	L2portmode               string `json:"l2portmode"`
	Numtaggedvlans           string `json:"numtaggedvlans"`
	Stpconfstate             string `json:"stpconfstate"`
	L2mtubytes               string `json:"l2mtubytes"`
	Interface                string `json:"interface"`
	Confspeed                string `json:"confspeed"`
	Inputignorederrors       string `json:"inputignorederrors"`
	Outputunicasts           string `json:"outputunicasts"`
	L2portstate              string `json:"l2portstate"`
	Inputnobuffers           string `json:"inputnobuffers"`
	Queue5dropped            string `json:"queue5dropped"`
	Confduplex               string `json:"confduplex"`
	Portname                 string `json:"portname"`
	Outputpacketssec         string `json:"outputpacketssec"`
	Queue4queued             string `json:"queue4queued"`
}

var ShowInterfaces_Template string = `Value interface ([a-zA-Z0-9\/\s]+)
Value Required intstate (up|down|disabled)
Value protostate (up|down)
Value stprootguard (enabled|disabled|Enabled|Disabled)
Value stpbpduguard (enabled|disabled|Enabled|Disabled)
Value hardwaretype ([a-zA-Z0-9\s]+)
Value mac ([0-9a-f\.]+)
Value bia ([0-9a-f\.]+)
Value confspeed ([0-9a-zA-Z]+)
Value confduplex ([0-9a-zA-Z]+)
Value actualspeed ([0-9a-zA-Z]+)
Value actualduplex ([0-9a-zA-Z]+)
Value l2portstate (Forwarding|Disabled|FORWARDING|DISABLED)
Value l2portmode (tagged|untagged|dual)
Value untaggedvlan (\d+)
Value numtaggedvlans (\d+)
Value stpconfstate (ON|OFF)
Value priority ([a-zA-Z0-9]+)
Value flowcontrolconfigstate (enabled|disabled)
Value flowcontrolopstate (enabled|disabled)
Value flowcontrolnegconfig (enabled|disabled)
Value mirrorstate (enabled|disabled)
Value monitorstate (enabled|disabled)
Value activelagports ([0-9\/\-\,]+)
Value activelagrole (primary|secondary)
Value activelagprimaryport ([0-9\/]+)
Value conflagports ([0-9\/\-\,]+)
Value conflagrole (primary|secondary)
Value conflagprimaryport ([0-9\/]+)
Value portname ([a-zA-Z0-9\.\ \"\(\)\#\/\-]+)
Value l2mtubytes (\d+)
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
Value outputpacketstotal (\d+)
Value outputbytestotal (\d+)
Value outputunderruns (\d+)
Value outputbroadcasts (\d+)
Value outputmulticasts (\d+)
Value outputunicasts (\d+)
Value outputerrors (\d+)
Value outputcollisions (\d+)
Value relayagentinfo (Disabled|Enabled)
Value queue0queued (\d+)
Value queue0dropped (\d+)
Value queue1queued (\d+)
Value queue1dropped (\d+)
Value queue2queued (\d+)
Value queue2dropped (\d+)
Value queue3queued (\d+)
Value queue3dropped (\d+)
Value queue4queued (\d+)
Value queue4dropped (\d+)
Value queue5queued (\d+)
Value queue5dropped (\d+)
Value queue6queued (\d+)
Value queue6dropped (\d+)
Value queue7queued (\d+)
Value queue7dropped (\d+)


Start
  ^\s+10Gig -> Continue.Record
  ^(Gig|10Gig|40Gig|\s+10Gig) -> Continue.Record
  ^\s*${interface} is ${intstate}, line protocol is ${protostate} -> Continue
  ^\s+Hardware is ${hardwaretype}, address is ${mac} \(bia ${bia}\) -> Continue
  ^\s+BPDU guard is ${stpbpduguard}, ROOT protect is ${stprootguard} -> Continue
  ^\s+Configured speed ${confspeed}, actual ${actualspeed}, configured duplex ${confduplex}, actual ${actualduplex} -> Continue
  ^\s+Member of L2 VLAN ID ${untaggedvlan}, port is ${l2portmode}, port state is ${l2portstate} -> Continue
  ^\s+Member of ${numtaggedvlans} L2 VLANs, port is ${l2portmode} mode in Vlan ${untaggedvlan}, port state is ${l2portstate} -> Continue
  ^\s+Member of ${numtaggedvlans} L2 VLANs, port is ${l2portmode}, port state is ${l2portstate} -> Continue
  ^\s+port is in ${l2portmode} mode, port state is ${l2portstate} -> Continue
  ^\s+STP configured to ${stpconfstate}, priority is ${priority} -> Continue
  ^\s+Flow Control is config ${flowcontrolconfigstate}, oper ${flowcontrolopstate}, negotiation ${flowcontrolnegconfig} -> Continue
  ^\s+Flow Control is ${flowcontrolconfigstate} -> Continue
  ^\s+Mirror ${mirrorstate}, Monitor ${monitorstate} -> Continue
  ^\s+Member of active trunk ports ${activelagports}, ${activelagrole} port, primary port is ${activelagprimaryport} -> Continue
  ^\s+Member of configured trunk ports ${conflagports}, ${conflagrole} port, primary port is ${conflagprimaryport} -> Continue
  ^\s+Member of active trunk ports ${activelagports}, ${activelagrole} port -> Continue
  ^\s+Member of configured trunk ports ${conflagports}, ${conflagrole} port -> Continue
  ^\s+Port name is ${portname} -> Continue
  ^\s+MTU ${l2mtubytes} bytes -> Continue
  ^\s+IP MTU ${l2mtubytes} bytes -> Continue
  ^\s+${loadinterval} second input rate: ${inputbitssec} bits/sec, ${inputpacketssec} packets/sec, ${inputpercentutilization}% utilization -> Continue
  ^\s+(\d+) second output rate: ${outputbitssec} bits/sec, ${outputpacketssec} packets/sec, ${outputpercentutilization}% utilization -> Continue
  ^\s+${inputpacketstotal} packets input, ${inputbytestotal} bytes, ${inputnobuffers} no buffer -> Continue
  ^\s+Received ${inputbroadcasts} broadcasts, ${inputmulticasts} multicasts, ${inputunicasts} unicasts -> Continue
  ^\s+${inputerrors} input errors, ${inputcrcerrors} CRC, ${inputframeerrors} frame, ${inputignorederrors} ignored -> Continue
  ^\s+${inputrunts} runts, ${inputgiants} giants -> Continue
  ^\s+${outputpacketstotal} packets output, ${outputbytestotal} bytes, ${outputunderruns} underruns -> Continue
  ^\s+Transmitted ${outputbroadcasts} broadcasts, ${outputmulticasts} multicasts, ${outputunicasts} unicasts -> Continue
  ^\s+${outputerrors} output errors, ${outputcollisions} collisions -> Continue
  ^\s+Relay Agent Information option: ${relayagentinfo} -> Continue
  ^\s+0\s+${queue0queued}\s+${queue0dropped} -> Continue
  ^\s+1\s+${queue1queued}\s+${queue1dropped} -> Continue
  ^\s+2\s+${queue2queued}\s+${queue2dropped} -> Continue
  ^\s+3\s+${queue3queued}\s+${queue3dropped} -> Continue
  ^\s+4\s+${queue4queued}\s+${queue4dropped} -> Continue
  ^\s+5\s+${queue5queued}\s+${queue5dropped} -> Continue
  ^\s+6\s+${queue6queued}\s+${queue6dropped} -> Continue
  ^\s+7\s+${queue7queued}\s+${queue7dropped} -> Continue
`
