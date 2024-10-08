import {ChangeDetectionStrategy, Component, inject, Input} from '@angular/core';
import {DatePipe} from "@angular/common";
import {IConversation} from "../../core/models/conversation.model";
import {AvatarComponent} from "../avatar/avatar.component";
import {Router} from "@angular/router";

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

    router = inject(Router)

    openConversation() {
        this.router.navigate(['./', this.conversation.id])
    }
}
