import {Injectable} from '@angular/core';
import {GoalService} from '../../api';

@Injectable({
  providedIn: 'root'
})
export class GoalCrudService {

  constructor(private goalService: GoalService) {
  }

}
