import {InjectionToken} from "@angular/core";
import {IMessageService} from "../models/message-service.model";

export const MessageService = new InjectionToken<IMessageService>("MessageService");