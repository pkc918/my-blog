// axios config
export interface InterceptorMethodType {
    <T>(query: T): T | Promise<T>;
}

export interface APIResult<T> {
    code: number;
    data: T;
    msg: string;
    success: boolean;
    timestamp: number;
}
