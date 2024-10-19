import {
    ChangeDetectionStrategy,
    Component,
    DestroyRef,
    HostListener,
    Input,
    OnChanges,
    OnInit,
    signal
} from '@angular/core';
import {ChatHeaderComponent} from "../../components/chat-header/chat-header.component";
import {ChatMessagesComponent} from "../../components/chat-messages/chat-messages.component";
import {ChatMessageFieldComponent} from "../../components/chat-message-field/chat-message-field.component";
import {Router} from "@angular/router";
import {ChatService} from "../../core/services/chat.service";
import {ClickOutsideDirective} from "../../core/directives/click-outside.directive";
import {IConversation} from "../../core/models/conversation.model";
import {Actions, ofActionDispatched, Store} from "@ngxs/store";
import {ChatState} from "../../state/chat/chat.state";
import {AsyncPipe, JsonPipe} from "@angular/common";
import {IMessage, IMessages} from "../../core/models/message.model";
import {map, Observable} from 'rxjs';
import {takeUntilDestroyed} from "@angular/core/rxjs-interop";
import {
    AddConversationsAction,
    AddNewMessageToQueueAction,
    RemoveMessageFromQueueAction,
    UpdateLastConversationMessage
} from "../../state/chat/chat.actions";


@Component({
    selector: 'app-chat',
    standalone: true,
    imports: [
        ChatHeaderComponent,
        ChatMessagesComponent,
        ChatMessageFieldComponent,
        ClickOutsideDirective,
        JsonPipe,
        AsyncPipe
    ],
    templateUrl: './chat.component.html',
    styleUrl: './chat.component.scss',
    changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ChatComponent implements OnInit, OnChanges {
    @Input() conversationId!: number;

    conversation$!: Observable<IConversation | undefined>;

    messages = signal<IMessages[]>([]);

    @HostListener('window:keyup.escape')
    exit() {
        const activeElement = document.activeElement;

        if (activeElement?.tagName.toLowerCase() !== 'body') return
        this.router.navigate(['/']);
    }

    constructor(
        private router: Router,
        private chatService: ChatService,
        private destroyRef: DestroyRef,
        private store: Store,
        private actions$: Actions
    ) {
    }

    ngOnInit() {
        this.actions$.pipe(
            ofActionDispatched(AddNewMessageToQueueAction),
            map(action => action.payload),
            takeUntilDestroyed(this.destroyRef),
        ).subscribe({
            next: message => {
                if (message.conversation === +this.conversationId) {
                    const messages = addMessageToMessageList(this.messages(), message)
                    this.messages.set(messages);
                    this.store.dispatch(new RemoveMessageFromQueueAction(message.id as number));
                    return
                }

                this.updateConversations()
            }
        })
    }


    ngOnChanges() {
        this.conversation$ = this.store.select(ChatState.getConversation(this.conversationId))
            .pipe(takeUntilDestroyed(this.destroyRef))

        this.chatService.getMessages(this.conversationId)
            .subscribe(res => {
                this.messages.set([]);
                res.forEach(msg => {
                    const messages = addMessageToMessageList(this.messages(), msg)
                    this.messages.set(messages)
                })
            })
    }

    createMessage(msg: string) {
        if (msg.trim().length < 1) return;

        this.chatService.sendMessage(this.conversationId, {
            message: msg,
        }).subscribe(res => {
            const messages = addMessageToMessageList(this.messages(), res)
            this.messages.set(messages);
            this.store.dispatch(new UpdateLastConversationMessage({msg: res, conversationId: +this.conversationId}))
        })
    }

    private updateConversations() {
        this.chatService.getConversation()
            .subscribe(res => {
                if (res && res.length) {
                    this.store.dispatch(new AddConversationsAction(res))
                }
            })
    }
}

function addMessageToMessageList(list: IMessages[], msg: IMessage): IMessages[] {
    const date = new Date(msg.createdAt as string).toJSON().substring(0, 10);
    const messages = [...list]
    const mIndex = messages.findIndex(m => m.date === date)
    if (mIndex === -1) {
        messages.push({date: date, messages: [msg]})
    } else {
        messages[mIndex].messages.push(msg)
    }

    return messages
}
