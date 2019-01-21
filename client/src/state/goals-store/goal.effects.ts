import {Injectable} from '@angular/core';
import {Actions, Effect, ofType} from '@ngrx/effects';
import {Observable, of} from 'rxjs';
import {Action} from '@ngrx/store';
import {mergeMap, map, catchError} from 'rxjs/operators';
import {HttpClient} from '@angular/common/http';
import {GoalService} from '../../api';
import {GoalActionTypes, LoadGoals} from './goal.actions';


@Injectable()
export class GoalEffects {
  @Effect()
  goalsBySection$: Observable<Action> = this.actions$.pipe(
    ofType<LoadGoals>(
      GoalActionTypes.LOAD_GOALS
    ),
    mergeMap(action =>
      this.goalService.goalsSection(action.payload.section).pipe(
        map(data => ({type: GoalActionTypes.LOAD_GOALS_SUCCESS, payload: data})),
        catchError(() => of({type: GoalActionTypes.LOAD_GOALS_FAILURE}))
      )
    )
  );

  constructor(private http: HttpClient,
              private goalService: GoalService,
              private actions$: Actions
  ) {
  }
}
