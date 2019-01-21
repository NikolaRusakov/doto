import { TestBed } from '@angular/core/testing';

import { GoalCrudService } from './goal-crud.service';

describe('GoalCrudService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: GoalCrudService = TestBed.get(GoalCrudService);
    expect(service).toBeTruthy();
  });
});
