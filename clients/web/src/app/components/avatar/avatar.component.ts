import {ChangeDetectionStrategy, Component, Input} from '@angular/core';
import {NgOptimizedImage} from "@angular/common";

@Component({
    selector: 'app-avatar',
    standalone: true,
    imports: [
        NgOptimizedImage
    ],
    templateUrl: './avatar.component.html',
    styleUrl: './avatar.component.scss',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class AvatarComponent {
    avatar!: string;

    @Input()
    set src(value: string | undefined) {
        if (!value) {
            this.avatar = "/images/default-avatar.png";
            return
        }

        this.avatar = value;
    }

    constructor() {
        this.src = undefined;
    }

    error() {
        this.src = undefined;
    }
}
