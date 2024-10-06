import {IConversation} from "../../core/models/conversation.model";

export class AddDialogsAction {
    static readonly type = '[Chat] add dialogs';

    constructor(readonly payload: IConversation[]) {
    }
}
