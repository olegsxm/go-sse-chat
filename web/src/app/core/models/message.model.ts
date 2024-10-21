import {IUser} from "./user.model";

export interface IMessage {
    id?: number;
    message: string;
    createdAt?: string | Date;
    sender?: IUser;
    conversation?: number;
    senderId?: number;
}

export interface IMessages {
    date: string;
    messages: IMessage[];
}