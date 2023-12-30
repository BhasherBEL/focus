import { toastAlert } from "../utils/toasts";

export interface Project {
    id: number;
    title: string;
}

export interface Card {
    id: number;
    project_id: number;
    title: string;
    content: string;
    tags: Tag[];
}

export interface Tag {
    card_id: number;
    tag_id: number;
    tag_title: string;
    value: string;
}

export function parseCard (c: any) {
    let card: Card = c;
    if (card.tags == null) card.tags = [];
    return card;
};

export function parseCards (cards: any) {
    if (cards == null) return [];

    let cardsArray;
    try {
       cardsArray = JSON.parse(cards);
    } catch (e) {
        toastAlert('Error', 'Could not parse cards');
        return [];
    }

    if (cardsArray == null) return [];

    return cardsArray.map(parseCard);
}