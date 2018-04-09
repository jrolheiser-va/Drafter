import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule } from '@angular/forms';

import {
  MatToolbarModule,
  MatIconModule,
  MatTableModule,
  MatInputModule,
  MatSortModule
} from '@angular/material';

import { AppComponent } from './app.component';
import { PlayerService } from './player';

@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    HttpClientModule,
    FormsModule,
    MatToolbarModule,
    MatIconModule,
    MatTableModule,
    MatInputModule,
    MatSortModule
  ],
  providers: [
    PlayerService
  ],
  bootstrap: [
    AppComponent
  ]
})
export class AppModule { }
