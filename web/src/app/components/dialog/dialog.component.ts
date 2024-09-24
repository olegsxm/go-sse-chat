import { Component, Input } from '@angular/core';
import { CommonModule } from '@angular/common';
import { TuiAvatar } from '@taiga-ui/kit';
import { TuiTitle } from '@taiga-ui/core';
import { IChat } from '../../core/models/chat.model';

@Component({
  selector: 'app-dialog',
  standalone: true,
  imports: [CommonModule, TuiAvatar, TuiTitle],
  templateUrl: './dialog.component.html',
  styleUrl: './dialog.component.scss'
})
export class DialogComponent {
  @Input() dialog!: IChat;
}
