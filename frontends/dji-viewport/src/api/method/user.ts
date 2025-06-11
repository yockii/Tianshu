import { alovaInstance } from "../alova";

export const login = (username: string, password: string) => {
    return alovaInstance.Post('/login', {
        params: {
            username,
            password
        }
    });
}