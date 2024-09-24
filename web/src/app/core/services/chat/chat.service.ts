import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { BehaviorSubject, Observable, of } from 'rxjs';
import { IChat } from '../../models/chat.model';
import { IUser } from '../../models/user.model';

@Injectable()
export class ChatService {
  chats$ = new BehaviorSubject<IChat[]>([]);
  activeChat$ = new BehaviorSubject<IChat | null>(null);

  constructor(private http: HttpClient) {
  }

  setActiveChat(activeChat: IChat) {
    this.activeChat$.next(activeChat);
  }

  getChats(): Observable<IChat[]> {
    return this.http.get<IChat[]>('/chat');
  }

  findUsers(login: string | null) {
    if (!login) return of([]);
    return this.http.get<IUser[]>(`/user/find?login=${login}`);
  }
}
