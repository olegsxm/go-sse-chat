import {ChangeDetectionStrategy, Component} from '@angular/core';
import {ChatComponent} from "../../components/chat/chat.component";
import {SidebarComponent} from "../../components/sidebar/sidebar.component";

@Component({
    selector: 'app-chat-page',
    standalone: true,
    imports: [
        ChatComponent,
        SidebarComponent
    ],
    templateUrl: './chat-page.component.html',
    styleUrl: './chat-page.component.scss',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class ChatPageComponent {

}
