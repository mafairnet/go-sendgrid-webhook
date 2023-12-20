package main

type WebHookEvent struct {
	Email         string `json:"email"`
	Timestamp      int `json:"timestamp"`
	SmtpId 	string `json:"smtp-id"`
	Category 	string `json:"category"`
	Event    string `json:"event"`
	SgEventId string `json:"sg_event_id"`
	SgMessageId string `json:"sg_message_id"`
	Response string `json:"response"`
	Reason string `json:"reason"`
	Status string `json:"status"`
	Attempt string `json:"attempt"`
	Useragent string `json:"useragent"`
	Ip string `json:"ip"`
	Url string `json:"url"`
	AsmGroupId string `json:"asm_group_id"`
}

type ProcessedWebHookEvent struct {
	Email         string `json:"email"`
	Timestamp      int `json:"timestamp"`
	SmtpId 	string `json:"smtp-id"`
	Category 	string `json:"category"`
	Event    string `json:"event"`
	SgEventId string `json:"sg_event_id"`
	SgMessageId string `json:"sg_message_id"`
}

type DeferredWebHookEvent struct {
	Email         string `json:"email"`
	Timestamp      int `json:"timestamp"`
	SmtpId 	string `json:"smtp-id"`
	Category 	string `json:"category"`
	Event    string `json:"event"`
	SgEventId string `json:"sg_event_id"`
	SgMessageId string `json:"sg_message_id"`
	Response string `json:"response"`
	Attempt string `json:"attempt"`
}

type DeliveredWebHookEvent struct {
	Email         string `json:"email"`
	Timestamp      int `json:"timestamp"`
	SmtpId 	string `json:"smtp-id"`
	Category 	string `json:"category"`
	Event    string `json:"event"`
	SgEventId string `json:"sg_event_id"`
	SgMessageId string `json:"sg_message_id"`
	Response string `json:"response"`
}

type OpenWebHookEvent struct {
	Email         string `json:"email"`
	Timestamp      int `json:"timestamp"`
	SmtpId 	string `json:"smtp-id"`
	Category 	string `json:"category"`
	Event    string `json:"event"`
	SgEventId string `json:"sg_event_id"`
	SgMessageId string `json:"sg_message_id"`
	Useragent string `json:"useragent"`
	Ip string `json:"ip"`
}

type ClickWebHookEvent struct {
	Email         string `json:"email"`
	Timestamp      int `json:"timestamp"`
	SmtpId 	string `json:"smtp-id"`
	Category 	string `json:"category"`
	Event    string `json:"event"`
	SgEventId string `json:"sg_event_id"`
	SgMessageId string `json:"sg_message_id"`
	Useragent string `json:"useragent"`
	Ip string `json:"ip"`
	Url string `json:"url"`
}

type BounceWebHookEvent struct {
	Email         string `json:"email"`
	Timestamp      int `json:"timestamp"`
	SmtpId 	string `json:"smtp-id"`
	Category 	string `json:"category"`
	Event    string `json:"event"`
	SgEventId string `json:"sg_event_id"`
	SgMessageId string `json:"sg_message_id"`
	Reason string `json:"reason"`
	Status string `json:"status"`
}

type DroppedWebHookEvent struct {
	Email         string `json:"email"`
	Timestamp      int `json:"timestamp"`
	SmtpId 	string `json:"smtp-id"`
	Category 	string `json:"category"`
	Event    string `json:"event"`
	SgEventId string `json:"sg_event_id"`
	SgMessageId string `json:"sg_message_id"`
	Reason string `json:"reason"`
	Status string `json:"status"`
}

type SpamReportWebHookEvent struct {
	Email         string `json:"email"`
	Timestamp      int `json:"timestamp"`
	SmtpId 	string `json:"smtp-id"`
	Category 	string `json:"category"`
	Event    string `json:"event"`
	SgEventId string `json:"sg_event_id"`
	SgMessageId string `json:"sg_message_id"`
}

type UnsuscribeWebHookEvent struct {
	Email         string `json:"email"`
	Timestamp      int `json:"timestamp"`
	SmtpId 	string `json:"smtp-id"`
	Category 	string `json:"category"`
	Event    string `json:"event"`
	SgEventId string `json:"sg_event_id"`
	SgMessageId string `json:"sg_message_id"`
}

type GroupUnsuscribeWebHookEvent struct {
	Email         string `json:"email"`
	Timestamp      int `json:"timestamp"`
	SmtpId 	string `json:"smtp-id"`
	Category 	string `json:"category"`
	Event    string `json:"event"`
	SgEventId string `json:"sg_event_id"`
	SgMessageId string `json:"sg_message_id"`
	Useragent string `json:"useragent"`
	Ip string `json:"ip"`
	Url string `json:"url"`
	AsmGroupId string `json:"asm_group_id"`
}

type GroupResuscribeWebHookEvent struct {
	Email         string `json:"email"`
	Timestamp      int `json:"timestamp"`
	SmtpId 	string `json:"smtp-id"`
	Category 	string `json:"category"`
	Event    string `json:"event"`
	SgEventId string `json:"sg_event_id"`
	SgMessageId string `json:"sg_message_id"`
	Useragent string `json:"useragent"`
	Ip string `json:"ip"`
	Url string `json:"url"`
	AsmGroupId string `json:"asm_group_id"`
}
