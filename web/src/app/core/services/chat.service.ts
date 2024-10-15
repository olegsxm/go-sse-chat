import {Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {IConversation} from "../models/conversation.model";
import {Observable} from "rxjs";
import {IUser} from "../models/user.model";

@Injectable()
export class ChatService {

    constructor(private http: HttpClient) {
    }

    findUsers(query: string): Observable<IUser[]> {
        return this.http.get<IUser[]>(`/users/find?query=${query}`);
    }

    getConversation(): Observable<IConversation[]> {
        return this.http.get<IConversation[]>('/chat/conversations');
    }
}
