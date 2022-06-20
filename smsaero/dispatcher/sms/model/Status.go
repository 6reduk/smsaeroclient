package smsModel

type ProcessingStatus string

type ProcessingStatusEnum map[string]ProcessingStatus

const (
	inQueueSmsStatus      = "IN_QUEUE"
	deliveredSmsStatus    = "DELIVERED"
	notDeliveredSmsStatus = "NOT_DELIVERED"
	transmittedSmsStatus  = "TRANSMITTED"
	onModerationSmsStatus = "ON_MODERATION"
	declinedSmsStatus     = "DECLINED"
	pendingSmsStatus      = "PENDING_FOR_STATUS"
)

var statusIDToHumanReadable = map[int]ProcessingStatus{
	0: inQueueSmsStatus,
	1: deliveredSmsStatus,
	2: notDeliveredSmsStatus,
	3: transmittedSmsStatus,
	4: pendingSmsStatus,
	6: declinedSmsStatus,
	8: onModerationSmsStatus,
}

var SmsProcessingStatuses = ProcessingStatusEnum{
	inQueueSmsStatus:      inQueueSmsStatus,
	deliveredSmsStatus:    deliveredSmsStatus,
	notDeliveredSmsStatus: notDeliveredSmsStatus,
	transmittedSmsStatus:  transmittedSmsStatus,
	pendingSmsStatus:      pendingSmsStatus,
	declinedSmsStatus:     declinedSmsStatus,
	onModerationSmsStatus: onModerationSmsStatus,
}

func (e ProcessingStatusEnum) InQueue() ProcessingStatus {
	return e[inQueueSmsStatus]
}

func (e ProcessingStatusEnum) Delivered() ProcessingStatus {
	return e[deliveredSmsStatus]
}

func (e ProcessingStatusEnum) NotDelivered() ProcessingStatus {
	return e[notDeliveredSmsStatus]
}

func (e ProcessingStatusEnum) Transmitted() ProcessingStatus {
	return e[transmittedSmsStatus]
}

func (e ProcessingStatusEnum) OnModeration() ProcessingStatus {
	return e[onModerationSmsStatus]
}

func (e ProcessingStatusEnum) Declined() ProcessingStatus {
	return e[declinedSmsStatus]
}

func (e ProcessingStatusEnum) Pending() ProcessingStatus {
	return e[pendingSmsStatus]
}

func (e ProcessingStatusEnum) Of(statusId int) (ProcessingStatus, error) {
	status, ok := statusIDToHumanReadable[statusId]
	if !ok {
		return "", ErrUnknownMessageStatusID
	}

	return status, nil
}
