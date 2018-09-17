import { Injectable } from '@angular/core';
import { Step } from './geo.model';

@Injectable({
  providedIn: 'root'
})
export class RouteService {
  steps: Step[] = [];
  total = 0;
  add(step: Step) {
    this.steps.push(step);
    this.total += step.distance.value;
  }

  clear() {
    this.total = 0;
    this.steps = [];
  }

  format(step: Step): string {
    return `${decodeURIComponent(step.html_instructions)}, ${step.distance.text}`;
  }
}
