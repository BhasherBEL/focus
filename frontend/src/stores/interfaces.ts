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
    tags: TagValue[];
}

export interface TagValue {
    card_id: number;
    tag_id: number;
    option_id: number;
    value: string;
}

export interface MeTag {
    id: number;
    project_id: number;
    title: string;
    type: number;
    options: TagOption[];
}

export interface TagOption {
    id: number;
    tag_id: number;
    value: string;
}

export interface View {
    id: number;
    project_id: number;
    primary_tag_id: number;
    secondary_tag_id: number;
    title: string;
}

export function parseCard (c: any) {
    let card: Card = c;
if (card.tags == null) card.tags = [];
    return card;
};

export function parseCards (cards: any) {
    if (cards == null) return [];
    return cards.map((c: any) => parseCard(c));
}

export function parseMeTag (t: any) {
    let tag: MeTag = t;
    if (tag.options == null) tag.options = [];
    return tag;
}

export function parseMeTags (tags: any) {
    if (tags == null) return {};
    return tags.map(parseMeTag).reduce((acc: any, tag: MeTag) => {
        acc[tag.id] = tag;
        return acc;
    });
}