import request from '../utils/request';

export const HOST = 'http://127.0.0.1:9000'

export const login = data => {
    return request({
        url: HOST + '/login',
        method: 'post',
        data: data
    });
}

export const fetchData = query => {
    return request({
        url: './table.json',
        method: 'get',
        params: query
    });
};
