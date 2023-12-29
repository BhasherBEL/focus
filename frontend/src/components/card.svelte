<script lang="ts">
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
</script>

<svelte:head>
	<link rel="stylesheet" type="text/css" href="/css/card.css" />
</svelte:head>

<div class="card" draggable={true}>
	<div class="title">{card.title}</div>
	<div class="tags">
		<span class="tag" style="background-color: #874d45;">HIGH</span>
		<span class="tag" style="background-color: #4a8645;">PERSONAL</span>
	</div>
</div>
