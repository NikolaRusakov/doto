import {Goal} from '../../api';
import {GoalActionsUnion, GoalActionTypes} from './goal.actions';

export interface State {
  goals: Goal[] | null;
  isLoading: boolean;
  error: string;
}

export const initialState: State = {
  goals: [],
  isLoading: false,
  error: null
};

export function reducer(
  state = initialState,
  action: GoalActionsUnion
): State {
  switch (action.type) {
    case GoalActionTypes.LOAD_GOALS: {
      return {
        ...state,
        isLoading: true
      };
    }
    case GoalActionTypes.LOAD_GOALS_SUCCESS: {
      return {
        ...state,
        isLoading: false,
        goals: [...action.payload.data]
      };
    }
    case GoalActionTypes.LOAD_GOALS_FAILURE: {
      return {
        ...state,
        isLoading: false,
        error: action.payload.error
      };
    }
    case GoalActionTypes.UPDATE_GOALS: {
      return {
        ...state,
        isLoading: true
      };
    }
    case GoalActionTypes.UPDATE_GOALS_SUCCESS: {
      return {
        ...state,
        isLoading: false,
        goals: [...action.payload.data]
      };
    }
    case GoalActionTypes.UPDATE_GOALS_FAILURE: {
      return {
        ...state,
        isLoading: false,
        error: action.payload.error
      };
    }
    case GoalActionTypes.REMOVE_GOALS: {
      return {
        ...state,
        isLoading: true
      };
    }
    case GoalActionTypes.REMOVE_GOALS_SUCCESS: {
      return {
        ...state,
        isLoading: false,
      };
    }
    case GoalActionTypes.REMOVE_GOALS_FAILURE: {
      return {
        ...state,
        isLoading: false,
        error: action.payload.error
      };
    }

    default:
      return state;
  }
}
