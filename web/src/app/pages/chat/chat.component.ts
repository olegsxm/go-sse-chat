import {ChangeDetectionStrategy, Component, DestroyRef, HostListener, Input, OnChanges} from '@angular/core';
import {ChatHeaderComponent} from "../../components/chat-header/chat-header.component";
import {ChatMessagesComponent} from "../../components/chat-messages/chat-messages.component";
import {ChatMessageFieldComponent} from "../../components/chat-message-field/chat-message-field.component";
import {Router} from "@angular/router";
import {ChatService} from "../../core/services/chat.service";

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
export class ChatComponent implements OnChanges {
    @Input() conversation: number | null = null;

    @HostListener('window:keyup.escape')
    exit() {
        this.router.navigate(['/']);
    }


    constructor(
        private router: Router,
        private chatService: ChatService,
        private destroyRef: DestroyRef
    ) {
    }


    ngOnChanges() {
        console.log(this.conversation);
    }
}
