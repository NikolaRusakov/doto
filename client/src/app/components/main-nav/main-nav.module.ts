import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {MainNavComponent} from './main-nav.component';
import {MatIconModule, MatSidenavModule, MatToolbarModule, MatListModule} from '@angular/material';
import {NgDependantModule} from '../../modules/ng-dependant/ng-dependant.module';

@NgModule({
  declarations: [MainNavComponent],
  imports: [
    CommonModule,
    NgDependantModule,
    MatSidenavModule,
    MatToolbarModule,
    MatListModule,
    MatIconModule
  ],
  exports: [MainNavComponent]
})
export class MainNavModule {
}
