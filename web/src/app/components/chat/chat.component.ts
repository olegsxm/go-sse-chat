import {ChangeDetectionStrategy, Component} from '@angular/core';
import {ChatHeaderComponent} from "../chat-header/chat-header.component";
import {ChatMessagesComponent} from "../chat-messages/chat-messages.component";
import {ChatMessageFieldComponent} from "../chat-message-field/chat-message-field.component";

@Component({
    selector: 'app-chat',
    standalone: true,
    imports: [
        ChatHeaderComponent,
        ChatMessagesComponent,
        ChatMessageFieldComponent
    ],
    templateUrl: './chat.component.html',
    styleUrl: './chat.component.scss',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class ChatComponent {

}
