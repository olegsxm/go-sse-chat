import {Injectable, NgZone} from '@angular/core';
import {IMessageService} from "../models/message-service.model";
import {Observable, Subscriber} from "rxjs";
import {environment} from "../../../environments/environment";
import {IMessage} from "../models/message.model";

export enum MessageType {
    message = "message",
    error = "error"
}

@Injectable()
export class SseService implements IMessageService {
    private eventSource!: EventSource;
    private eventNames: string[] = [MessageType.message];

    constructor(private zone: NgZone) {
    }

    connect(userId: number): Observable<IMessage> {
        const api = environment.api.replace("/v1", "")

        const url = `${api}/sse?userId=${userId}`;

        this.eventSource = new EventSource(url)

        return new Observable((subscriber: Subscriber<IMessage>) => {
            this.eventSource.onerror = error => {
                this.zone.runOutsideAngular(() => subscriber.error(error));
            };

            this.eventNames.forEach((event: string) => {
                this.eventSource.addEventListener(event, data => {
                    if (event === MessageType.error) {
                        console.log("error", event);
                    } else {
                        this.zone.runOutsideAngular(() => {
                            subscriber.next(JSON.parse(data.data))
                        });
                    }
                });
            });
        });
    }

    close() {
        this.eventSource?.close();
    }
}
