import axios, { Axios, type AxiosResponse } from 'axios';
import { setupCache } from 'axios-cache-interceptor';
import { toastAlert } from './toasts';
// import { env } from '$env/dynamic/public';

// const backend = env.PUBLIC_BACKEND_URL || 'http://localhost:3000';
const backendUrl = 'http://localhost:3000';

let pendingRequests = 0;

export function hasPendingRequests() {
	return pendingRequests > 0;
}

const axiosInstance = new Axios({
	...axios.defaults,
	baseURL: backendUrl + '/api',
	validateStatus: () => true,
	headers: {
		'Content-Type': 'application/json'
	}
});

const cachedInstance = setupCache(axiosInstance, {
	interpretHeader: true,
	modifiedSince: true
});

axiosInstance.interceptors.request.use(
	(config) => {
		pendingRequests++;
		return config;
	},
	(error) => {
		pendingRequests--;
		return Promise.reject(error);
	}
);

axiosInstance.interceptors.response.use(
	(response) => {
		pendingRequests--;
		return response;
	},
	(error) => {
		pendingRequests--;
		return Promise.reject(error);
	}
);

export default axiosInstance;

export function processError(response: AxiosResponse<any, any>, message: string = '') {
	let title = `${response.status} ${response.statusText}`;
	let subtitle = message;

	if (response.headers['content-type'] === 'application/json') {
		const parsed = response.data;
		subtitle = parsed.error;
		if (response.data.trace) {
			subtitle += '<br><br>' + parsed.trace;
		}
	}

	toastAlert(title, subtitle);
}
