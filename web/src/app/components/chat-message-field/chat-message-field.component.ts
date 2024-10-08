import { ChangeDetectionStrategy, Component } from '@angular/core';

@Component({
  selector: 'app-chat-message-field',
  standalone: true,
  imports: [],
  templateUrl: './chat-message-field.component.html',
  styleUrl: './chat-message-field.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class ChatMessageFieldComponent {

}
