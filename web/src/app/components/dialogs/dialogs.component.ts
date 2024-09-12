import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import {
  TuiIcon,
  TuiScrollable,
  TuiScrollbar,
  tuiScrollbarOptionsProvider,
  TuiTextfieldComponent,
  TuiTextfieldDirective,
  TuiTextfieldOptionsDirective,
} from '@taiga-ui/core';
import { DialogComponent } from '../dialog/dialog.component';
import {
  CdkFixedSizeVirtualScroll,
  CdkVirtualForOf,
  CdkVirtualScrollViewport,
} from '@angular/cdk/scrolling';

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
  ],
  templateUrl: './dialogs.component.html',
  styleUrl: './dialogs.component.scss',
  providers: [
    tuiScrollbarOptionsProvider({
      mode: 'hover',
    }),
  ],
})
export class DialogsComponent {
  dialogs = new Array(25).fill(0);

  constructor() {
    console.log(this.dialogs);
  }
}
