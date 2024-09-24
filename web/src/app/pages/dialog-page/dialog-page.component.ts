import { ChangeDetectionStrategy, Component, DestroyRef, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { DialogsComponent } from '../../components/dialogs/dialogs.component';
import { ChatComponent } from '../../components/chat/chat.component';
import { ChatService } from '../../core/services/chat/chat.service';
import { Observable } from 'rxjs';
import { IChat } from '../../core/models/chat.model';

@Component({
  selector: 'app-dialog-page',
  standalone: true,
  imports: [CommonModule, DialogsComponent, ChatComponent],
  templateUrl: './dialog-page.component.html',
  styleUrl: './dialog-page.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class DialogPageComponent implements OnInit {
  chats$!: Observable<IChat[]>;
  chat$!: Observable<IChat | null>;

  constructor(private chatService: ChatService, private destroy: DestroyRef) {
  }

  ngOnInit() {
    this.chats$ = this.chatService.chats$;
    this.chat$ = this.chatService.activeChat$;
  }

  setActiveChat(chat: IChat) {
    this.chatService.setActiveChat(chat);
  }
}
