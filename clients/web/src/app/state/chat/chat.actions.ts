import {IConversation} from "../../core/models/conversation.model";
import {IMessage} from "../../core/models/message.model";

export class AddConversationsAction {
    static readonly type = '[Chat] add conversations';

    constructor(readonly payload: IConversation[]) {
    }
}

export class UpdateLastConversationMessage {
    static readonly type = '[Chat] update last conversation message';

    constructor(readonly payload: { conversationId: number, msg: IMessage }) {
    }

}

export class AddNewMessageToQueueAction {
    static readonly type = '[Chat] add new message';

    constructor(readonly payload: IMessage) {
    }
}

export class RemoveMessageFromQueueAction {
    static readonly type = '[Chat] remove message from queue';

    constructor(readonly payload: number) {
    }
}