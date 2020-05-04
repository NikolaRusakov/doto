import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {MainNavComponent} from './main-nav.component';
import {MatIconModule, MatSidenavModule, MatToolbarModule, MatListModule, MatButtonModule, MatRadioModule} from '@angular/material';
import {NgDependantModule} from '../../modules/ng-dependant/ng-dependant.module';
import {FormsModule} from '@angular/forms';

@NgModule({
  declarations: [MainNavComponent],
  imports: [
    CommonModule,
    NgDependantModule,
    FormsModule,
    MatSidenavModule,
    MatToolbarModule,
    MatListModule,
    MatIconModule,
    MatButtonModule,
    MatRadioModule
  ],
  exports: [MainNavComponent]
})
export class MainNavModule {
}
