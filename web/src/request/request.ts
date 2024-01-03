import axios, { AxiosInstance } from "axios";
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
    return new Promise((resolve, reject) => {
        request({
            url,
            data,
            method: "post",
        })
            .then((data) => {
                if (data.data && data.data.code != 10000) {
                    reject(data.data.message);
                } else {
                    resolve(data.data);
                }
            })
            .catch((reason) => {
                reject(reason);
            });
    });
};

const get = <T>(url: string, params?: any): Promise<APIResult<T>> => {
    return request({
        url,
        params,
        method: "get",
    });
};

export { request, get, post };
