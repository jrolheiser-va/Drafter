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
    this.playersService.players$.subscribe(
      players => {
        this.players = players;
        this.playoffPlayers = this.players.filter(
          player => PLAYOFF_TEAM_IDS.indexOf(player.TeamAbbreviation) > -1
        );
      }
    );
  }
}
