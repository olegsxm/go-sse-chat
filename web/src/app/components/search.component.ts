import { ChangeDetectionStrategy, ChangeDetectorRef, Component, DestroyRef, inject, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormControl, ReactiveFormsModule } from '@angular/forms';
import { ChatService } from '../core/services/chat/chat.service';
import { debounceTime, switchMap } from 'rxjs';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { IUser } from '../core/models/user.model';

@Component({
  selector: 'app-search',
  standalone: true,
  imports: [CommonModule, ReactiveFormsModule],
  templateUrl: './search.component.html',
  styleUrl: './search.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class SearchComponent implements OnInit {
  search = new FormControl<string | null>(null);
  destroyRef = inject(DestroyRef);

  users: IUser[] = [];

  constructor(private chatService: ChatService, private cdr: ChangeDetectorRef) {
  }

  ngOnInit() {
    this.search.valueChanges
      .pipe(
        debounceTime(200),
        switchMap((value) => this.chatService.findUsers(value)),
        takeUntilDestroyed(this.destroyRef)
      ).subscribe(v => {
      this.users = v;
      this.cdr.detectChanges();
    });
  }
}
