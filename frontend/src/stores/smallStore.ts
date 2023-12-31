import { writable } from 'svelte/store';
import { parseCards, type Card, type View } from './interfaces';
import { deleteCardApi, newCardApi } from '../api/cards';
import { getProjectCardsAPI } from '../api/projects';

export const currentView = writable(null as View | null);

export const currentModalCard = writable(-1);

export const cards = (() => {
	const { subscribe, set, update } = writable([] as Card[]);

	return {
		subscribe,
		init: async (projectId: number) => {
			getProjectCardsAPI(projectId).then((c) => {
				set(parseCards(c));
			});
		},
		add: async (projectId: number) => {
			await newCardApi(projectId).then((card) => {
				currentModalCard.set(card.id);
				update((cards) => [...cards, card]);
			});
		},
		remove: async (card: Card) => {
			await deleteCardApi(card.id).then(() => {
				update((cards) => cards.filter((c) => c.id !== card.id));
				currentModalCard.set(-1);
			});
		}
	};
})();
