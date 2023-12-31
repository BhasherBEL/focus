<script lang="ts">
	import type { Card, TagOption } from '../stores/interfaces';
	import CardC from './card.svelte';

	export let tag_id: number;
	export let option: TagOption;
	export let cards: Card[] = [];
	export let deleteCard: (id: number) => void;

	let columnCards = cards.filter(
		(card) => card.tags.find((t) => t.tag_id === tag_id)?.option_id === option.id
	);
</script>

<div class="column">
	<h3>{option.value}</h3>
	<ul>
		{#each columnCards as card}
			<CardC {card} onDelete={async () => await deleteCard(card.id)} />
		{/each}
	</ul>
</div>

<style>
	h3 {
		text-align: center;
	}

	.column {
		width: 200px;
		margin: 0 10px;
	}
</style>
