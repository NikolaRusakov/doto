import { TestBed, inject } from '@angular/core/testing';
import { provideMockActions } from '@ngrx/effects/testing';
import { Observable } from 'rxjs';

import { GoalEffects } from './goal.effects';

describe('GoalEffects', () => {
  let actions$: Observable<any>;
  let effects: GoalEffects;

  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [
        GoalEffects,
        provideMockActions(() => actions$)
      ]
    });

    effects = TestBed.get(GoalEffects);
  });

  it('should be created', () => {
    expect(effects).toBeTruthy();
  });
});
