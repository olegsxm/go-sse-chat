import {ChangeDetectionStrategy, Component} from '@angular/core';
import {CommonModule} from '@angular/common';
import {SearchComponent} from "../search/search.component";
import {ContactListComponent} from "../contact-list/contact-list.component";

@Component({
    selector: 'app-sidebar',
    standalone: true,
    imports: [CommonModule, SearchComponent, ContactListComponent],
    templateUrl: './sidebar.component.html',
    styleUrl: './sidebar.component.scss',
    changeDetection: ChangeDetectionStrategy.OnPush,
})
export class SidebarComponent {
}
