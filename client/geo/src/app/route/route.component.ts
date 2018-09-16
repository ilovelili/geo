import { Component, OnInit } from '@angular/core';
import { RouteService } from '../route.service';

@Component({
  selector: 'app-route',
  templateUrl: './route.component.html',
  styleUrls: ['./route.component.css']
})
export class RouteComponent implements OnInit {
  constructor(public routeService: RouteService) { }

  ngOnInit() {
  }
}
