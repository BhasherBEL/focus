<script lang="ts">
	import type { TagOption, Card, MeTag, TagValue } from '../../stores/interfaces';
	import { cards, currentDraggedCard } from '../../stores/smallStore';
	import api, { processError } from '../../utils/api';
	import status from '../../utils/status';
	import CardC from './card/card.svelte';
	import AddIcon from '../icons/addIcon.svelte';
	import projectTags from '../../stores/projectTags';
	import { updateTagAPI as updateTagOptionAPI } from '../../api/tags';
	import { get } from 'svelte/store';

	export let projectId: number;
	export let editable: boolean = true;
	export let option: TagOption;
	export let columnCards: Card[] = [];

	let lastOptionValue = option.value;

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

	async function addCard() {
		const tags: TagValue[] = [];
		for (let tag of Object.values(get(projectTags))) {
			tags.push({
				card_id: -1,
				tag_id: tag.id,
				option_id: tag.id === option.tag_id ? option.id : -1,
				value: ''
			});
		}

		await cards.add(projectId, tags);
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
	<header>
		<input
			bind:value={option.value}
			type="text"
			on:blur={async () => {
				if (lastOptionValue === option.value) return;
				await updateTagOptionAPI(option);
				lastOptionValue = option.value;
				cards.reload();
			}}
			disabled={!editable}
		/>
		<span>
			<span>{columnCards.length}</span>
			<span
				class="add"
				on:click={addCard}
				role="button"
				tabindex="0"
				on:keypress={(e) => {
					if (e.key === 'Enter') {
						addCard();
					}
				}}
			>
				<AddIcon />
			</span>
		</span>
	</header>
	<ul>
		{#each columnCards as card}
			<CardC {card} />
		{/each}
	</ul>
	<div
		class="addEnd"
		on:click={addCard}
		role="button"
		tabindex="0"
		on:keypress={(e) => {
			if (e.key === 'Enter') {
				addCard();
			}
		}}
	>
		+
	</div>
</div>

<style lang="less">
	.column {
		margin: 20px 10px;
		width: 250px;
	}

	header {
		margin-bottom: 20px;
		display: flex;
		flex-direction: row;
		justify-content: space-between;

		input {
			background-color: #444;
			flex: 1;
			margin-right: 10px;
			padding: 3px 10px;
			border-radius: 3px;
			overflow: hidden;
			color: inherit;
			border: none;
		}

		:last-child {
			display: flex;
		}

		.add {
			cursor: pointer;
			border-radius: 4px;
			width: 20px;
			height: 20px;
			display: flex;
			justify-content: center;
			align-items: center;
			margin-left: 10px;

			&:hover {
				background-color: #fff1;
			}
		}
	}

	.addEnd {
		cursor: pointer;
		border-radius: 4px;
		height: 40px;
		font-size: 150%;
		display: flex;
		justify-content: center;
		align-items: center;
		margin-top: 10px;
		margin-bottom: 10px;

		&:hover {
			background-color: #fff1;
		}
	}
</style>
