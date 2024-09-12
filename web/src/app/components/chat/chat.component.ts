import { ChangeDetectionStrategy, Component } from '@angular/core';
import { TuiAvatar } from '@taiga-ui/kit';
import {
  TuiAutoColorPipe,
  TuiButton,
  TuiScrollable,
  TuiScrollbar,
} from '@taiga-ui/core';
import { MessageComponent } from '../message/message.component';
import {
  CdkFixedSizeVirtualScroll,
  CdkVirtualForOf,
  CdkVirtualScrollViewport,
} from '@angular/cdk/scrolling';
import { FormControl, ReactiveFormsModule } from '@angular/forms';

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
    TuiButton,
  ],
  templateUrl: './chat.component.html',
  styleUrl: './chat.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ChatComponent {
  messages = new Array(250).fill(0);
  messageForm = new FormControl();
}
