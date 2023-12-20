package main

import (
	"fmt"
	"encoding/json"
)

func processWebHookEvent(account string, payloadString string){

	fmt.Printf("\nPayload: %v",payloadString)
	
	var events []WebHookEvent

	err := json.Unmarshal([]byte(payloadString), &events)

	_ = err
	
	fmt.Printf("\nEvents: %v\n",events)


	/*
		Events:
			-processed 			
			 [email,timestamp,smtp-id,category,sg_event_id,sg_message_id]
			-deferred 			
			 [email,timestamp,smtp-id,category,sg_event_id,sg_message_id,response,attempt]
			-delivered 			
			 [email,timestamp,smtp-id,category,sg_event_id,sg_message_id,response]
			-open 				
			 [email,timestamp,smtp-id,category,sg_event_id,sg_message_id,useragent,ip]
			-click 				
			 [email,timestamp,smtp-id,category,sg_event_id,sg_message_id,useragent,ip,url]
			-bounce 			
			 [email,timestamp,smtp-id,category,sg_event_id,sg_message_id,reason,status]
			-dropped 			
			 [email,timestamp,smtp-id,category,sg_event_id,sg_message_id,reason,status]
			-spamreport 		
			 [email,timestamp,smtp-id,category,sg_event_id,sg_message_id]
			-unsubscribe 		
			 [email,timestamp,smtp-id,category,sg_event_id,sg_message_id]
			-group_unsubscribe 	
			 [email,timestamp,smtp-id,category,sg_event_id,sg_message_id,useragent,ip,url,asm_group_id]
			-group_resubscribe 	
			 [email,timestamp,smtp-id,category,sg_event_id,sg_message_id,useragent,ip,url,asm_group_id]
	*/

	for _, event := range events{

		eventString := ""

		switch event.Event {

			case "processed":

				tEvent := ProcessedWebHookEvent{event.Email, event.Timestamp, event.SmtpId, event.Category,event.Event, event.SgEventId, event.SgMessageId}
				eventJsonMarshall, _ := json.Marshal(tEvent)
				eventString = string(eventJsonMarshall)

			case "deferred":

				tEvent := &DeferredWebHookEvent{event.Email, event.Timestamp, event.SmtpId, event.Category,event.Event, event.SgEventId, event.SgMessageId, event.Response, event.Attempt}
				eventJsonMarshall, _ := json.Marshal(tEvent)
				eventString = string(eventJsonMarshall)

			case "delivered":

				tEvent := DeliveredWebHookEvent{event.Email,event.Timestamp,event.SmtpId,event.Category,event.Event,event.SgEventId,event.SgMessageId,event.Response}
				eventJsonMarshall, _ := json.Marshal(tEvent)
				eventString = string(eventJsonMarshall)

			case "open":

				tEvent := OpenWebHookEvent{event.Email,event.Timestamp,event.SmtpId,event.Category,event.Event,event.SgEventId,event.SgMessageId,event.Useragent,event.Ip}
				eventJsonMarshall, _ := json.Marshal(tEvent)
				eventString = string(eventJsonMarshall)

			case "click":

				tEvent := ClickWebHookEvent{event.Email,event.Timestamp,event.SmtpId,event.Category,event.Event,event.SgEventId,event.SgMessageId,event.Useragent,event.Ip,event.Url}
				eventJsonMarshall, _ := json.Marshal(tEvent)
				eventString = string(eventJsonMarshall)

			case "bounce":

				tEvent := BounceWebHookEvent{event.Email,event.Timestamp,event.SmtpId,event.Category,event.Event,event.SgEventId,event.SgMessageId,event.Reason,event.Status}
				eventJsonMarshall, _ := json.Marshal(tEvent)
				eventString = string(eventJsonMarshall)

			case "dropped":

				tEvent := DroppedWebHookEvent{event.Email,event.Timestamp,event.SmtpId,event.Category,event.Event,event.SgEventId,event.SgMessageId,event.Reason,event.Status}
				eventJsonMarshall, _ := json.Marshal(tEvent)
				eventString = string(eventJsonMarshall)

			case "spamreport":

				tEvent := SpamReportWebHookEvent{event.Email,event.Timestamp,event.SmtpId,event.Category,event.Event,event.SgEventId,event.SgMessageId}
				eventJsonMarshall, _ := json.Marshal(tEvent)
				eventString = string(eventJsonMarshall)

			case "unsubscribe":

				tEvent := UnsuscribeWebHookEvent{event.Email,event.Timestamp,event.SmtpId,event.Category,event.Event,event.SgEventId,event.SgMessageId}
				eventJsonMarshall, _ := json.Marshal(tEvent)
				eventString = string(eventJsonMarshall)

			case "group_unsubscribe":

				tEvent := GroupUnsuscribeWebHookEvent{event.Email,event.Timestamp,event.SmtpId,event.Category,event.Event,event.SgEventId,event.SgMessageId,event.Useragent,event.Ip,event.Url,event.AsmGroupId}
				eventJsonMarshall, _ := json.Marshal(tEvent)
				eventString = string(eventJsonMarshall)

			case "group_resubscribe":

				tEvent := GroupResuscribeWebHookEvent{event.Email,event.Timestamp,event.SmtpId,event.Category,event.Event,event.SgEventId,event.SgMessageId,event.Useragent,event.Ip,event.Url,event.AsmGroupId}
				eventJsonMarshall, _ := json.Marshal(tEvent)
				eventString = string(eventJsonMarshall)

			default:

				tEvent := event
				eventJsonMarshall, _ := json.Marshal(tEvent)
				eventString = string(eventJsonMarshall)

		} 

		
		
		saveWebHookFile(account,eventString,event.Timestamp,event.Event)

		if event.Event == "spamreport" {
			fmt.Printf("Sending alert for %v\n",event.Event )
			sendMessage("EVENT: "+ event.Event + " ACCOUNT: "+account+"\n" + eventString, configuration.SlackChannel)
		}
	}
	//saveIntoLog(payloadString)
	//saveWebHookEventToS3()
	//insertWebHookEvents(events)
}