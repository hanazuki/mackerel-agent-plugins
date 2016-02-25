package main

import (
	mp "github.com/mackerelio/go-mackerel-plugin-helper"
)

var mapping = map[string]string{
	"TcpExtSyncookiesSent":            "tcp.syncookies.sent",
	"TcpExtSyncookiesRecv":            "tcp.syncookies.recv",
	"TcpExtSyncookiesFailed":          "tcp.syncookies.failed",
	"TcpExtEmbryonicRsts": "",
	"TcpExtPruneCalled": "",
	"TcpExtRcvPruned": "",
	"TcpExtOfoPruned": "",
	"TcpExtOutOfWindowIcmps": "",
	"TcpExtLockDroppedIcmps": "",
	"TcpExtArpFilter": "",
	"TcpExtTW": "tcp.timewait.timewait",
	"TcpExtTWRecycled": "tcp.timewait.recycled",
	"TcpExtTWKilled": "tcp.timewait.killed",
	"TcpExtPAWSPassive": "",
	"TcpExtPAWSActive": "",
	"TcpExtPAWSEstab": "",
	"TcpExtDelayedACKs": "",
	"TcpExtDelayedACKLocked": "",
	"TcpExtDelayedACKLost": "",
	"TcpExtListenOverflows": "",
	"TcpExtListenDrops": "",
	"TcpExtTCPPrequeued": "",
	"TcpExtTCPDirectCopyFromBacklog": "",
	"TcpExtTCPDirectCopyFromPrequeue": "",
	"TcpExtTCPPrequeueDropped": "",
	"TcpExtTCPHPHits": "",
	"TcpExtTCPHPHitsToUser": "",
	"TcpExtTCPPureAcks": "",
	"TcpExtTCPHPAcks": "",
	"TcpExtTCPRenoRecovery": "tcp.tcpreno.recovery",
	"TcpExtTCPSackRecovery": "tcp.tcpsack.recovery",
	"TcpExtTCPSACKReneging": "tcp.tcpsack.reneging",
	"TcpExtTCPFACKReorder": "tcp.tcpfack.reorder",
	"TcpExtTCPSACKReorder": "tcp.tcpsack.reorder",
	"TcpExtTCPRenoReorder": "tcp.tcpreno.reorder",
	"TcpExtTCPTSReorder": "tcp.tcpts.reorder",
	"TcpExtTCPFullUndo": "",
	"TcpExtTCPPartialUndo": "",
	"TcpExtTCPDSACKUndo": "",
	"TcpExtTCPLossUndo": "",
	"TcpExtTCPLoss": "",
	"TcpExtTCPLostRetransmit": "",
	"TcpExtTCPRenoFailures": "tcp.tcpreno.failures",
	"TcpExtTCPSackFailures": "tcp.tcpreno.failures",
	"TcpExtTCPLossFailures": "",
	"TcpExtTCPFastRetrans": "",
	"TcpExtTCPForwardRetrans": "",
	"TcpExtTCPSlowStartRetrans": "",
	"TcpExtTCPTimeouts": "",
	"TcpExtTCPRenoRecoveryFail": "",
	"TcpExtTCPSackRecoveryFail": "",
	"TcpExtTCPSchedulerFailed": "",
	"TcpExtTCPRcvCollapsed": "",
	"TcpExtTCPDSACKOldSent": "",
	"TcpExtTCPDSACKOfoSent": "",
	"TcpExtTCPDSACKRecv": "",
	"TcpExtTCPDSACKOfoRecv": "",
	"TcpExtTCPAbortOnData": "tcp.abort.on_data",
	"TcpExtTCPAbortOnClose": "tcp.abort.on_close",
	"TcpExtTCPAbortOnMemory": "tcp.abort.on_memory",
	"TcpExtTCPAbortOnTimeout": "tcp.abort.on_timeout",
	"TcpExtTCPAbortOnLinger": "tcp.abort.on_linger",
	"TcpExtTCPAbortFailed": "tcp.abort.failed",
	"TcpExtTCPMemoryPressures": "",
	"TcpExtTCPSACKDiscard": "",
	"TcpExtTCPDSACKIgnoredOld": "",
	"TcpExtTCPDSACKIgnoredNoUndo": "",
	"TcpExtTCPSpuriousRTOs": "",
	"TcpExtTCPMD5NotFound": "",
	"TcpExtTCPMD5Unexpected": "",
	"TcpExtTCPSackShifted": "",
	"TcpExtTCPSackMerged": "",
	"TcpExtTCPSackShiftFallback": "",
	"TcpExtTCPBacklogDrop": "",
	"TcpExtTCPMinTTLDrop": "",
	"TcpExtTCPChallengeACK": "",
	"TcpExtTCPSYNChallenge": "",
	"IpExtInNoRoutes": "",
	"IpExtInTruncatedPkts": "",
	"IpExtInMcastPkts": "",
	"IpExtOutMcastPkts": "",
	"IpExtInBcastPkts": "",
	"IpExtOutBcastPkts": "",
	"IpExtInOctets": "",
	"IpExtOutOctets": "",
	"IpExtInMcastOctets": "",
	"IpExtOutMcastOctets": "",
	"IpExtInBcastOctets": "",
	"IpExtOutBcastOctets": "",
	"IpForwarding": "",
	"IpDefaultTTL": "",
	"IpInReceives": "",
	"IpInHdrErrors": "",
	"IpInAddrErrors": "",
	"IpForwDatagrams": "",
	"IpInUnknownProtos": "",
	"IpInDiscards": "",
	"IpInDelivers": "",
	"IpOutRequests": "",
	"IpOutDiscards": "",
	"IpOutNoRoutes": "",
	"IpReasmTimeout": "ip.reasm.timeout",
	"IpReasmReqds": "ip.reams.reqds",
	"IpReasmOKs": "ip.reasm.oks",
	"IpReasmFails": "ip.reasm.fails",
	"IpFragOKs": "ip.frag.oks",
	"IpFragFails": "ip.frag.fails",
	"IpFragCreates": "",
	"IcmpInMsgs": "",
	"IcmpInErrors": "",
	"IcmpInDestUnreachs": "",
	"IcmpInTimeExcds": "",
	"IcmpInParmProbs": "",
	"IcmpInSrcQuenchs": "",
	"IcmpInRedirects": "",
	"IcmpInEchos": "",
	"IcmpInEchoReps": "",
	"IcmpInTimestamps": "",
	"IcmpInTimestampReps": "",
	"IcmpInAddrMasks": "",
	"IcmpInAddrMaskReps": "",
	"IcmpOutMsgs": "",
	"IcmpOutErrors": "",
	"IcmpOutDestUnreachs": "",
	"IcmpOutTimeExcds": "",
	"IcmpOutParmProbs": "",
	"IcmpOutSrcQuenchs": "",
	"IcmpOutRedirects": "",
	"IcmpOutEchos": "",
	"IcmpOutEchoReps": "",
	"IcmpOutTimestamps": "",
	"IcmpOutTimestampReps": "",
	"IcmpOutAddrMasks": "",
	"IcmpOutAddrMaskReps": "",
	"IcmpMsgInType0": "",
	"IcmpMsgInType3": "",
	"IcmpMsgInType8": "",
	"IcmpMsgInType11": "",
	"IcmpMsgOutType0": "",
	"IcmpMsgOutType3": "",
	"TcpRtoAlgorithm": "",
	"TcpRtoMin": "",
	"TcpRtoMax": "",
	"TcpMaxConn": "",
	"TcpActiveOpens": "tcp.sockets.active_opens",
	"TcpPassiveOpens": "tcp.sockets.passive_opens",
	"TcpAttemptFails": "tcp.sockets.attempt_fails",
	"TcpEstabResets": "tcp.sockets.estab_resets",
	"TcpCurrEstab": "",
	"TcpInSegs": "tcp.segments.in_segs",
	"TcpOutSegs": "tcp.segments.out_segs",
	"TcpRetransSegs": "tcp.segments.retrans_segs",
	"TcpInErrs": "tcp.segments.in_errs",
	"TcpOutRsts": "",
	"UdpInDatagrams": "udp.databgrams.in_datagrams",
	"UdpNoPorts": "upd.datagrams.no_ports",
	"UdpInErrors": "udp.datagrams.in_errors",
	"UdpOutDatagrams": "udp.datagrams.out_datagrams",
	"UdpRcvbufErrors": "",
	"UdpSndbufErrors": "",
	"UdpLiteInDatagrams": "",
	"UdpLiteNoPorts": "",
	"UdpLiteInErrors": "",
	"UdpLiteOutDatagrams": "",
	"UdpLiteRcvbufErrors": "",
	"UdpLiteSndbufErrors": "",
}

func (p *pluginOpts) GraphDefinition() map[string]mp.Graphs {
	return map[string]mp.Graphs{
		p.prefix + "tcp.syncookies": {
			Label: "TCP SYN cookies",
			Unit:  "integer",
			Metrics: []mp.Metrics{
				{Name: "*", Label: '%1', Diff: true},
			},
		},
		p.prefix + "tcp.sockets": {
			Label: "TCP Sockets",
			Unit: "integer",
			Metrics: []mp.Metrics{
				{Name: "*", Diff: true},
			},
		},
		p.prefix + "tcp.segments": {
			Label: "TCP Segments",
			Unit: "integer",
			Metrics: []mp.Metrics{
				{Name: "*", Diff: true},
			},
		},
		p.prefix + "udp.datagrams": {
			Label: "UDP datagrams",
			Unit: "integer",
			Metrics: []mp.Metrics{
				{Name: "*", Diff: true},
			},
		},
	}
}
