package socket

type ProcessorStatus int8

var (
	Init        ProcessorStatus = 1
	UnHealth    ProcessorStatus = 2
	Established ProcessorStatus = 3
	Sending     ProcessorStatus = 5
	Closing     ProcessorStatus = 6
	Closed      ProcessorStatus = 7
)
