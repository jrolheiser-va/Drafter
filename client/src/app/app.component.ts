import { Component, OnInit } from '@angular/core';

import { Sort } from '@angular/material';
import { DataSource } from '@angular/cdk/table';

import { Player, PlayerService } from './player';
import { Observable } from 'rxjs/Observable';
import { PLAYOFF_TEAM_IDS } from './team/team';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html'
})
export class AppComponent {

  players: Player[] = [];
  playoffPlayers: Player[] = [];

  constructor(private playersService: PlayerService) {
    Observable.combineLatest(
      this.playersService.players$,
      this.playersService.playoffPlayers$
    ).subscribe(
      ([players, playoffPlayers]) => {
        this.players = players;
        this.playoffPlayers = playoffPlayers;
      }
    );
  }
}
