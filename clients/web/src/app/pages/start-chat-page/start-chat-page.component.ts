import { ChangeDetectionStrategy, Component } from '@angular/core';

@Component({
  selector: 'app-start-chat-page',
  standalone: true,
  imports: [],
  templateUrl: './start-chat-page.component.html',
  styleUrl: './start-chat-page.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class StartChatPageComponent {

}
