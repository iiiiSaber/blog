import axios from 'axios'
import Vue from 'vue'

let Url='http://192.168.1.17:3131/interfaces/version1'

axios.defaults.baseURL = Url
axios.interceptors.request.use(config=>{
  config.headers.Authorization=`Bearer ${window.sessionStorage.getItem('token')}`
  return config
})
Vue.prototype.$http = axios
export { Url }