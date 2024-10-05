interface IMessage {
    id?: number;
    message: string;
    createdAt?: string | Date;
}

export interface IContact {
    id: number;
    avatar?: string;
    name: string;
    message?: IMessage
}