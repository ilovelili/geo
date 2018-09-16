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
  }

  clear() {
    this.total = 0;
    this.steps = [];
  }

  format(step: Step): string {
    this.total += step.distance.value;
    return `${decodeURIComponent(step.html_instructions)}, ${step.distance.text}`;
  }
}
