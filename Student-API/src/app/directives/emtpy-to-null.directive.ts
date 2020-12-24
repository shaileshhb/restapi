import { Directive, HostListener, Self } from '@angular/core';
import { NgControl } from '@angular/forms';

@Directive({
  selector: '[appEmtpyToNull]'
})
export class EmtpyToNullDirective {

  constructor(@Self() private ngControl: NgControl) { }


  @HostListener('keyup', ['$event']) 
  onKeyDowns(event: KeyboardEvent) {
    
    if (this.ngControl.value?.trim() === '') {
      console.log(this.ngControl.value);
      this.ngControl.reset(null);
    }
  }



}
