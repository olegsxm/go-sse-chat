import {ChangeDetectionStrategy, Component, EventEmitter, Output} from '@angular/core';

@Component({
    selector: 'app-chat-message-field',
    standalone: true,
    imports: [],
    templateUrl: './chat-message-field.component.html',
    styleUrl: './chat-message-field.component.scss',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class ChatMessageFieldComponent {
    @Output() send = new EventEmitter<string>();

    sendMessage(message: HTMLInputElement) {
        const txt = message.value;
        this.send.emit(txt);

        message.value = ""
    }
}
