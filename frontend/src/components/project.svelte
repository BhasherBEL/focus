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

	let modalID = -1;

	async function newCard() {
		const response = await axios.post(`${backend}/api/card`, {
			project_id: projectId,
			title: 'Untitled',
			content: ''
		});

		if (response.data.status !== 'ok') {
			console.error(response.data);
			return;
		}

		const id: number = response.data.id;

		let card: Card = {
			id: id,
			project_id: projectId,
			title: 'Untitled',
			content: '',
			tags: []
		};

		cards = [...cards, card];
		modalID = id;
	}

	async function deleteCard(cardID: number) {
		const response = await axios.delete(`${backend}/api/card/${cardID}`);

		if (response.status !== 204) {
			console.error(response.data);
			return;
		}

		cards = cards.filter((card) => card.id !== cardID);
	}
</script>

<svelte:head>
	<link rel="stylesheet" type="text/css" href="/css/project.css" />
	<link rel="stylesheet" type="text/css" href="/css/card.css" />
	<link rel="stylesheet" type="text/css" href="/css/modalCard.css" />
</svelte:head>

{#if project}
	<div id="project">
		<header>
			<h2>{project.title}</h2>
			<button on:click={newCard}>New card</button>
		</header>
		<ul>
			{#if cards}
				{#each cards as card}
					<CardC
						{card}
						showModal={modalID === card.id}
						onDelete={async () => await deleteCard(card.id)}
					/>
				{/each}
			{/if}
		</ul>
	</div>
{/if}
