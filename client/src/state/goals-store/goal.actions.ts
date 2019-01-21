import {Action} from '@ngrx/store';
import {Goal, GoalSections} from '../../api';

export enum GoalActionTypes {
  LOAD_GOALS = '[GOAL] Load Goals',
  LOAD_GOALS_FAILURE = '[GOAL] Load Goals Failure',
  UPDATE_GOALS = '[GOAL] Update Goals',
  UPDATE_GOALS_FAILURE = '[GOAL] Update Goals Failure',
  UPDATE_GOALS_SUCCESS = '[GOAL] Update Goals Success',
  REMOVE_GOALS = '[GOAL] Remove Goals',
  REMOVE_GOALS_FAILURE = '[GOAL] Remove Goals Failure',
  REMOVE_GOALS_SUCCESS = '[GOAL] Remove Goals Success',
  LOAD_GOALS_SUCCESS = '[GOAL] Load Goals Success'
}

export class LoadGoals implements Action {
  readonly type = GoalActionTypes.LOAD_GOALS;

  constructor(public payload: { section: GoalSections }) {
  }
}

export class LoadGoalsSuccess implements Action {
  readonly type = GoalActionTypes.LOAD_GOALS_SUCCESS;

  constructor(public payload: { data: Goal[] }) {
  }
}

export class LoadGoalsFailure implements Action {
  readonly type = GoalActionTypes.LOAD_GOALS_FAILURE;

  constructor(public payload: { error: string }) {
  }
}

export class UpdateGoals implements Action {
  readonly type = GoalActionTypes.UPDATE_GOALS;

  constructor(public payload: { message: string }) {
  }
}

export class UpdateGoalsSuccess implements Action {
  readonly type = GoalActionTypes.UPDATE_GOALS_SUCCESS;

  constructor(public payload: { data: Goal[] }) {
  }
}

export class UpdateGoalsFailure implements Action {
  readonly type = GoalActionTypes.UPDATE_GOALS_FAILURE;

  constructor(public payload: { error: string }) {
  }
}


export class RemoveGoals implements Action {
  readonly type = GoalActionTypes.REMOVE_GOALS;

  constructor(public payload: { message: string }) {
  }
}

export class RemoveGoalsSuccess implements Action {
  readonly type = GoalActionTypes.REMOVE_GOALS_SUCCESS;

  constructor(public payload: { onRemove: string }) {
  }
}

export class RemoveGoalsFailure implements Action {
  readonly type = GoalActionTypes.REMOVE_GOALS_FAILURE;

  constructor(public payload: { error: string }) {
  }
}

export type GoalActionsUnion =
  LoadGoals
  | LoadGoalsSuccess
  | LoadGoalsFailure
  | RemoveGoals
  | RemoveGoalsSuccess
  | RemoveGoalsFailure
  | UpdateGoals
  | UpdateGoalsSuccess
  | UpdateGoalsFailure;
