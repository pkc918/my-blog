import { InterceptorMethodType } from "./types";

export const serverConfig = {
    baseURL: import.meta.env.VITE_BASE_URL,
    timeout: 30000,
    withCredentials: true,
    useTokenAuthorization: false,
};

// 请求拦截器
export const requestInterceptor: InterceptorMethodType = (config: any) => {
    return config;
};

// 请求错误拦截器
export const requestInterceptorError: InterceptorMethodType = (err: any) => {
    return Promise.reject(err);
};

// 响应拦截器
export const responseInterceptor: InterceptorMethodType = (res: any) => {
    return res;
};

// 响应错误拦截器
export const responseInterceptorError: InterceptorMethodType = (error: any) => {
    return Promise.reject(error);
};
