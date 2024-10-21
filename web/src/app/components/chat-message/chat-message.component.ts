import {ChangeDetectionStrategy, Component, Input} from '@angular/core';
import {DatePipe, NgTemplateOutlet} from "@angular/common";
import {IMessage} from "../../core/models/message.model";

@Component({
    selector: 'app-chat-message',
    standalone: true,
    imports: [
        NgTemplateOutlet,
        DatePipe
    ],
    templateUrl: './chat-message.component.html',
    styleUrl: './chat-message.component.scss',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class ChatMessageComponent {
    @Input() isUserMessage = false
    @Input() message!: IMessage;
}
