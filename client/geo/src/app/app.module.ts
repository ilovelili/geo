import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule } from '@angular/forms';

import { AppComponent, SafePipe } from './app.component';
import { MessagesComponent } from './messages/messages.component';
import { RouteComponent } from './route/route.component';

@NgModule({
  declarations: [
    AppComponent,
    SafePipe,
    MessagesComponent,
    RouteComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    FormsModule,
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
