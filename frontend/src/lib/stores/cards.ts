import { deleteCardApi, newCardApi, updateCardApi } from '$lib/api/cards';
import { getProjectCardsAPI } from '$lib/api/projects';
import type Card from '$lib/types/Card';
import type TagValue from '$lib/types/TagValue';
import { writable } from 'svelte/store';
import { parseCards } from '../utils/parser';
import currentModalCard from './currentModalCard';

const { subscribe, set, update } = writable([] as Card[]);

async function init(projectId: number) {
	getProjectCardsAPI(projectId).then((c) => {
		set(parseCards(c));
	});
}

async function add(projectId: number, tags: TagValue[]) {
	await newCardApi(projectId, tags).then((card) => {
		update((cards) => [...cards, card]);
		currentModalCard.set(card.id);
	});
}

async function remove(card: Card) {
	await deleteCardApi(card.id).then(() => {
		update((cards) => cards.filter((c) => c.id !== card.id));
		currentModalCard.set(null);
	});
}

async function edit(card: Card): Promise<boolean> {
	if (await updateCardApi(card)) {
		update((cards) => cards.map((c) => (c.id === card.id ? card : c)));
		return true;
	}
	return false;
}

/**
 * @deprecated The method should not be used. It is only used to force a reload of the cards.
 */
function reload() {
	update((cards) => cards);
}

export default {
	subscribe,
	init,
	add,
	remove,
	edit,
	reload
};
