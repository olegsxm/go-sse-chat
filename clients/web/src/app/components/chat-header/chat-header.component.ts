import {ChangeDetectionStrategy, Component, Input} from '@angular/core';
import {AvatarComponent} from "../avatar/avatar.component";

@Component({
    selector: 'app-chat-header',
    standalone: true,
    imports: [
        AvatarComponent
    ],
    templateUrl: './chat-header.component.html',
    styleUrl: './chat-header.component.scss',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class ChatHeaderComponent {
    @Input() avatar?: string;
    @Input({required: true}) name!: string;
}
