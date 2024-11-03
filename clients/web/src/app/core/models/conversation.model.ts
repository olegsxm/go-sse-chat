import {IMessage} from "./message.model";

export interface IConversation {
    id?: number;
    avatar?: string;
    name?: string;
    message?: IMessage;
    draft?: boolean;
}