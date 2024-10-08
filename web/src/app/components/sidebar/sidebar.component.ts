import {ChangeDetectionStrategy, Component, Input} from '@angular/core';
import {CommonModule} from '@angular/common';
import {SearchComponent} from "../search/search.component";
import {ContactListComponent} from "../contact-list/contact-list.component";
import {AvatarComponent} from "../avatar/avatar.component";
import {IConversation} from "../../core/models/conversation.model";
import {IUser} from "../../core/models/user.model";

export interface ISidebarData {
    user: IUser | null;
    dialogs: IConversation[]
}

@Component({
    selector: 'app-sidebar',
    standalone: true,
    imports: [CommonModule, SearchComponent, ContactListComponent, AvatarComponent],
    templateUrl: './sidebar.component.html',
    styleUrl: './sidebar.component.scss',
    changeDetection: ChangeDetectionStrategy.OnPush,
})
export class SidebarComponent {
    @Input({required: true}) data!: ISidebarData;
}
