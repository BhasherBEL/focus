<script lang="ts">
	import cards from '$lib/stores/cards';
	import currentDraggedCard from '$lib/stores/currentDraggedCard';
	import type Card from '$lib/types/Card';
	import type TagValue from '$lib/types/TagValue';
	import { get } from 'svelte/store';
	import projectTags from '../../stores/projectTags';
	import CardComponent from '../card/Card.svelte';
	import AddIcon from '../icons/AddIcon.svelte';
	import { import } from { createCardTagApi, deleteCardTagApi, updateCardTagApi };

	export let projectId: number;
	export let optionId: number | null = null;
	export let primary_tag_id: number | null = null;
	export let title: string;
	export let columnCards: Card[] = [];

	let lastTitle = title;

	async function onDrop(e: DragEvent) {
		e.preventDefault();
		if (!$currentDraggedCard || !$currentDraggedCard.tags) return;
		for (let tag of $currentDraggedCard.tags) {
			if (tag.tag_id !== primary_tag_id) continue;
			if (tag.option_id == optionId) return;

			try {
				if (tag.option_id && optionId) await deleteCardTagApi(tag.card_id, tag.tag_id);
				else if (tag.option_id && optionId)
					await createCardTagApi(tag.card_id, tag.tag_id, tag.option_id, tag.value);
				else await updateCardTagApi(tag.card_id, tag.tag_id, optionId, tag.value);

				tag.option_id = optionId;
				cards.reload();
			} catch (e) {}
			break;
		}
		currentDraggedCard.set(null);
	}

	async function addCard() {
		const tags: TagValue[] = [];
		for (let tag of Object.values(get(projectTags))) {
			if (tag.id === primary_tag_id) {
				tags.push({
					card_id: -1,
					tag_id: tag.id,
					option_id: optionId,
					value: null
				});
			}
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
			bind:value={title}
			type="text"
			on:blur={async () => {
				if (lastTitle === title) return;
				if (!optionId || !primary_tag_id) return;
				await updateTagOptionAPI({
					id: optionId,
					tag_id: primary_tag_id,
					value: title
				});
				lastTitle = title;
				cards.reload();
			}}
			disabled={optionId === null}
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
			<CardComponent {card} />
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
