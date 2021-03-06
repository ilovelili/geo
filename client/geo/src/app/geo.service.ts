import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { GeoRequest, Route } from './geo.model';
import { MessageService } from './message.service';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { catchError, tap } from 'rxjs/operators';
import { environment } from '../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class GeoService {
  constructor(private http: HttpClient, private messageService: MessageService) { }

  getGeo(origin: string, destination: string, waypoints: string[]): Observable<Route[]> {
    const headers = new HttpHeaders({
      'Content-Type': 'application/json',
    });
    const options = { headers: headers };
    const request = new GeoRequest(origin, destination, waypoints);

    const r: Route[] = [];
    return this.http.post<Route[]>(this.resolveUrl(), request, options).pipe(
      tap(geo => this.log(`fetched geo ${geo}`)),
      catchError(this.handleError('getGeo', r))
    );
  }

  private handleError<T>(operation = 'operation', result?: T) {
    return (error: Error): Observable<T> => {
      console.error(error);
      this.log(`${operation} failed: ${error.message}`);
      return of(result as T);
    };
  }

  /** Log a GeoService message with the MessageService */
  private log(message: string) {
    this.messageService.add(`GeoService: ${message}`);
  }

  private resolveUrl(): string {
    if (environment.production) {
      return `http://188.166.244.244:3200/geo`;
    }
    return `http://0.0.0.0:3200/geo`;
  }
}
