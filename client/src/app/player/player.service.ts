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
    private regularPlayers$$: ReplaySubject<Player[]> = new ReplaySubject(1);
    private playoffPlayers$$: ReplaySubject<Player[]> = new ReplaySubject(1);

    constructor(private http: HttpClient) {
        this.http.get('/api/player/list').subscribe(
            (players: Player[]) => this.regularPlayers$$.next(players)
        );
        this.http.get('/api/playoffplayer/list').subscribe(
            (players: Player[]) => this.playoffPlayers$$.next(players)
        );
    }

    get players$(): Observable<Player[]> {
        return this.regularPlayers$$.asObservable();
    }

    get playoffPlayers$(): Observable<Player[]> {
        return this.playoffPlayers$$.asObservable();
    }
}
