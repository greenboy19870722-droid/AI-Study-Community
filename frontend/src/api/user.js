import request from "./request";

export function addUser(username) {
    return request({
        url: '/api/user/login',
        method: 'post',
        data: { username }
    })
}