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

    constructor(private http: HttpClient) {
        this.http.get('/api/player/list').subscribe(
            (players: Player[]) => this.allPlayers$$.next(players)
        );
    }

    get players$(): Observable<Player[]> {
        return this.allPlayers$;
    }

    private get allPlayers$(): Observable<Player[]> {
        return this.allPlayers$$.asObservable();
    }
}
