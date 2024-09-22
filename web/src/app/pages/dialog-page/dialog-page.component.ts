import { ChangeDetectionStrategy, Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { DialogsComponent } from '../../components/dialogs/dialogs.component';
import { ChatComponent } from '../../components/chat/chat.component';

@Component({
  selector: 'app-dialog-page',
  standalone: true,
  imports: [CommonModule, DialogsComponent, ChatComponent],
  templateUrl: './dialog-page.component.html',
  styleUrl: './dialog-page.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class DialogPageComponent {}
