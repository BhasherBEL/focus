<script lang="ts">
	import axios from 'axios';
	import { onMount } from 'svelte';

	const backend = 'http://127.0.0.1:3000';

	interface Card {
		id: number;
		title: string;
		content: string;
		tags: Tag[];
	}

	interface Tag {
		tag_id: number;
		tag_title: string;
		value: string;
	}

	export let card: Card = {
		id: 0,
		title: 'No title',
		content: 'Nocontent',
		tags: []
	};
</script>

<svelte:head>
	<link rel="stylesheet" type="text/css" href="/css/card.css" />
</svelte:head>

<div class="card" draggable={true}>
	<div class="title">{card.title}</div>
	{#if card.tags}
		<div class="tags">
			{#each card.tags as tag}
				<span class="tag" style="border: 1px solid #333">{tag.value}</span>
			{/each}
		</div>
	{/if}
</div>
