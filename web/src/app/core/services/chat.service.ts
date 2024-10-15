import {Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {IConversation} from "../models/conversation.model";
import {Observable, of} from "rxjs";
import {IUser} from "../models/user.model";

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
}
