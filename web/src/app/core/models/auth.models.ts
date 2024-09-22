export interface ISignInRequest {
  login: string;
  password: string;
}

export interface ISignInResponse {
  token: string;
}

export interface ISignUpRequest {
  login: string;
  password: string;
}

export interface ISignUpResponse {
  token: string;
}