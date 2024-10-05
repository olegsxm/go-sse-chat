import {ChangeDetectionStrategy, Component, Input} from '@angular/core';
import {IContact} from "../../core/models/contact.model";
import {DatePipe} from "@angular/common";

@Component({
    selector: 'app-contact',
    standalone: true,
    imports: [
        DatePipe
    ],
    templateUrl: './contact.component.html',
    styleUrl: './contact.component.scss',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class ContactComponent {
    @Input() contact!: IContact;

}
