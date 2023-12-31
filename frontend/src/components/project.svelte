<script lang="ts">
	import { onMount } from 'svelte';
	import CardC from './card.svelte';
	import { type Project, type Card, parseCards, type View } from '../stores/interfaces';
	import status from '../utils/status';
	import api, { processError } from '../utils/api';
	import projectTags from '../stores/projectTags';
	import currentView from '../stores/currentView';
	import Card from './card.svelte';

	export let projectId: number;

	let project: Project;
	let cards: Card[];

	onMount(async () => {
		let response = await api.get(`/v1/projects/${projectId}`);

		if (response.status !== status.OK) {
			processError(response, 'Failed to fetch project');
			return;
		}

		project = response.data;

		response = await api.get(`/v1/projects/${projectId}/cards`);

		if (response.status === status.OK) {
			cards = parseCards(response.data);
		} else {
			cards = [];
			processError(response, 'Failed to fetch cards');
			return;
		}

		if (!(await projectTags.init(projectId))) {
			return;
		}
	});

	let modalID = -1;

	async function newCard() {
		const response = await api.post(`/v1/cards`, {
			project_id: projectId,
			title: 'Untitled',
			content: ''
		});

		if (response.status !== status.Created) {
			processError(response, 'Failed to create card');
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
		const response = await api.delete(`/v2/cards/${cardID}`);

		if (response.status !== status.NoContent) {
			processError(response, 'Failed to delete card');
			return;
		}

		cards = cards.filter((card) => card.id !== cardID);
	}

	let view: View | null = null;
	let columns: { id: number; title: string; cards: Card[] }[] = [];

	currentView.subscribe((v) => {
		view = v;
		if (!v) return;
		let primary_tag_id = v.primary_tag_id;
		columns = $projectTags[primary_tag_id].options.map((o) => {
			return {
				id: o.id,
				title: o.value,
				cards: cards.filter((c) => c.tags.map((t) => t.option_id).includes(o.id))
			};
		});
		columns.push({
			id: -1,
			title: 'No tag',
			cards: cards.filter((c) => {
				const tag = c.tags.find((t) => t.tag_id === primary_tag_id);

				return tag?.option_id == -1;
			})
		});
	});
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
		{#if view}
			<div class="grid">
				{#each columns as column}
					<div class="column">
						<h3>{column.title}</h3>
						<ul>
							{#each column.cards as card}
								<CardC
									{card}
									showModal={modalID === card.id}
									onDelete={async () => await deleteCard(card.id)}
								/>
							{/each}
						</ul>
					</div>
				{/each}
			</div>
		{:else}
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
		{/if}
	</div>
{/if}

<style>
	#project .grid {
		display: flex;
		flex-direction: row;
	}

	#project .column {
		width: 200px;
		margin: 0 10px;
	}

	#project .column h3 {
		text-align: center;
	}
</style>
