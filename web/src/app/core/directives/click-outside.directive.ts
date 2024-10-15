import {Directive, ElementRef, EventEmitter, HostListener, Output} from '@angular/core';

@Directive({
    // eslint-disable-next-line @angular-eslint/directive-selector
    selector: '[clickOutside]',
    standalone: true
})
export class ClickOutsideDirective {
    constructor(private elementRef: ElementRef) {
    }

    @Output() clickOutside: EventEmitter<null> = new EventEmitter();

    @HostListener('document: click', ['$event.target']) onMouseEnter(targetElement: unknown) {
        const clickInside = this.elementRef.nativeElement.contains(targetElement);
        if (!clickInside) {
            this.clickOutside.emit(null);
        }
    }

}
