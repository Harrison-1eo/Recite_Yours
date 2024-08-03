import axios from 'axios';
import { Toast } from 'bootstrap';
import {showErrorToast} from "@/utils/showMessage";
import router from "@/router/router"; // 如果你使用 Bootstrap 的 Toast 组件来显示消息

// 创建一个 Axios 实例
const axios_instance = axios.create({
    // baseURL: 'http://47.92.84.147:8010', // 替换为你的后端接口基础 URL
    baseURL: 'http://localhost:8010', // 替换为你的后端接口基础 URL
    timeout: 10000, // 请求超时时间（毫秒）
    headers: { 'Content-Type': 'application/json' }
});


// 请求拦截器
// 添加Token
axios_instance.interceptors.request.use(
    config => {
        // 从 localStorage 中获取 Token
        const token = localStorage.getItem('token');

        // 如果 Token 存在，添加到请求头
        if (token) {
            config.headers.Authorization = `Bearer ${token}`;
        }

        return config;
    },
    error => {
        return Promise.reject(error);
    }
);

// 响应拦截器
axios_instance.interceptors.response.use(
    response => {
        const { code, message, data } = response.data;

        // 处理成功响应
        if (code === 0) {
            return data; // 直接返回数据部分
        } else if (code === 2) {
            showErrorToast(message);
            // 重定向到登录页
            router.push('/login');
        } else {
            // 处理业务逻辑错误
            showErrorToast(message);

            console.log(response.data);
            return Promise.reject(response.data); // 拒绝 Promise
        }
    },
    error => {
        // 处理 HTTP 错误
        if (error.response) {
            const status = error.response.status;
            let errorMessage = 'An error occurred';

            if (status === 404) {
                errorMessage = 'Resource not found';
            } else if (status === 500) {
                errorMessage = 'Server error';
            }

            showErrorToast(errorMessage);
        } else {
            showErrorToast('Network error');
        }

        console.log(error);
        return Promise.reject(error); // 拒绝 Promise
    }
);

// 统一定义后端接口
export const API = {
    LOGIN: '/login',
    REGISTER: '/register',
    SOURCES_LIST: '/api/sources/list',
    SOURCES_ADD: '/api/sources/add',
    WORDS_ADD: '/api/words/add',
    WORDS_AUTOCOMPLETE: '/api/words/autocomplete',
    WORDS_MEANINGS: '/api/words/meaning',
    WORDS_LIST: '/api/words/list',
    // 添加其他接口
};

// 导出 Axios 实例
export default axios_instance;
