import {HttpInterceptorFn, HttpResponse} from '@angular/common/http';
import {environment} from "../../../environments/environment";
import {of} from "rxjs";
import {IConversation} from "../models/conversation.model";

export const apiInterceptor: HttpInterceptorFn = (req, next) => {
    if (req.url.includes("/auth")) {
        const {login} = req.body as { login: string }
        return of(new HttpResponse({status: 200, body: {token: '4343434223 token', user: {id: 1, login: login}}}));
    }

    if (req.url.includes("conversations")) {
        const mocks: IConversation[] = [
            {
                "id": 78785,
                "avatar": "http://placeimg.com/640/480",
                "sender": {
                    "id": 42875,
                    "login": "4bA(lzWx?b",
                    "avatar": "http://placeimg.com/640/480"
                },
                "message": {
                    "id": 35878,
                    "message": "U5<n#\\=6jd",
                    "createdAt": "2024-05-23T15:33:08.918Z"
                }
            },
            {
                "id": 11739,
                "avatar": "http://placeimg.com/640/480",
                "sender": {
                    "id": 48895,
                    "login": "|Lk>7S=ktG",
                    "avatar": "http://placeimg.com/640/480"
                },
                "message": {
                    "id": 98267,
                    "message": "n/7tD,4>(|",
                    "createdAt": "2024-09-15T04:41:13.065Z"
                }
            },
            {
                "id": 29053,
                "avatar": "http://placeimg.com/640/480",
                "sender": {
                    "id": 33864,
                    "login": "?pw$94-,#m",
                    "avatar": "http://placeimg.com/640/480"
                },
                "message": {
                    "id": 57229,
                    "message": "ywQ*@<BxUN",
                    "createdAt": "2024-04-03T08:51:25.040Z"
                }
            },
            {
                "id": 89461,
                "avatar": "http://placeimg.com/640/480",
                "sender": {
                    "id": 13114,
                    "login": "3x/Sge%%2T",
                    "avatar": "http://placeimg.com/640/480"
                },
                "message": {
                    "id": 19120,
                    "message": "0C805RU'mr",
                    "createdAt": "2023-12-24T23:43:46.959Z"
                }
            },
            {
                "id": 18884,
                "avatar": "http://placeimg.com/640/480",
                "sender": {
                    "id": 85846,
                    "login": "3:Aiq|j}-H",
                    "avatar": "http://placeimg.com/640/480"
                },
                "message": {
                    "id": 37591,
                    "message": "N7sV.[4f6z",
                    "createdAt": "2024-04-18T15:40:10.708Z"
                }
            },
            {
                "id": 69628,
                "avatar": "http://placeimg.com/640/480",
                "sender": {
                    "id": 38218,
                    "login": "z]d6E@^k8D",
                    "avatar": "http://placeimg.com/640/480"
                },
                "message": {
                    "id": 44965,
                    "message": "_ClMGU/n=R",
                    "createdAt": "2024-03-14T14:17:30.413Z"
                }
            }
        ];

        return of(new HttpResponse({status: 200, body: mocks}));
    }

    return next(req.clone({
        url: `${environment.api}${req.url}`,
    }));
};
