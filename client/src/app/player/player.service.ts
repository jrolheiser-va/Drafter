import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { Observable } from 'rxjs/Observable';
import { ReplaySubject } from 'rxjs/ReplaySubject';
import { BehaviorSubject } from 'rxjs/BehaviorSubject';
import { tap, map, combineLatest } from 'rxjs/operators';

import { Player } from './player';

import { MOCK_DATA } from './mock';
import { Sort, SortDirection } from '@angular/material';


@Injectable()
export class PlayerService {
    private allPlayers$$: ReplaySubject<Player[]> = new ReplaySubject(1);
    private searchTerm$$: BehaviorSubject<string> = new BehaviorSubject('');
    private sort$$: BehaviorSubject<Sort> = new BehaviorSubject({active: 'Points', direction: <SortDirection>'asc'});

    constructor(private http: HttpClient) {
        this.http.get('/api/player/list').subscribe(
            (players: Player[]) => this.allPlayers$$.next(players)
        );
    }

    get players$(): Observable<Player[]> {
        return Observable.combineLatest(
            this.allPlayers$,
            this.searchTerm$,
            this.sort$
        ).pipe(
            map(
                ([players, searchTerm, sort]) => players.filter(
                    player => player.FirstName.toLowerCase().indexOf(searchTerm.toLowerCase()) > -1 ||
                              player.LastName.toLowerCase().indexOf(searchTerm.toLowerCase()) > -1 ||
                              player.TeamAbbreviation.toLowerCase().indexOf(searchTerm.toLowerCase()) > -1
                ).sort(
                    (a, b) => {
                        if (sort.direction === <SortDirection>'desc') {
                            return a[sort.active] > b[sort.active] ? 1 : -1;
                        }
                        return a[sort.active] < b[sort.active] ? 1 : -1;
                    }
                )
            )
        );
    }

    private get allPlayers$(): Observable<Player[]> {
        return this.allPlayers$$.asObservable();
    }

    private get searchTerm$(): Observable<string> {
        return this.searchTerm$$.asObservable();
    }

    private get sort$(): Observable<Sort> {
        return this.sort$$.asObservable();
    }

    applySearch(searchTerm: string): void {
        this.searchTerm$$.next(searchTerm);
    }

    applySort(sort: Sort): void {
        this.sort$$.next(sort);
    }
}
