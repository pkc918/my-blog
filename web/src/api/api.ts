import { get } from "@/request/request.ts";

const getProfile = <T>() => {
    return get<T>("/api/user");
};

export {
    getProfile
};
