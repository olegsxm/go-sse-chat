export class ChatAction {
  static readonly type = '[Chat] Add item';
  constructor(readonly payload: string) { }
}
