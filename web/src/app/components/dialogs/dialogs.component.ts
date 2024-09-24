import { Component, EventEmitter, Input, Output } from '@angular/core';
import { CommonModule } from '@angular/common';
import {
  TuiIcon,
  TuiScrollable,
  TuiScrollbar,
  tuiScrollbarOptionsProvider,
  TuiTextfieldComponent,
  TuiTextfieldDirective,
  TuiTextfieldOptionsDirective
} from '@taiga-ui/core';
import { DialogComponent } from '../dialog/dialog.component';
import { CdkFixedSizeVirtualScroll, CdkVirtualForOf, CdkVirtualScrollViewport } from '@angular/cdk/scrolling';


import { IChat } from '../../core/models/chat.model';
import { ChatService } from '../../core/services/chat/chat.service';
import { FormControl } from '@angular/forms';
import { SearchComponent } from '../search.component';

@Component({
  selector: 'app-dialogs',
  standalone: true,
  imports: [
    CommonModule,
    TuiTextfieldComponent,
    TuiTextfieldDirective,
    TuiTextfieldOptionsDirective,
    TuiIcon,
    DialogComponent,
    TuiScrollbar,
    CdkVirtualScrollViewport,
    TuiScrollable,
    CdkFixedSizeVirtualScroll,
    CdkVirtualForOf,
    SearchComponent
  ],
  templateUrl: './dialogs.component.html',
  styleUrl: './dialogs.component.scss',
  providers: [
    tuiScrollbarOptionsProvider({
      mode: 'hover'
    })
  ]
})
export class DialogsComponent {
  @Input() dialogs: IChat[] = [];
  @Output() setActiveChat = new EventEmitter<IChat>();

  search = new FormControl(null);

  constructor(private chatService: ChatService) {
  }

}
