import axios, { AxiosInstance, AxiosResponse } from "axios";
import {
    requestInterceptor,
    requestInterceptorError,
    responseInterceptor,
    responseInterceptorError,
    serverConfig
} from "@/request/config.ts";
import { APIResult } from "@/request/types.ts";

const request: AxiosInstance = axios.create({
    baseURL: serverConfig.baseURL,
    timeout: serverConfig.timeout,
    withCredentials: serverConfig.withCredentials,
});

request.interceptors.request.use(
    (config) => {
        return requestInterceptor<typeof config>(config);
    },
    (error) => {
        return requestInterceptorError<typeof error>(error);
    }
);
request.interceptors.response.use(
    (res) => {
        return responseInterceptor<typeof res>(res);
    },
    (error) => {
        return responseInterceptorError<typeof error>(error);
    }
);

const post = <T>(url: string, data?: any): Promise<APIResult<T>> => {
    return new Promise(async (resolve, reject) => {
        try {
            const res: AxiosResponse<APIResult<T>> = await request({
                url,
                data,
                method: "post",
            });
            if (res.data.code === 10000) {
                resolve(res.data);
            }
        } catch (e) {
            reject(e);
        }
    });
};

const get = <T>(url: string, params?: any): Promise<APIResult<T>> => {
    return new Promise(async (resolve, reject) => {
        try {
            const res: AxiosResponse<APIResult<T>> = await request({
                url,
                params,
                method: "get",
            });
            if (res.data.code === 10000) {
                resolve(res.data);
            }
        } catch (e) {
            reject(e);
        }
    });
};

export { request, get, post };
