import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { ISignInRequest, ISignInResponse, ISignUpRequest, ISignUpResponse } from '../../models/auth.models';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private readonly tokenLsKey = 'chat-token';
  private _token: string | null = localStorage.getItem(this.tokenLsKey);

  set token(token: string) {
    this._token = token;
    localStorage.setItem(this.tokenLsKey, token);
  }

  get token(): string | null {
    return this._token;
  }

  constructor(private http: HttpClient) {
  }


  signUp(form: ISignUpRequest): Observable<ISignUpResponse> {
    return this.http.post<ISignUpResponse>(`/auth/sign-up`, form);
  }

  signIn(form: ISignInRequest): Observable<ISignInResponse> {
    return this.http.post<ISignInResponse>(`/auth/sign-in`, form);
  }

}
