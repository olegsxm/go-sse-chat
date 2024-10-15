import {ChangeDetectionStrategy, Component, Input} from '@angular/core';
import {ChatMessageComponent} from "../chat-message/chat-message.component";
import {IMessages} from 'src/app/core/models/message.model';
import {DatePipe} from "@angular/common";

@Component({
    selector: 'app-chat-messages',
    standalone: true,
    imports: [
        ChatMessageComponent,
        DatePipe
    ],
    templateUrl: './chat-messages.component.html',
    styleUrl: './chat-messages.component.scss',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class ChatMessagesComponent {
    @Input() messages!: IMessages[];

}
