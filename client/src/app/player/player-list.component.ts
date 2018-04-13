import { Component, OnInit, Input, ViewChild } from '@angular/core';
import { Player } from './player';
import { DataSource } from '@angular/cdk/table';
import { MatTableDataSource, MatPaginator, MatSort } from '@angular/material';
import { PLAYOFF_TEAMS } from '../team/team';

@Component({
    selector: 'player-list',
    templateUrl: './player-list.component.html'
})

export class PlayerListComponent {
    @Input() teamAbbreviation: string; 
    @Input() players: Player[];

    displayedColumns = [
        //'TeamAbbreviation',
        'FirstName',
        'LastName',
        'Points',
        'GamesPlayed',
        'PointsPerGame',
        'Goals',
        'Assists'
    ];

    playersSource: MatTableDataSource<Player>;
    teamName = '';

    @ViewChild(MatPaginator) paginator: MatPaginator;
    @ViewChild(MatSort) sort: MatSort;

    constructor() {}

    // tslint:disable-next-line:use-life-cycle-interface
    ngOnInit() {
        this.playersSource = new MatTableDataSource(this.players);
        const team = PLAYOFF_TEAMS.find(t => t.ID === this.teamAbbreviation);
        this.teamName = team ? team.Name : this.teamAbbreviation;
    }

    // tslint:disable-next-line:use-life-cycle-interface
    ngAfterViewInit() {
        this.playersSource.paginator = this.paginator;
        this.playersSource.sort = this.sort;
    }

    applyFilter(filterValue: string) {
        filterValue = filterValue.trim();
        filterValue = filterValue.toLowerCase();
        this.playersSource.filter = filterValue;
    }
}
