import {createAlova} from 'alova';
import adapterFetch  from 'alova/fetch';
import vueHook from 'alova/vue';
import LZString from 'lz-string';

export const alovaInstance = createAlova({
    requestAdapter: adapterFetch(),
    baseURL: import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api/v1',
    statesHook: vueHook,
    beforeRequest: (method) => {
        if(method.data && !(method.data instanceof FormData)) {
            // Compress the request data using LZString
            const compressedData = LZString.compressToUTF16(JSON.stringify(method.data));
            method.data = compressedData;
        }

        if (!method.meta?.ignoreToken) {
            const token = localStorage.getItem('token');
            if (token) {
                method.config.headers.Authorization = `Bearer ${token}`;
            }
        }
    },
    responded: {
        onSuccess: async (response, method) => {
            if (response.status >= 400) {
                throw new Error(response.statusText)
            }

            const data = await response.text()

            const json = JSON.parse(LZString.decompressFromUTF16(data))
            if (json.code !== 200) {
                throw new Error(json.message || 'Request failed')
            }

            return json.data;
        },
        onError: (err, method) => {
            console.error('Request failed:', err);
        },
        onComplete: async method => {

        }
    },
})