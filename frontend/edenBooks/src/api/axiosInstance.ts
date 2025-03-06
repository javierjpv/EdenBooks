// axiosInstance.ts
import axios from 'axios';
import store, { RootState } from '../Store/store';


export const axiosInstance = axios.create({
  baseURL: 'http://localhost:6969', // Cambia esto a tu API base
});

// const {user}=useAuthStore()
axiosInstance.interceptors.request.use(
  (config) => {
    // const token=user.Token
    const token = (store.getState() as RootState).auth.user.Token;
    console.log("Token enviado:", token)
    if (token) {
      config.headers['Authorization'] = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);


