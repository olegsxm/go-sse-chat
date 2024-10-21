import {Observable} from "rxjs";
import {IMessage} from "./message.model";

export interface IMessageService {
    connect(userId: number): Observable<IMessage>;

    close(): void;
}