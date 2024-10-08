import {ChangeDetectionStrategy, Component, Input} from '@angular/core';
import {NgTemplateOutlet} from "@angular/common";

@Component({
    selector: 'app-chat-message',
    standalone: true,
    imports: [
        NgTemplateOutlet
    ],
    templateUrl: './chat-message.component.html',
    styleUrl: './chat-message.component.scss',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class ChatMessageComponent {
    @Input() isUserMessage = false
}
