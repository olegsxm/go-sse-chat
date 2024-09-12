import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterLink } from '@angular/router';
import { TuiAutoColorPipe, TuiIcon } from '@taiga-ui/core';
import { TuiAvatar, TuiAvatarOutline } from '@taiga-ui/kit';

@Component({
  selector: 'app-navigation',
  standalone: true,
  imports: [
    CommonModule,
    RouterLink,
    TuiIcon,
    TuiAvatar,
    TuiAvatarOutline,
    TuiAutoColorPipe,
  ],
  templateUrl: './navigation.component.html',
  styleUrl: './navigation.component.scss',
})
export class NavigationComponent {}
