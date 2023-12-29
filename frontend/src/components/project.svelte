<script lang="ts">
	import axios from 'axios';
	import { onMount } from 'svelte';
	import CardC from './card.svelte';
	import { backend } from '../stores/config';
	import type { Card, Project } from '../stores/interfaces';

	export let projectId: number = 0;

	let project: Project;
	let cards: Card[];

	onMount(async () => {
		let response = await axios.get(`${backend}/api/project/${projectId}`);

		project = response.data;

		response = await axios.get(`${backend}/api/cards/${projectId}`);

		if (response.data.status === 'ok') {
			cards = response.data.data;
		} else {
			console.error(response.data);
		}
	});
</script>

<svelte:head>
	<link rel="stylesheet" type="text/css" href="/css/project.css" />
	<link rel="stylesheet" type="text/css" href="/css/card.css" />
	<link rel="stylesheet" type="text/css" href="/css/modalCard.css" />
</svelte:head>

{#if project}
	<div id="project">
		<h2>{project.title}</h2>

		<ul>
			{#if cards}
				{#each cards as card}
					<CardC {card} />
				{/each}
			{/if}
		</ul>
	</div>
{/if}
