import type Card from '$lib/types/Card';
import { writable } from 'svelte/store';

const { subscribe, set, update } = writable(null as Card | null);

export default {
	subscribe,
	set
};
