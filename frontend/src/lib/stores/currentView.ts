import type View from '$lib/types/View';
import { writable } from 'svelte/store';

const { subscribe, set, update } = writable(null as View | null);

export default {
	subscribe,
	set
};
