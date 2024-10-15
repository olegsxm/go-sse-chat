import {Injectable} from '@angular/core';
import {Action, createSelector, Selector, State, StateContext} from '@ngxs/store';
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
    static getConversations(state: IChatState) {
        return state.conversations;
    }

    static getConversation(id: number) {
        return createSelector([ChatState], (state: IChatState) => {
            return state.conversations.find(conversation => {
                return conversation.id === +id
            });
        })
    }

    @Action(AddConversationsAction)
    addConversations(ctx: StateContext<IChatState>, {payload}: AddConversationsAction) {
        ctx.patchState({
            conversations: [...payload, ...ctx.getState().conversations]
        })
    }
}
