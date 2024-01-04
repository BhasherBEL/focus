<script lang="ts">
	import type { Card } from '../../../stores/interfaces';
	import projectTags from '../../../stores/projectTags';
	import { currentModalCard, currentDraggedCard } from '../../../stores/smallStore';
	import ModalCard from './modal_card.svelte';

	export let card: Card;
</script>

<div
	class="card"
	tabindex="0"
	draggable={true}
	on:dragstart={() => currentDraggedCard.set(card)}
	on:click={() => ($currentModalCard = card.id)}
	role="button"
	on:keydown={(e) => {
		if (e.key === 'Enter') {
			$currentModalCard = card.id;
		}
	}}
>
	<div class="title">{card.title}</div>
	{#if card.tags}
		<div class="tags">
			{#each card.tags as tag}
				{#if tag.option_id}
					{#if $projectTags[tag.tag_id]}
						<span class="tag" style="border: 1px solid #333"
							>{$projectTags[tag.tag_id]?.options.find((o) => o.id == tag.option_id)?.value}</span
						>
					{/if}
				{:else if tag.value}
					<span class="tag" style="border: 1px solid #333">{tag.value}</span>
				{/if}
			{/each}
		</div>
	{/if}
</div>

<ModalCard bind:card />

<style lang="less">
	.card {
		padding: 10px;
		border-radius: 6px;
		border: 1px solid #555;
		font-family: 'Open Sans', sans-serif;
		font-size: 14px;
		margin-bottom: 10px;
	}

	.card:hover {
		background-color: #303030;
		cursor: pointer;
	}

	.card .title {
		font-weight: normal;
	}

	.card .tags {
		padding-top: 10px;
		font-weight: lighter;
	}

	.card .tag {
		padding: 2px 8px;
		margin: 4px 4px 0 0;
		text-transform: uppercase;
		border-radius: 3px;
		font-size: 90%;
	}
</style>
