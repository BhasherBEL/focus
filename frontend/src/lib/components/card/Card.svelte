<script lang="ts">
	import currentDraggedCard from '$lib/stores/currentDraggedCard';
	import type Card from '$lib/types/Card';
	import ModalCard from './ModalCard.svelte';

	export let card: Card;

	let showModal = false;
</script>

<div
	class="card"
	tabindex="0"
	draggable={true}
	on:dragstart={() => currentDraggedCard.set(card)}
	on:click={() => (showModal = true)}
	role="button"
	on:keydown={(e) => {
		if (e.key === 'Enter') showModal = true;
	}}
>
	<div class="title">{card.id} - {card.title}</div>
	<div class="tags">
		{#each card.cardTags as tag}
			{#if tag.option}
				<span class="tag" style="border: 1px solid #333">{tag.option.value}</span>
			{:else if tag.value}
				<span class="tag" style="border: 1px solid #333">{tag.value}</span>
			{/if}
		{/each}
	</div>
</div>

{#if showModal}
	<ModalCard {card} bind:showModal />
{/if}

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
