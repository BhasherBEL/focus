import { writable } from 'svelte/store';

const { subscribe, set, update } = writable(null as number | null);

export default {
	subscribe,
	set
};
