<script lang="ts">
	import axios from 'axios';
	import { onMount } from 'svelte';

	const backend = 'http://127.0.0.1:3000';

	interface Card {
		id: number;
		title: string;
		content: string;
	}

	export let card: Card = {
		id: 0,
		title: 'No title',
		content: 'Nocontent'
	};

	interface FullTag {
		tag_id: number;
		tag_title: string;
		value: string;
	}

	let tags: FullTag[];

	onMount(async () => {
		const response = await axios.get(`${backend}/api/cardtags/${card.id}`);

		if (response.data.status === 'ok') {
			tags = response.data.data;
		} else {
			console.error(response.data);
		}
	});
</script>

<svelte:head>
	<link rel="stylesheet" type="text/css" href="/css/card.css" />
</svelte:head>

<div class="card" draggable={true}>
	<div class="title">{card.title}</div>
	{#if tags}
		<div class="tags">
			<!-- <span class="tag" style="background-color: #874d45;">HIGH</span>
			<span class="tag" style="background-color: #4a8645;">PERSONAL</span> -->
			{#each tags as tag}
				<span class="tag" style="border: 1px solid #333">{tag.value}</span>
			{/each}
		</div>
	{/if}
</div>
