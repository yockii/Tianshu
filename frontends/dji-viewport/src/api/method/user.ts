import { type LoginResponse } from "@/types/user";
import { alovaInstance } from "../alova";

export const login = (loginRequest: {username: string, password: string, tenantName: string}) => {
    return alovaInstance.Post<LoginResponse>('/user/login',loginRequest);
}