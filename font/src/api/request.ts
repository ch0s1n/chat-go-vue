import axios from "axios";

import { ElMessage } from "element-plus";
import { deleteLocalUser } from "@/utils/localutils";

const service = axios.create({
  baseURL: "/api",
  timeout: 20000, // 设置超时时间为20秒
});

service.interceptors.request.use(
  (config) => {
    config.headers["Authorization"] = localStorage.getItem("token");
    return config;
  },
  (error) => {
    ElMessage(error.message ?? "No response");
    // 对请求错误做些什么
    return Promise.reject(error);
  }
);

service.interceptors.response.use(
  (response) => {
    return response;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// 添加响应拦截器
service.interceptors.response.use(
  function (response) {
    // 2xx 范围内的状态码都会触发该函数。
    // 对响应数据做点什么
	ElMessage(response.data.message);
    return response;
  },
  function (error) {
    // 超出 2xx 范围的状态码都会触发该函数。
    // 对响应错误做点什么
    ElMessage(error.response.data.message);
    return Promise.reject(error);
  }
);

export default service;
