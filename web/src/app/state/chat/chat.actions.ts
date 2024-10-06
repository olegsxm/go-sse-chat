import {IConversation} from "../../core/models/conversation.model";

export class AddConversationsAction {
    static readonly type = '[Chat] add conversations';

    constructor(readonly payload: IConversation[]) {
    }
}
