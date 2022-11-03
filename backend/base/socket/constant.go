package socket

import LOG "backend/base/log"

type InstanceStatus int8

var (
	InstanceOnline    InstanceStatus = 1
	InstanceUnHealth  InstanceStatus = 2
	InstanceUnderLine InstanceStatus = 3
)

type InstanceModel int8

var (
	SeverInstance  InstanceModel = 1
	ClientInstance InstanceModel = 2
)

var log = LOG.GetLog()
