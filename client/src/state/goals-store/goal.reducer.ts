import {Goal} from '../../api';
import {GoalActionsUnion, GoalActionTypes} from './goal.actions';
import {createEntityAdapter, EntityState, EntityAdapter} from '@ngrx/entity';

const adapter = createEntityAdapter<Goal>({
  selectId: goal => goal.id,
  sortComparer: (goalA, goalB) => goalA.name.localeCompare(goalB.name)
});

/*export interface State {
  goals: Goal[] | null;
  isLoading: boolean;
  error: string;
}*/
export type State = EntityState<Goal>;


/*export const initialState: State = {
  goals: [],
  isLoading: false,
  error: null
};*/
export const initialState: State = adapter.getInitialState();

export function reducer(
  state = initialState,
  action: GoalActionsUnion
): State {
  switch (action.type) {
    /*case GoalActionTypes.LOAD_GOALS: {
      /!* return {
         ...state,
         isLoading: true
       };*!/
    }*/
    case GoalActionTypes.LOAD_GOALS_SUCCESS: {
      return adapter.upsertMany(action.payload, state);
      /* return {
         ...state,
         isLoading: false,
         goals: [
           ...state.goals,
           ...action.payload
         ]
       };*/
    }
    /*case GoalActionTypes.LOAD_GOALS_FAILURE: {
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
    }*/

    default:
      return state;
  }
}

export const {
  selectIds,
  selectEntities,
  selectAll,
  selectTotal
} = adapter.getSelectors();
// export const getGoals = (state: State) => state.goals;
