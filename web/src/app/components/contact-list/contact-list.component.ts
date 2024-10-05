import {ChangeDetectionStrategy, Component} from '@angular/core';
import {ContactComponent} from "../contact/contact.component";
import {IContact} from "../../core/models/contact.model";

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
    contacts: IContact[] = [
        {
            id: 1,
            name: 'New Movie! Expendables 4',
            message: {
                id: 1,
                message: 'Get Andrés on this movie ASAP!',
                createdAt: new Date(2024, 11, 1, 12, 45)
            }
        },
        {
            id: 2,
            name: 'Arnold',
            message: {
                id: 2,
                message: "I'll be back",
                createdAt: new Date(2024, 11, 1, 12, 40)
            }
        },
        {
            id: 3,
            name: 'Russell Crowe',
            avatar: 'https://www.famousbirthdays.com/headshots/russell-crowe-6.jpg',
            message: {
                id: 3,
                message: "Hold the line!",
                createdAt: new Date(2024, 11, 1, 12, 39)
            }
        },
        {
            id: 4,
            name: 'Tom Cruise',
            avatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQGpYTzuO0zLW7yadaq4jpOz2SbsX90okb24Z9GtEvK6Z9x2zS5',
            message: {
                id: 3,
                message: "Show me the money!",
                createdAt: new Date(2024, 11, 1, 12, 38)
            }
        },
        {
            id: 5,
            name: 'Harrison Ford',
            message: {
                id: 3,
                message: "Tell Java I have the money",
                createdAt: new Date(2024, 10, 1, 12, 35)
            }
        },
    ];
}
