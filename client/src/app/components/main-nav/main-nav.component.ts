import {Component, Output, EventEmitter, Input} from '@angular/core';
import {BreakpointObserver, Breakpoints} from '@angular/cdk/layout';
import {Observable} from 'rxjs';
import {map} from 'rxjs/operators';
import {Goal, GoalSections} from '../../../api';

@Component({
  selector: 'dt-main-nav',
  templateUrl: './main-nav.component.html',
  styleUrls: ['./main-nav.component.scss']
})
export class MainNavComponent {
  @Input()
  goals: Goal[];
  @Output()
  getGoals = new EventEmitter<GoalSections>();

  selectedSection: GoalSections;
  goalSections: GoalSections[] = Object.keys(GoalSections)
    .reduce((previousValue, currentValue) => [...previousValue, currentValue.toLowerCase()], []);
  isHandset$: Observable<boolean> = this.breakpointObserver.observe(Breakpoints.Handset)
    .pipe(
      map(result => result.matches)
    );

  constructor(private breakpointObserver: BreakpointObserver) {

  }


}
