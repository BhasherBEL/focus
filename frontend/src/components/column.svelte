<script lang="ts">
	import type { Card, TagOption } from '../stores/interfaces';
	import CardC from './card.svelte';

	export let title: string;
	export let cards: Card[] = [];
</script>

<div class="column">
	<h3>{title}</h3>
	<ul>
		{#each cards as card}
			<CardC {card} />
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
