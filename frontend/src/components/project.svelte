<script lang="ts">
	import { onMount } from 'svelte';
	import { type Project, type Card, type TagOption, type View } from '../stores/interfaces';
	import projectTags from '../stores/projectTags';
	import { getProjectAPI } from '../api/projects';
	import Column from './column.svelte';
	import { cards, currentView } from '../stores/smallStore';

	export let projectId: number;

	let project: Project;
	// let cards: Card[] = [];
	let view: View | null = null;
	let columns: { id: number; title: string; cards: Card[] }[] = [];

	onMount(async () => {
		getProjectAPI(projectId).then((p) => {
			project = p;
		});

		cards.init(projectId);

		cards.subscribe((c) => {
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

	function loadColumns() {
		if (!view) return;
		let primary_tag_id = view.primary_tag_id;
		columns =
			$projectTags[primary_tag_id]?.options.map((o) => {
				return {
					id: o.id,
					title: o.value,
					cards: $cards?.filter((c) => c.tags.map((t) => t.option_id).includes(o.id)) || []
				};
			}) || [];
		columns.push({
			id: -1,
			title: 'No tag',
			cards:
				$cards?.filter((c) => {
					const tag = c.tags.find((t) => t.tag_id === primary_tag_id);
					return tag?.option_id == -1;
				}) || []
		});
	}

	async function newCard() {
		await cards.add(projectId);
	}
</script>

<svelte:head>
	<link rel="stylesheet" type="text/css" href="/css/project.css" />
	<link rel="stylesheet" type="text/css" href="/css/card.css" />
	<link rel="stylesheet" type="text/css" href="/css/modalCard.css" />
</svelte:head>

{#if project}
	<section>
		<header>
			<h2>{project.title}</h2>
			<button on:click={newCard}>New card</button>
		</header>
		{#if view && $projectTags[view.primary_tag_id] && $cards}
			<div class="grid">
				{#each $projectTags[view.primary_tag_id].options as option}
					<Column
						{option}
						columnCards={$cards.filter((c) => c.tags.map((t) => t.option_id).includes(option.id))}
					/>
				{/each}
				<Column
					option={{
						id: -1,
						tag_id: view.primary_tag_id,
						value: `No ${$projectTags[view.primary_tag_id].title}`
					}}
					columnCards={$cards.filter((c) => c.tags.find((t) => t.tag_id)?.option_id == -1 || false)}
				/>
			</div>
		{/if}
	</section>
{/if}

<style>
	section {
		display: flex;
		flex-direction: column;
	}

	.grid {
		display: flex;
		flex-direction: row;
		flex: 1;
	}
</style>
