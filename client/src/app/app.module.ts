import {BrowserModule} from '@angular/platform-browser';
import {NgModule} from '@angular/core';

import {AppRoutingModule} from './app-routing.module';
import {AppComponent} from './app.component';
import {MainNavComponent} from './components/main-nav/main-nav.component';
import {
  MatGridListModule,
  MatCardModule,
  MatMenuModule,
  MatIconModule,
  MatButtonModule,
  MatTreeModule,
  MatToolbarModule,
  MatSidenavModule,
  MatListModule
} from '@angular/material';
import {LayoutModule} from '@angular/cdk/layout';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {TemplateModule} from './components/template/template.module';
import {FlexLayoutModule} from '@angular/flex-layout';
import {CUSTOM_HEIGHT_BREAKPOINTS} from './components/test/breakpoints';
import {CustomLayoutDirective} from './components/test/flex-width.directive';
import {CustomFlexOverrideDirective} from './components/test/flex-override.directive';
import {StoreModule} from '@ngrx/store';
import {reducers, metaReducers} from '../state/goals-store/reducers';
import {EffectsModule} from '@ngrx/effects';
import {environment} from '../environments/environment';
import {StoreDevtoolsModule} from '@ngrx/store-devtools';
import {GoalEffects} from '../state/goals-store/goal.effects';
import {HttpClientModule} from '@angular/common/http';

@NgModule({
  declarations: [
    AppComponent,
    // MainNavComponent,
    // NavTreeComponent,
    // MainDashboardComponent
    // CustomShowHideDirective
    CustomFlexOverrideDirective,
    CustomLayoutDirective
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    HttpClientModule,
    AppRoutingModule,
    MatGridListModule,
    MatCardModule,
    MatMenuModule,
    MatIconModule,
    MatButtonModule,
    LayoutModule,
    MatTreeModule,
    MatToolbarModule,
    MatSidenavModule,
    MatListModule,
    TemplateModule,
    /*FlexLayoutModule.withConfig({disableDefaultBps: true}, PRINT_BREAKPOINTS),*/
    FlexLayoutModule.withConfig({}, CUSTOM_HEIGHT_BREAKPOINTS),
    StoreModule.forRoot(reducers, {metaReducers}),
    !environment.production && StoreDevtoolsModule.instrument(),
    EffectsModule.forRoot([GoalEffects]),
    StoreDevtoolsModule.instrument({maxAge: 25, logOnly: environment.production})
  ],
  providers: [
    // CustomBreakPointsProvider
  ],
  bootstrap: [AppComponent]
})
export class AppModule {
}
