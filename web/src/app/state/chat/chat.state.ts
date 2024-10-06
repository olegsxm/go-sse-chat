import {Injectable} from '@angular/core';
import {Action, Selector, State, StateContext} from '@ngxs/store';
import {IConversation} from "../../core/models/conversation.model";
import {AddDialogsAction} from "./chat.actions";

export interface ChatStateModel {
    dialogs: IConversation[];
}

@State<ChatStateModel>({
    name: 'chat',
    defaults: {
        dialogs: []
    }
})
@Injectable()
export class ChatState {

    @Selector()
    static getState(state: ChatStateModel) {
        return state;
    }

    @Selector()
    static getDialogs(state: ChatStateModel) {
        return state.dialogs;
    }

    @Action(AddDialogsAction)
    addDialogs(ctx: StateContext<ChatStateModel>, {payload}: AddDialogsAction) {
        ctx.patchState({
            dialogs: [...payload, ...ctx.getState().dialogs]
        })
    }
}
