import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {MainDashboardComponent} from './main-dashboard.component';
import {MatListModule, MatGridListModule, MatCardModule, MatMenuModule, MatIconModule} from '@angular/material';
import {NgDependantModule} from '../../modules/ng-dependant/ng-dependant.module';

@NgModule({
  declarations: [MainDashboardComponent],
  imports: [
    CommonModule,
    NgDependantModule,
    MatGridListModule,
    MatCardModule,
    MatMenuModule,
    MatIconModule,
    MatListModule
  ], exports: [MainDashboardComponent]
})
export class MainDashboardModule {
}
