import cardsApi from '$lib/api/cardsApi';
import { get, writable } from 'svelte/store';
import CardTag from './CardTag';

const cards = writable([] as Card[]);

export default class Card {
	private _id: number;
	private _project_id: number;
	private _title: string;
	private _content: string;
	private _tags: CardTag[];

	private constructor(
		id: number,
		project_id: number,
		title: string,
		content: string,
		tags: CardTag[]
	) {
		this._id = id;
		this._project_id = project_id;
		this._title = title;
		this._content = content;
		this._tags = tags;
	}

	get id(): number {
		return this._id;
	}

	get project_id(): number {
		return this._project_id;
	}

	get title(): string {
		return this._title;
	}

	get content(): string {
		return this._content;
	}

	get tags(): CardTag[] {
		return this._tags;
	}

	static getAll(): Card[] {
		return get(cards);
	}

	static fromId(id: number): Card | null {
		for (const card of get(cards)) {
			if (card.id === id) {
				return card;
			}
		}

		return null;
	}

	static async create(project_id: number): Promise<Card | null> {
		const id = await cardsApi.create(project_id);

		if (!id) return null;

		const card = new Card(id, project_id, 'Untilted', '', []);

		cards.update((cards) => [...cards, card]);

		return card;
	}

	async delete(): Promise<boolean> {
		const res = await cardsApi.delete(this.id);

		if (!res) return false;

		cards.update((cards) => cards.filter((c) => c.id !== this.id));

		return true;
	}

	static parse(json: any): Card | null {
		if (json === null) {
			return null;
		}

		const card = new Card(json.id, json.project_id, json.title, json.content, []);

		card._tags = CardTag.parseAll(json.tags, card);

		cards.update((cards) => {
			if (!cards.find((c) => c.id === card.id)) {
				return [...cards, card];
			}

			return cards.map((c) => (c.id === card.id ? card : c));
		});

		return card;
	}

	static parseAll(json: any): Card[] {
		if (json === null) {
			return [];
		}

		const cards: Card[] = [];

		for (const jsonCard of json) {
			const card = this.parse(jsonCard);
			if (!card) continue;
			cards.push(card);
		}

		return cards;
	}
}
