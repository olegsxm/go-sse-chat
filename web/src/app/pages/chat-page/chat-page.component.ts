import {ChangeDetectionStrategy, Component, Inject, inject, OnInit} from '@angular/core';
import {ChatComponent} from "../chat/chat.component";
import {SidebarComponent} from "../../components/sidebar/sidebar.component";
import {ChatService} from "../../core/services/chat.service";
import {Store} from "@ngxs/store";
import {AuthState} from "../../state/auth/auth.state";
import {ChatState} from "../../state/chat/chat.state";
import {
    AddConversationsAction,
    AddNewMessageToQueueAction,
    UpdateLastConversationMessage
} from "../../state/chat/chat.actions";
import {RouterOutlet} from "@angular/router";
import {IUser} from "../../core/models/user.model";
import {AsyncPipe} from "@angular/common";
import {SseService} from "../../core/services/sse.service";
import {environment} from "../../../environments/environment";
import {WsService} from "../../core/services/ws.service";
import {MessageService} from "../../core/tokens/message-service.token";
import {IMessageService} from "../../core/models/message-service.model";
import {IMessage} from "../../core/models/message.model";

@Component({
    selector: 'app-chat-page',
    standalone: true,
    imports: [
        ChatComponent,
        SidebarComponent,
        RouterOutlet,
        AsyncPipe
    ],
    providers: [
        ChatService,
        {
            provide: MessageService,
            useClass: environment.broker === 'ws' ? WsService : SseService,
        }
    ],
    templateUrl: './chat-page.component.html',
    styleUrl: './chat-page.component.scss',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class ChatPageComponent implements OnInit {
    private store = inject(Store)

    user$ = this.store.select(AuthState.getUser);
    dialogs$ = this.store.select(ChatState.getConversations);

    constructor(private chatService: ChatService, @Inject(MessageService) private messageService: IMessageService) {
        this.chatService.getConversation()
            .subscribe(res => {
                if (res && res.length) {
                    this.store.dispatch(new AddConversationsAction(res))
                }
            })
    }

    ngOnInit() {
        const userid = this.store.selectSnapshot(AuthState.getState).user?.id

        this.messageService.connect(userid as number)
            .subscribe({
                next: (res: IMessage) => {
                    this.store.dispatch(new AddNewMessageToQueueAction(res))
                    this.store.dispatch(new UpdateLastConversationMessage({
                        msg: res,
                        conversationId: res.conversation as number
                    }))
                }
            })
    }

    writeToUser(user: IUser) {
        this.chatService.createConversation(user)
            .subscribe(
                res => this.store.dispatch(new AddConversationsAction([res]))
            )
    }
}
