import {
    ChangeDetectionStrategy,
    Component,
    DestroyRef,
    ElementRef,
    EventEmitter,
    HostListener,
    OnInit,
    Output,
    signal,
    ViewChild
} from '@angular/core';
import {CommonModule} from '@angular/common';
import {FormControl, ReactiveFormsModule} from "@angular/forms";
import {debounceTime, switchMap} from "rxjs";
import {takeUntilDestroyed} from "@angular/core/rxjs-interop";
import {ChatService} from "../../core/services/chat.service";
import {IUser} from "../../core/models/user.model";
import {ClickOutsideDirective} from "../../core/directives/click-outside.directive";

@Component({
    selector: 'app-search',
    standalone: true,
    imports: [CommonModule, ReactiveFormsModule, ClickOutsideDirective],
    templateUrl: './search.component.html',
    styleUrl: './search.component.scss',
    changeDetection: ChangeDetectionStrategy.OnPush,
})
export class SearchComponent implements OnInit {
    @Output() selectUser = new EventEmitter<IUser>();
    @ViewChild("field") searchField: ElementRef | undefined;

    search = new FormControl();
    users = signal<IUser[]>([])

    @HostListener('window:keyup.escape')
    escape() {
        if (this.search.value?.length === 0) return

        if (!document.activeElement?.className.includes('search-field')) return;

        this.clear()
    }

    constructor(private destroyRef: DestroyRef, private chatService: ChatService) {
    }

    ngOnInit() {
        this.search.valueChanges
            .pipe(
                debounceTime(300),
                switchMap(query => this.chatService.findUsers(query)),
                takeUntilDestroyed(this.destroyRef)
            )
            .subscribe({
                next: users => this.users.set(users),
                error: () => {
                    this.users.set([]);
                }
            })
    }

    createConversations(user: IUser) {
        this.selectUser.emit(user);
        this.clear()
    }

    clear() {
        this.users.set([])
        this.search.setValue(null);
    }
}