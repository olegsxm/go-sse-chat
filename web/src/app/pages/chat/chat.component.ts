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
import {Store} from "@ngxs/store";
import {ChatState} from "../../state/chat/chat.state";
import {AsyncPipe, JsonPipe} from "@angular/common";
import {IMessage, IMessages} from "../../core/models/message.model";
import {Observable} from 'rxjs';
import {takeUntilDestroyed} from "@angular/core/rxjs-interop";


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
        private store: Store
    ) {
    }

    ngOnInit() {
        this.chatService.getMessages(this.conversationId)
            .subscribe(res => {
                res.forEach(msg => {
                    const messages = addMessageToMessageList(this.messages(), msg)
                    this.messages.set(messages)
                })
            })
    }

    ngOnChanges() {
        this.conversation$ = this.store.select(ChatState.getConversation(this.conversationId))
            .pipe(takeUntilDestroyed(this.destroyRef))
    }

    createMessage(msg: string) {
        if (msg.trim().length < 1) return;
        
        this.chatService.sendMessage(this.conversationId, {
            message: msg,
        }).subscribe(res => {
            const messages = addMessageToMessageList(this.messages(), res)
            this.messages.set(messages)
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
