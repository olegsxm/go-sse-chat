import {ChangeDetectionStrategy, Component, Input} from '@angular/core';
import {ContactComponent} from "../contact/contact.component";
import {IConversation} from "../../core/models/conversation.model";

@Component({
    selector: 'app-contact-list',
    standalone: true,
    imports: [
        ContactComponent
    ],
    templateUrl: './contact-list.component.html',
    styleUrl: './contact-list.component.scss',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class ContactListComponent {
    @Input() conversations: IConversation[] = [];


}
