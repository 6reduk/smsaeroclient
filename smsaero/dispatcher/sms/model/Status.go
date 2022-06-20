package smsModel

type Status string

type StatusEnum map[string]Status

const (
	inQueueSmsStatus      = "IN_QUEUE"
	deliveredSmsStatus    = "DELIVERED"
	notDeliveredSmsStatus = "NOT_DELIVERED"
	transmittedSmsStatus  = "TRANSMITTED"
	onModerationSmsStatus = "ON_MODERATION"
	declinedSmsStatus     = "DECLINED"
	pendingSmsStatus      = "PENDING_FOR_STATUS"
)

var statusIDToHumanReadable = map[int]Status{
	0: inQueueSmsStatus,
	1: deliveredSmsStatus,
	2: notDeliveredSmsStatus,
	3: transmittedSmsStatus,
	4: pendingSmsStatus,
	6: declinedSmsStatus,
	8: onModerationSmsStatus,
}

var SmsStatuses = StatusEnum{
	inQueueSmsStatus:      inQueueSmsStatus,
	deliveredSmsStatus:    deliveredSmsStatus,
	notDeliveredSmsStatus: notDeliveredSmsStatus,
	transmittedSmsStatus:  transmittedSmsStatus,
	pendingSmsStatus:      pendingSmsStatus,
	declinedSmsStatus:     declinedSmsStatus,
	onModerationSmsStatus: onModerationSmsStatus,
}

func (e StatusEnum) InQueue() Status {
	return e[inQueueSmsStatus]
}

func (e StatusEnum) Delivered() Status {
	return e[deliveredSmsStatus]
}

func (e StatusEnum) NotDelivered() Status {
	return e[notDeliveredSmsStatus]
}

func (e StatusEnum) Transmitted() Status {
	return e[transmittedSmsStatus]
}

func (e StatusEnum) OnModeration() Status {
	return e[onModerationSmsStatus]
}

func (e StatusEnum) Declined() Status {
	return e[declinedSmsStatus]
}

func (e StatusEnum) Pending() Status {
	return e[pendingSmsStatus]
}

func (e StatusEnum) Of(statusId int) (Status, error) {
	status, ok := statusIDToHumanReadable[statusId]
	if !ok {
		return "", ErrUnknownMessageStatusID
	}

	return status, nil
}
