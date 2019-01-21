import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {TemplateComponent} from './template.component';
import {FlexLayoutModule, DEFAULT_BREAKPOINTS, BREAKPOINTS} from '@angular/flex-layout';
import {CustomBreakPointsProvider} from '../../utils/custom-breakpoints';
import {MainDashboardModule} from '../main-dashboard/main-dashboard.module';
import {MainNavModule} from '../main-nav/main-nav.module';
import {NavTreeModule} from '../nav-tree/nav-tree.module';

export const BreakPointsProvider = {
  provide: BREAKPOINTS,
  useValue: DEFAULT_BREAKPOINTS,
  multi: true
};

@NgModule({
  declarations: [
    TemplateComponent
  ],
  imports: [
    CommonModule,
    FlexLayoutModule,
    MainDashboardModule,
    MainNavModule,
    // NavTreeModule
  ],
  providers: [
    CustomBreakPointsProvider
  ],
  exports: [
    TemplateComponent
  ]
})
export class TemplateModule {
}
