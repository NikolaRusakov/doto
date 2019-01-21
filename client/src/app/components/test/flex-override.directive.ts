import { Directive } from '@angular/core';
import { ShowHideDirective ,LayoutDirective} from '@angular/flex-layout';

const selector = `[fxHide.yba-min-height]`;
const inputs = ['fxHide.yba-min-height'];

// tslint:disable-next-line:use-input-property-decorator
@Directive({ selector, inputs })
export class CustomFlexOverrideDirective extends ShowHideDirective {
  protected inputs = inputs;
}

@Directive({ selector, inputs })
export class CustomLayoutDirective extends LayoutDirective {
  protected inputs = inputs;
}
