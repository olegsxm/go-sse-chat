import {ChangeDetectionStrategy, Component} from '@angular/core';
import {RouterLink} from "@angular/router";

@Component({
    selector: 'app-sign-up-page',
    standalone: true,
    imports: [
        RouterLink
    ],
    templateUrl: './sign-up-page.component.html',
    styleUrl: './sign-up-page.component.scss',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class SignUpPageComponent {

}
