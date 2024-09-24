import { ChangeDetectionStrategy, Component, Input, OnChanges, OnDestroy } from '@angular/core';
import { TuiAvatar } from '@taiga-ui/kit';
import { TuiAutoColorPipe, TuiButton, TuiScrollable, TuiScrollbar } from '@taiga-ui/core';
import { MessageComponent } from '../message/message.component';
import { CdkFixedSizeVirtualScroll, CdkVirtualForOf, CdkVirtualScrollViewport } from '@angular/cdk/scrolling';
import { FormControl, ReactiveFormsModule } from '@angular/forms';
import { IChat } from '../../core/models/chat.model';
import { ChatService } from '../../core/services/chat/chat.service';

@Component({
  selector: 'app-chat',
  standalone: true,
  imports: [
    TuiAvatar,
    TuiAutoColorPipe,
    MessageComponent,
    CdkFixedSizeVirtualScroll,
    TuiScrollbar,
    TuiScrollable,
    CdkVirtualForOf,
    CdkVirtualScrollViewport,
    ReactiveFormsModule,
    TuiButton
  ],
  templateUrl: './chat.component.html',
  styleUrl: './chat.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class ChatComponent implements OnChanges, OnDestroy {
  @Input() chat!: IChat | null;

  messages = new Array(250).fill(0);
  messageForm = new FormControl();

  constructor(private chatService: ChatService) {
  }

  ngOnChanges() {
    console.log(this.chat?.id);
  }

  ngOnDestroy() {
    this.messageForm.setValue(null);
  }
}
