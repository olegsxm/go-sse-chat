import {Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {Observable} from "rxjs";
import {IAuthResponse} from "../models/responce.models";

@Injectable({
    providedIn: 'root'
})
export class AuthService {

    constructor(private http: HttpClient) {
    }

    signIn(login: string, password: string): Observable<IAuthResponse> {
        return this.http.post<IAuthResponse>('/auth/sign-in', {login, password},)
    }

    signUp(login: string, password: string): Observable<IAuthResponse> {
        return this.http.post<IAuthResponse>('/auth/sign-up', {login, password},)
    }
}
