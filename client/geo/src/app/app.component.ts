import { Component, ViewEncapsulation, ViewChild, ElementRef, PipeTransform, Pipe, OnInit } from '@angular/core';
import { GeoService } from './geo.service';
import { Step } from './geo.model';
import { RouteService } from './route.service';
import { DomSanitizer } from '@angular/platform-browser';

@Pipe({ name: 'safe' })
export class SafePipe implements PipeTransform {
  constructor(private sanitizer: DomSanitizer) { }
  transform(url) {
    return this.sanitizer.bypassSecurityTrustResourceUrl(url);
  }
}

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  src = '';
  constructor(private geoService: GeoService, private routeService: RouteService) { }

  create(origin: string, destination: string, waypoints: string) {
    this.src = `https://www.google.com/maps/embed/v1/directions?key=AIzaSyDtR1i_vcnrR-MhsNUBTxNGLKneEPQtrfI&origin=${origin}&destination=${destination}&waypoints=${waypoints}`;
    const waypointsarr = waypoints.split('|');
    this.routeService.clear();
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
