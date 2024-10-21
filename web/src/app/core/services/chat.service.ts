import {Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {IConversation} from "../models/conversation.model";
import {Observable, of} from "rxjs";
import {IUser} from "../models/user.model";
import {IMessage} from "../models/message.model";

@Injectable()
export class ChatService {

    constructor(private http: HttpClient) {
    }

    findUsers(query: string): Observable<IUser[]> {
        if (!query || query.length === 0) {
            return of([])
        }

        return this.http.get<IUser[]>(`/users/find?query=${query}`);
    }

    getConversation(): Observable<IConversation[]> {
        return this.http.get<IConversation[]>('/chat/conversations');
    }

    createConversation(to: IUser): Observable<IConversation> {
        return this.http.post<IConversation>(`/chat/conversations`, {
            to: to.id
        });
    }

    sendMessage(conversationId: number, message: IMessage): Observable<IMessage> {
        return this.http.post<IMessage>(`/chat/conversation/${conversationId}/create-message`, message);
    }

    getMessages(conversationId: number): Observable<IMessage[]> {
        // /chat/conversation/:conversationId/messages
        return this.http.get<IMessage[]>(`/chat/conversation/${conversationId}/messages`);
    }
}
