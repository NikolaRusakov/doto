import { Directive } from '@angular/core';
import { ShowHideDirective ,LayoutDirective} from '@angular/flex-layout';

const selector = `[fxLayout.yba-min-width]`;
const inputs = ['fxLayout.yba-min-width'];

@Directive({ selector, inputs })
export class CustomLayoutDirective extends LayoutDirective {
  protected inputs = inputs;
}
