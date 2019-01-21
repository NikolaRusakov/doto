import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import {MatGridListModule, MatCardModule, MatIconModule, MatMenuModule, MatTreeModule, MatButtonModule, MatSidenavModule, MatListModule,
  MatToolbarModule} from '@angular/material';
import {TemplateModule} from '../../components/template/template.module';

@NgModule({
  declarations: [],
  imports: [
   /* CommonModule,
    MatGridListModule,
    MatCardModule,
    MatMenuModule,
    MatIconModule,
    MatButtonModule,
    MatTreeModule,
    MatToolbarModule,
    MatSidenavModule,
    MatListModule,*/
    MatGridListModule,
    MatCardModule,
    MatMenuModule,
    MatIconModule,
    MatButtonModule,
    MatTreeModule,
    MatToolbarModule,
    MatSidenavModule,
    MatListModule,
  ]
})
export class NgDependantModule { }
