export class GeoRequest {
    origin: string;
    destination: string;
    waypoints: string[];

    constructor(origin: string, destination: string, waypoints: string[]) {
        this.origin = origin;
        this.destination = destination;
        this.waypoints = waypoints;
    }
}

export class GeoResponse {
    // tbd
}
