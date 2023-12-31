import { get, writable } from "svelte/store";
import type { MeTag, TagOption } from "./interfaces";
import api, { processError } from "../utils/api";
import status from "../utils/status";

const { subscribe, set, update } = writable({} as { [key: number]: MeTag });

export default {
    subscribe,
    init: async (projectID: number) : Promise<boolean> => {
        const response = await api.get(`/v1/projects/${projectID}/tags`);

        if (response.status !== 200) {
            processError(response);
            return false;
        }

        const metags: MeTag[] = response.data;

        const tags: { [key: number]: MeTag } = {};

        metags.forEach((tag: MeTag) => {
            if(tag.options === null) tag.options = [];
            tags[tag.id] = tag;
        });

        set(tags);

        return true;
    },
    add: async (tag_id: number, value: string) => {
		const response = await api.post(`/v1/tags/${tag_id}/options`, {
			value
		});

		if (response.status !== status.Created) {
			processError(response, 'Failed to create tag option');
			return;
		}

        const option: TagOption = {
            id: response.data.id,
            tag_id,
            value,
        };

        update(tags => {
            tags[tag_id].options.push(option);
            return tags;
        });
    },
}