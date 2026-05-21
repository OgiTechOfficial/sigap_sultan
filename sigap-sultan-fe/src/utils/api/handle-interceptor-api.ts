import { AxiosStatic } from 'axios';
import { USER_STORAGE_KEY, ACCESS_TOKEN_STORAGE_KEY, REFRESH_TOKEN_STORAGE_KEY } from '@/constants/storage';
import { UserAccount } from '@/types/user-account';
import { createResourceId } from '../create-resource-id';

function handleInterceptorAPI(axios: AxiosStatic) {

  axios.interceptors.request.use(
    (config) => {
      const tokenStorage = window.sessionStorage.getItem(ACCESS_TOKEN_STORAGE_KEY);
      const userStorage = window.sessionStorage.getItem(USER_STORAGE_KEY);
      const user: UserAccount | null = userStorage ? JSON.parse(userStorage) : null;
      // TODO: adding condition for public API
      if (tokenStorage) {
        config.headers['Authorization'] = `Bearer ${tokenStorage}`;
        config.headers['Accept-Language'] = window.navigator.language;
        config.headers['User-Agent'] = window.navigator.userAgent;
        config.headers['X-Service-Id'] = 'gateway';
        config.headers['Accept'] = 'application/json, text/plain'
        // if (user?.accountId) {
        //   config.headers['X-Account-Id '] = user.accountId;
        // }
        // config.headers['X-Request-Id'] = createResourceId();
        // if (user?.business) {
        //   config.headers['X-Business-Id'] = user.business.businessId;
        // }

      }
      return config;
    },
    (error) => {
      return Promise.reject(error);
    }
  );

  axios.interceptors.response.use(
    (res) => {
      return res;
    },
    async (err) => {
      const originalConfig = err.config;
      if (err.response) {
        // Access Token was expired
        if (err.response.status === 401 && !originalConfig._retry) {
          originalConfig._retry = true;
          try {
            // TODO: implement refresh token
            // const response = await sessionApi.refresh();
            // const { result, error } = response;
            // if (result && !error) {
            //   sessionStorage.setItem(ACCESS_TOKEN_STORAGE_KEY, result.accessToken);
            //   sessionStorage.setItem(REFRESH_TOKEN_STORAGE_KEY, result.refreshToken);
            //   sessionStorage.setItem(USER_STORAGE_KEY, JSON.stringify(result));
            // }

            return axios(originalConfig);
          } catch (_error) {
            return Promise.reject(_error);
          }
        }
      }
      return Promise.reject(err);
    }
  )

  return axios;
}

export default handleInterceptorAPI;