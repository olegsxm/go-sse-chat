import {ChangeDetectionStrategy, Component} from '@angular/core';
import {ChatMessageComponent} from "../chat-message/chat-message.component";

@Component({
    selector: 'app-chat-messages',
    standalone: true,
    imports: [
        ChatMessageComponent
    ],
    templateUrl: './chat-messages.component.html',
    styleUrl: './chat-messages.component.scss',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class ChatMessagesComponent {

}
