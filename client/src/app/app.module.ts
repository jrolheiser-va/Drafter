import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule } from '@angular/forms';
import { ServiceWorkerModule } from '@angular/service-worker';
import { environment } from '../environments/environment';

import {
  MatToolbarModule,
  MatIconModule,
  MatTableModule,
  MatInputModule,
  MatSortModule,
  MatPaginatorModule,
  MatButtonToggleModule,
  MatGridListModule
} from '@angular/material';

import { AppComponent } from './app.component';
import { PlayerService } from './player';
import { PlayerListComponent } from './player/player-list.component';

@NgModule({
  declarations: [
    AppComponent, PlayerListComponent
  ],
  imports: [
    BrowserModule,
    environment.production ? ServiceWorkerModule.register('ngsw-worker.js') : [],
    BrowserAnimationsModule,
    HttpClientModule,
    FormsModule,
    MatToolbarModule,
    MatIconModule,
    MatTableModule,
    MatInputModule,
    MatSortModule,
    MatPaginatorModule,
    MatButtonToggleModule,
    MatGridListModule
  ],
  providers: [
    PlayerService
  ],
  bootstrap: [
    AppComponent
  ]
})
export class AppModule { }
