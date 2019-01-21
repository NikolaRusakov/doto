import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {NavTreeComponent} from './nav-tree.component';
import {MatTreeModule} from '@angular/material';
import {NgDependantModule} from 'src/app/modules/ng-dependant/ng-dependant.module';

@NgModule({
  declarations: [NavTreeComponent],
  imports: [
    CommonModule,
    NgDependantModule
  ],
  exports: [NavTreeComponent]
})
export class NavTreeModule {
}
