import {ChangeDetectionStrategy, Component, DestroyRef, HostListener, Input} from '@angular/core';
import {ChatHeaderComponent} from "../../components/chat-header/chat-header.component";
import {ChatMessagesComponent} from "../../components/chat-messages/chat-messages.component";
import {ChatMessageFieldComponent} from "../../components/chat-message-field/chat-message-field.component";
import {Router} from "@angular/router";
import {ChatService} from "../../core/services/chat.service";
import {ClickOutsideDirective} from "../../core/directives/click-outside.directive";
import {IConversation} from "../../core/models/conversation.model";
import {Store} from "@ngxs/store";
import {ChatState} from "../../state/chat/chat.state";
import {JsonPipe} from "@angular/common";
import {IMessage, IMessages} from "../../core/models/message.model";

@Component({
    selector: 'app-chat',
    standalone: true,
    imports: [
        ChatHeaderComponent,
        ChatMessagesComponent,
        ChatMessageFieldComponent,
        ClickOutsideDirective,
        JsonPipe
    ],
    templateUrl: './chat.component.html',
    styleUrl: './chat.component.scss',
    changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ChatComponent {
    @Input()
    set conversationId(id: number | null) {
        if (!id) return

        this.conversation = this.store.selectSnapshot(ChatState.getConversation(id)) || null;
    }

    conversation: IConversation | null = null;

    messages: IMessages[] = [];

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

    createMessage(msg: string) {
        const message: IMessage = {
            message: msg
        }
    }
}
