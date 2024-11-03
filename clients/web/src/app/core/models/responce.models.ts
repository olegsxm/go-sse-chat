import {IUser} from "./user.model";

export interface IAuthResponse {
    token: string;
    user: IUser;
}