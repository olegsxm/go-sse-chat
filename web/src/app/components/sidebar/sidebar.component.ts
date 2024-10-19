import {ChangeDetectionStrategy, Component, EventEmitter, Input, Output} from '@angular/core';
import {CommonModule} from '@angular/common';
import {SearchComponent} from "../search/search.component";
import {ContactListComponent} from "../contact-list/contact-list.component";
import {AvatarComponent} from "../avatar/avatar.component";
import {IConversation} from "../../core/models/conversation.model";
import {IUser} from "../../core/models/user.model";
import {Router} from "@angular/router";
import {Store} from "@ngxs/store";

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
    @Input({required: true}) user: IUser | null = null;
    @Input() conversations: IConversation[] = [];

    @Output() writeToUser = new EventEmitter<IUser>();

    constructor(private router: Router, private store: Store) {
    }

}
