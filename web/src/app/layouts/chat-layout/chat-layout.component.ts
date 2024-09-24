import { ChangeDetectionStrategy, Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { NavigationComponent } from '../../components/navigation/navigation.component';
import { DialogsComponent } from '../../components/dialogs/dialogs.component';
import { DialogComponent } from '../../components/dialog/dialog.component';
import { ContactDetailsComponent } from '../../components/contact-details/contact-details.component';
import { RouterOutlet } from '@angular/router';
import { ChatService } from '../../core/services/chat/chat.service';

@Component({
  selector: 'app-chat-layout',
  standalone: true,
  imports: [
    CommonModule,
    NavigationComponent,
    DialogsComponent,
    DialogComponent,
    ContactDetailsComponent,
    RouterOutlet
  ],
  providers: [
    ChatService
  ],
  templateUrl: './chat-layout.component.html',
  styleUrl: './chat-layout.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class ChatLayoutComponent implements OnInit {
  constructor(private chatService: ChatService) {
  }


  ngOnInit() {
    this.chatService.getChats()
      .subscribe(res => {
        this.chatService.chats$.next(res);
      });
  }
}
