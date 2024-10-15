import {IUser} from "./user.model";
import {IMessage} from "./message.model";

export interface IConversation {
    id?: number;
    avatar?: string;
    name?: string;
    sender?: IUser;
    message?: IMessage;
    draft?: boolean;
}