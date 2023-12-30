import axios, { Axios, type AxiosResponse } from "axios";
import { backend } from "../stores/config";
import { toastAlert } from "./toasts";

export default new Axios({
    ...axios.defaults,
    baseURL: backend + '/api',
    validateStatus: () => true,
    headers: {
        'Content-Type': 'application/json'
    },
});

export function processError (response: AxiosResponse<any, any>, message: string = '') {
    let title = `${response.status} ${response.statusText}`;
    let subtitle = message;

    console.log(response.headers)

    if(response.headers["content-type"] === "application/json") {
        const parsed = response.data;
        subtitle = parsed.error;
        if(response.data.trace) {
            subtitle += '<br><br>' + parsed.trace;
        }
    } 

    axios.get(backend + '/api/trace');

    toastAlert(title, subtitle);
} 