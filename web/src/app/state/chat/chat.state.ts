import { Injectable } from '@angular/core';
import { State, Action, Selector, StateContext } from '@ngxs/store';
import { ChatAction } from './chat.actions';

export interface ChatStateModel {
  items: string[];
}

@State<ChatStateModel>({
  name: 'chat',
  defaults: {
    items: []
  }
})
@Injectable()
export class ChatState {

  @Selector()
  static getState(state: ChatStateModel) {
    return state;
  }

  @Action(ChatAction)
  add(ctx: StateContext<ChatStateModel>, { payload }: ChatAction) {
    const stateModel = ctx.getState();
    stateModel.items = [...stateModel.items, payload];
    ctx.setState(stateModel);
  }
}
