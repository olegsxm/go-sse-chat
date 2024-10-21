import {
    AfterViewInit,
    ChangeDetectionStrategy,
    Component,
    ElementRef,
    inject,
    Input,
    OnChanges,
    ViewChild
} from '@angular/core';
import {ChatMessageComponent} from "../chat-message/chat-message.component";
import {IMessages} from 'src/app/core/models/message.model';
import {DatePipe} from "@angular/common";
import {Store} from "@ngxs/store";
import {AuthState} from "../../state/auth/auth.state";

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
export class ChatMessagesComponent implements OnChanges, AfterViewInit {
    @Input() messages!: IMessages[];

    @ViewChild('messagesViewBox') messagesViewBox!: ElementRef<HTMLDivElement>;

    store = inject(Store)
    owner = this.store.selectSnapshot(AuthState.getState)

    ngOnChanges() {
        this.scroll();
    }

    ngAfterViewInit() {
        this.scroll();
    }

    private scroll() {
        setTimeout(() => {
            if (this.messagesViewBox) {
                this.messagesViewBox.nativeElement.scrollIntoView(false)
            }
        }, 100)
    }
}
