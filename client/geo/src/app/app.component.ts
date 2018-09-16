import { Component } from '@angular/core';
import { GeoService } from './geo.service';
import { Step } from './geo.model';
import { RouteService } from './route.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  constructor(private geoService: GeoService, private routeService: RouteService) { }

  create(origin: string, destination: string, waypoints: string) {
    const waypointsarr = waypoints.split('|');
    this.geoService.getGeo(origin, destination, waypointsarr)
      .subscribe(
        route => route.forEach(r => this.addSteps(r.legs.map(l => l.steps)))
      );
  }

  /** Add steps */
  private addSteps(steps: Step[][]) {
    for (const i of steps) {
      for (const j of i) {
        this.routeService.add(j);
      }
    }
  }
}
