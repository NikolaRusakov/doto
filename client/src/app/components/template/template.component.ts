import {Component, OnInit} from '@angular/core';
import {Store, select} from '@ngrx/store';
import * as goalAction from '../../../state/goals-store/goal.actions';
import {GoalSections} from '../../../api';
import * as fromStore from 'src/state/goals-store/reducers/index';
import {map} from 'rxjs/operators';

@Component({
  selector: 'dt-template',
  templateUrl: './template.component.html',
  styleUrls: ['./template.component.scss']
})
export class TemplateComponent implements OnInit {

  goals$;

  constructor(private store: Store<fromStore.State>) {
    this.goals$ = this.store.pipe(
      select(fromStore.selectAllGoals))
      .pipe(
        map((item) => {
          console.log(item);
          return item;
        }));
    this.store.pipe(select(fromStore.selectAllGoals)).pipe(
      map((entities) => {
        console.log(entities);
        return entities;
      }),
    );


  }

  ngOnInit() {
  }

  loadGoals(selected: GoalSections) {
    this.store.dispatch(new goalAction.LoadGoals({section: selected}));
  }

}
