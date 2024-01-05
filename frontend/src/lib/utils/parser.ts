import type Card from '../types/Card';
import type MeTag from '../types/MeTag';

export function parseCard(c: any) {
	let card: Card = c;
	if (card.tags == null) card.tags = [];
	return card;
}

export function parseCards(cards: any) {
	if (cards == null) return [];
	return cards.map((c: any) => parseCard(c));
}

export function parseMeTag(t: any) {
	let tag: MeTag = t;
	if (tag.options == null) tag.options = [];
	return tag;
}

export function parseMeTags(tags: any) {
	if (tags == null) return {};
	return tags.map(parseMeTag).reduce((acc: any, tag: MeTag) => {
		acc[tag.id] = tag;
		return acc;
	});
}
