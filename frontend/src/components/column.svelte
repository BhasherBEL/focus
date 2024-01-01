<script lang="ts">
	import { updateCardTagApi } from '../api/cards';
	import type { Card, TagOption } from '../stores/interfaces';
	import { cards, currentDraggedCard } from '../stores/smallStore';
	import api, { processError } from '../utils/api';
	import status from '../utils/status';
	import CardC from './card.svelte';

	export let option: TagOption;
	export let columnCards: Card[] = [];

	async function onDrop(e: DragEvent) {
		e.preventDefault();
		if ($currentDraggedCard && $currentDraggedCard.tags) {
			for (let tag of $currentDraggedCard.tags) {
				if (tag.tag_id == option.tag_id) {
					try {
						if (tag.option_id == option.id) return;
						// DELETE
						if (tag.option_id !== -1 && option.id === -1) {
							const response = await api.delete(`/v1/cards/${tag.card_id}/tags/${tag.tag_id}`);

							if (response.status !== status.NoContent) {
								processError(response, 'Failed to delete tag');
								return;
							}
						}
						// CREATE
						else if (tag.option_id == -1 && option.id !== -1) {
							const response = await api.post(`/v1/cards/${tag.card_id}/tags/${tag.tag_id}`, {
								value: tag.value,
								option_id: option.id
							});
							if (response.status !== status.Created) {
								processError(response, 'Failed to create tag');
								return;
							}
						}
						// UPDATE
						else {
							const response = await api.put(`/v1/cards/${tag.card_id}/tags/${tag.tag_id}`, {
								value: tag.value,
								option_id: option.id
							});
							if (response.status !== status.NoContent) {
								processError(response, 'Failed to update tag');
								return;
							}
						}

						tag.option_id = option.id;
						cards.reload();
					} catch (e) {}
					break;
				}
			}
			currentDraggedCard.set(null);
		}
	}
</script>

<div
	class="column"
	role="listbox"
	tabindex="-1"
	on:drop={onDrop}
	on:dragover={(e) => {
		e.preventDefault();
	}}
>
	<h3>{option.value}</h3>
	<ul>
		{#each columnCards as card}
			<CardC {card} />
		{/each}
	</ul>
</div>

<style>
	h3 {
		text-align: center;
	}

	.column {
		margin: 0 10px;
		width: 200px;
		height: 100%;
	}
</style>
