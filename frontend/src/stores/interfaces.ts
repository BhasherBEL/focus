export interface Project {
    id: number | undefined;
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
    return cards.map(parseCard);
}