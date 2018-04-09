import { Component, OnInit } from '@angular/core';

import { Sort } from '@angular/material';
import { DataSource } from '@angular/cdk/table';

import { Player, PlayerService } from './player';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html'
})
export class AppComponent implements OnInit {

  displayedColumns = [
    'TeamAbbreviation',
    'FirstName',
    'LastName',
    'Points',
    'GamesPlayed',
    'PointsPerGame',
    'Goals',
    'Assists',
    'PlusMinus'
  ];
  playersSource: DataSource<Player>;

  constructor(private players: PlayerService) {}

  ngOnInit(): void {
    this.playersSource = {
      connect: () => this.players.players$,
      disconnect: () => null
    };
  }

  search(searchTerm: string): void {
    this.players.applySearch(searchTerm);
  }

  sortPlayers(sort: Sort): void {
    this.players.applySort(sort);
  }
}
