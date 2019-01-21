import {BREAKPOINT} from '@angular/flex-layout';
import {Directive, AfterViewInit} from '@angular/core';
import {ShowHideDirective} from '@angular/flex-layout';

export const PRINT_BREAKPOINTS = [{
  alias: 'xs.print',
  suffix: 'XsPrint',
  mediaQuery: 'screen and (max-width: 897px)',
  overlapping: false
}];

export const CustomBreakPointsProvider = {
  provide: BREAKPOINT,
  useValue: PRINT_BREAKPOINTS,
  multi: true
};


const selector = `[fxHide.xs.print]`;
const inputs = ['fxHide.xs.print'];

// tslint:disable-next-line:use-input-property-decorator
@Directive({selector, inputs})
export class CustomShowHideDirective extends ShowHideDirective implements AfterViewInit{
  protected inputs = inputs;

  ngAfterViewInit() {
    super.ngAfterViewInit();
    console.log(this.inputs);
  }
}
