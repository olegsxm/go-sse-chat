import {IUser} from "./user.model";

export interface IMessage {
    id?: number;
    message: string;
    createdAt?: string | Date;
    sender?: IUser;
}

export interface IMessages {
    date: string;
    messages: IMessage[];
}