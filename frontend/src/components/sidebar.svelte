<script lang="ts">
	import { onMount } from 'svelte';
	import api, { processError } from '../utils/api';
	import type { View } from '../stores/interfaces';
	import currentView from '../stores/currentView';

	export let projectID: number;
	let views: View[];

	onMount(async () => {
		const response = await api.get(`/v1/projects/${projectID}/views`);

		if (response.status !== 200) {
			processError(response, 'Failed to fetch views');
			return;
		}

		views = response.data;
	});
</script>

<svelte:head>
	<link rel="stylesheet" type="text/css" href="/css/sidebar.css" />
</svelte:head>

<div id="sidebar" class="sidebar">
	<div class="logo">
		<a href="/">
			<img src="img/icon.svg" alt="" />
			<span class="title">Focus</span>
			<span class="version">v0.0.1</span>
		</a>
	</div>
	<div class="boards">
		<h2>Projects</h2>
		{#if views}
			<ul>
				{#each views as view}
					<li>
						<span
							on:click={() => {
								currentView.set(view);
							}}>{view.title}</span
						>
					</li>
				{/each}
			</ul>
		{/if}
	</div>
</div>
