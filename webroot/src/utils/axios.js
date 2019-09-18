import axios from 'axios'

axios.defaults.withCredentials = true;

// Add a request interceptor
axios.interceptors.request.use(function (config) {
  // Do something before request is sent
  return config
}, function (error) {
  // Do something with request error
  return Promise.reject(error)
});

// Add a response interceptor
axios.interceptors.response.use(function (res) {
  // Do something with response data
  return res
}, function (error) {
  // Do something with response error
  return Promise.reject(error)
});

export const http = axios;