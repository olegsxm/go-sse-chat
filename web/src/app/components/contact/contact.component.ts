import {ChangeDetectionStrategy, Component, Input} from '@angular/core';
import {DatePipe} from "@angular/common";
import {IConversation} from "../../core/models/conversation.model";
import {AvatarComponent} from "../avatar/avatar.component";

@Component({
    selector: 'app-contact',
    standalone: true,
    imports: [
        DatePipe,
        AvatarComponent
    ],
    templateUrl: './contact.component.html',
    styleUrl: './contact.component.scss',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class ContactComponent {
    @Input() conversation!: IConversation;

}
