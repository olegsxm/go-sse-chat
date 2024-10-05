import {Component, OnInit} from '@angular/core';
import {RouterModule} from '@angular/router';
import {SidebarComponent} from "./components/sidebar/sidebar.component";
import {ChatComponent} from "./components/chat/chat.component";
import {initFlowbite} from "flowbite";

@Component({
    standalone: true,
    imports: [RouterModule, SidebarComponent, ChatComponent],
    selector: 'app-root',
    templateUrl: './app.component.html',
    styleUrl: './app.component.scss',
})
export class AppComponent implements OnInit {
    ngOnInit() {
        initFlowbite();
    }
}
