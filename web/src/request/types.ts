// axios config
export interface InterceptorMethodType {
    <T>(query: T): T | Promise<T>;
}

export interface APIResult<T> {
    code: number;
    data: T;
    message: string;
    success: boolean;
    timestamp: number;
}
