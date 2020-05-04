import {
  ActionReducerMap,
  createFeatureSelector,
  createSelector,
  MetaReducer
} from '@ngrx/store';
import {environment} from '../../../environments/environment';
import * as fromGoal from '../goal.reducer';

export interface State {
  goals: fromGoal.State;
}

export const reducers: ActionReducerMap<State> = {
  goals: fromGoal.reducer
};


export const metaReducers: MetaReducer<State>[] = !environment.production ? [] : [];

export const selectGoalsState = createFeatureSelector<fromGoal.State>('goals');

export const selectAllGoals = createSelector(
  selectGoalsState,
  fromGoal.selectAll
);
