import axios from 'axios'
import type {AxiosInstance} from 'axios'

const api:AxiosInstance = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL,
  headers: {
    'Content-Type': 'application/json'
  }
})

// Optional: Add request interceptor to attach auth token
// api.interceptors.request.use(
//     (config) => {
//       const token = localStorage.getItem('auth_token');
//       if (token) {
//         config.headers.Authorization = `Bearer ${token}`;
//       }
//       return config;
//     },
//     (error) => Promise.reject(error)
//   );
  
//   // Optional: Response interceptor for global error handling
//   api.interceptors.response.use(
//     (response) => response,
//     (error: AxiosError) => {
//       // Log or handle specific status codes globally
//       console.error('API Error:', error.response?.status, error.message);
//       return Promise.reject(error);
//     }
//   );

export default api