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

export class Route {
    summary: string;
    legs: Leg[];
}

export class Leg {
    steps: Step[];
    distance: Distance;
}

export class Step {
    html_instructions: string;
    distance: Distance;
    duration: number;
}

export class Distance {
    text: string;
    value: number;
}
