import {Injectable} from '@angular/core';
import {Action, createSelector, Selector, State, StateContext} from '@ngxs/store';
import {IConversation} from "../../core/models/conversation.model";
import {
    AddConversationsAction,
    AddNewMessageToQueueAction,
    RemoveMessageFromQueueAction,
    UpdateLastConversationMessage
} from "./chat.actions";
import {IMessage} from "../../core/models/message.model";

export interface IChatState {
    conversations: IConversation[];
    messagesQueue: IMessage[] // Messages from sse or ws
}

@State<IChatState>({
    name: 'chat',
    defaults: {
        conversations: [],
        messagesQueue: []
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

    @Action(UpdateLastConversationMessage)
    updateLastConversationMessage(state: StateContext<IChatState>, {payload}: UpdateLastConversationMessage) {
        const conversations = state.getState().conversations.map(c => {
            if (c.id === payload.conversationId) {
                return {...c, message: payload.msg}
            }
            return c
        })

        state.patchState({
            conversations: conversations
        })
    }

    @Action(AddNewMessageToQueueAction)
    addNewMessageToQueue(state: StateContext<IChatState>, {payload}: AddNewMessageToQueueAction) {
        const messages = [...state.getState().messagesQueue, payload]
        state.patchState({
            messagesQueue: messages
        })
    }

    @Action(RemoveMessageFromQueueAction)
    removeMessageFromQueue(state: StateContext<IChatState>, {payload}: RemoveMessageFromQueueAction) {
        const messages = state.getState().messagesQueue.filter(m => m.id !== payload)
        state.patchState({
            messagesQueue: messages
        })
    }
}
