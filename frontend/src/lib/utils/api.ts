import axios, { Axios, type AxiosResponse } from 'axios';
import { toastAlert } from './toasts';
import { setupCache } from 'axios-cache-interceptor';
import { env } from '$env/dynamic/public';

const backend = env.PUBLIC_BACKEND_URL || 'http://localhost:3000';

export default setupCache(
	new Axios({
		...axios.defaults,
		baseURL: backend + '/api',
		validateStatus: () => true,
		headers: {
			'Content-Type': 'application/json'
		}
	}),
	{
		interpretHeader: true,
		modifiedSince: true
	}
);

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
