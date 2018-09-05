import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { GeoRequest, GeoResponse } from './geo.model';
import { MessageService } from './message.service';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { catchError, tap } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class GeoService {
  // move to config
  static url = 'http:localhost:3000/geo';

  constructor(private http: HttpClient, private messageService: MessageService) { }

  getGeo(origin: string, destination: string, waypoints: string[]): Observable<GeoResponse> {
    const headers = new HttpHeaders({
      'Content-Type': 'application/json',
    });
    const options = { headers: headers };
    const request = new GeoRequest(origin, destination, waypoints);

    return this.http.post<GeoResponse>(GeoService.url, request, options).pipe(
      tap(geo => this.log(`fetched geo ${geo}`)),
      catchError(this.handleError('getGeo', []))
    );
  }

  private handleError<T>(operation = 'operation', result?: T) {
    return (error: any): Observable<T> => {
      console.error(error);
      this.log(`${operation} failed: ${error.message}`);
      return of(result as T);
    };
  }

  /** Log a GeoService message with the MessageService */
  private log(message: string) {
    this.messageService.add(`GeoService: ${message}`);
  }
}
