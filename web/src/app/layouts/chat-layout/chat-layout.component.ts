import { ChangeDetectionStrategy, Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { NavigationComponent } from '../../components/navigation/navigation.component';
import { DialogsComponent } from '../../components/dialogs/dialogs.component';
import { DialogComponent } from '../../components/dialog/dialog.component';
import { ContactDetailsComponent } from '../../components/contact-details/contact-details.component';
import { RouterOutlet } from '@angular/router';

@Component({
  selector: 'app-chat-layout',
  standalone: true,
  imports: [
    CommonModule,
    NavigationComponent,
    DialogsComponent,
    DialogComponent,
    ContactDetailsComponent,
    RouterOutlet,
  ],
  templateUrl: './chat-layout.component.html',
  styleUrl: './chat-layout.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ChatLayoutComponent {}
