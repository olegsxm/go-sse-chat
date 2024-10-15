import {ChangeDetectionStrategy, Component, inject} from '@angular/core';
import {ChatComponent} from "../chat/chat.component";
import {SidebarComponent} from "../../components/sidebar/sidebar.component";
import {ChatService} from "../../core/services/chat.service";
import {Store} from "@ngxs/store";
import {AuthState} from "../../state/auth/auth.state";
import {ChatState} from "../../state/chat/chat.state";
import {AddConversationsAction} from "../../state/chat/chat.actions";
import {RouterOutlet} from "@angular/router";
import {IUser} from "../../core/models/user.model";

@Component({
    selector: 'app-chat-page',
    standalone: true,
    imports: [
        ChatComponent,
        SidebarComponent,
        RouterOutlet
    ],
    providers: [
        ChatService,
    ],
    templateUrl: './chat-page.component.html',
    styleUrl: './chat-page.component.scss',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class ChatPageComponent {
    private store = inject(Store)

    readonly user = this.store.selectSignal(AuthState.getUser);
    readonly dialogs = this.store.selectSignal(ChatState.getConversations);


    constructor(private chatService: ChatService) {
        this.chatService.getConversation()
            .subscribe(res => {
                if (res && res.length) {
                    this.store.dispatch(new AddConversationsAction(res))
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
