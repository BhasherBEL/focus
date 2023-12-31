<script lang="ts">
	import { onMount } from 'svelte';
	import CardC from './card.svelte';
	import { type Project, type Card, parseCards, type View } from '../stores/interfaces';
	import projectTags from '../stores/projectTags';
	import currentView from '../stores/currentView';
	import { deleteCardApi, newCardApi } from '../api/cards';
	import { getProjectAPI, getProjectCardsAPI } from '../api/projects';
	import Column from './column.svelte';
	import currentModalCard from '../stores/currentModalCard';

	export let projectId: number;

	let project: Project;
	let cards: Card[] = [];
	let view: View | null = null;
	let columns: { id: number; title: string; cards: Card[] }[] = [];

	onMount(async () => {
		getProjectAPI(projectId).then((p) => {
			project = p;
		});

		getProjectCardsAPI(projectId).then((c) => {
			cards = parseCards(c);
			loadColumns();
		});

		if (!(await projectTags.init(projectId))) {
			return;
		}

		currentView.subscribe((v) => {
			view = v;
			loadColumns();
		});
	});

	function newCard() {
		newCardApi(projectId).then((card) => {
			cards = [...cards, card];
			currentModalCard.set(card.id);
			loadColumns();
		});
	}

	function deleteCard(id: number) {
		deleteCardApi(id).then(() => {
			cards = cards.filter((card) => card.id !== id);
			loadColumns();
		});
	}

	function loadColumns() {
		if (!view) return;
		let primary_tag_id = view.primary_tag_id;
		columns =
			$projectTags[primary_tag_id]?.options.map((o) => {
				return {
					id: o.id,
					title: o.value,
					cards: cards.filter((c) => c.tags.map((t) => t.option_id).includes(o.id))
				};
			}) || [];
		columns.push({
			id: -1,
			title: 'No tag',
			cards: cards.filter((c) => {
				const tag = c.tags.find((t) => t.tag_id === primary_tag_id);
				return tag?.option_id == -1;
			})
		});
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
		{#if view && $projectTags[view.primary_tag_id]}
			<div class="grid">
				{#each $projectTags[view.primary_tag_id].options as option}
					<Column tag_id={view.primary_tag_id} {option} bind:cards deleteCard />
				{/each}
			</div>
		{:else}
			<!-- <ul>
				{#if cards}
					{#each cards as card}
						<CardC
							{card}
							showModal={modalID === card.id}
							onDelete={async () => await deleteCard(card.id)}
						/>
					{/each}
				{/if}
			</ul> -->
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
