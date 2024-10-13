import {Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {IConversation} from "../models/conversation.model";
import {Observable} from "rxjs";

@Injectable()
export class ChatService {

    constructor(private http: HttpClient) {
    }

    getConversation(): Observable<IConversation[]> {
        return this.http.get<IConversation[]>('/chat/conversations');
    }
}
