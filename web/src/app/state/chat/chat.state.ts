import {Injectable} from '@angular/core';
import {Action, Selector, State, StateContext} from '@ngxs/store';
import {IConversation} from "../../core/models/conversation.model";
import {AddConversationsAction} from "./chat.actions";

export interface IChatState {
    conversations: IConversation[];
}

@State<IChatState>({
    name: 'chat',
    defaults: {
        conversations: []
    }
})
@Injectable()
export class ChatState {

    @Selector()
    static getState(state: IChatState) {
        return state;
    }

    @Selector()
    static getDialogs(state: IChatState) {
        return state.conversations;
    }

    @Action(AddConversationsAction)
    addDialogs(ctx: StateContext<IChatState>, {payload}: AddConversationsAction) {
        ctx.patchState({
            conversations: [...payload, ...ctx.getState().conversations]
        })
    }
}
