import {IMessage} from "./message.model";

export interface IContact {
    id: number;
    avatar?: string;
    name: string;
    message?: IMessage
}