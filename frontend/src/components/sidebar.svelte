<script lang="ts">
	import { onMount } from 'svelte';
	import api, { processError } from '../utils/api';
	import type { Project, View } from '../stores/interfaces';
	import { currentView } from '../stores/smallStore';
	import ViewIcon from './icons/viewIcon.svelte';

	export let project: Project;
	let views: View[];

	onMount(async () => {
		const response = await api.get(`/v1/projects/${project.id}/views`);

		if (response.status !== 200) {
			processError(response, 'Failed to fetch views');
			return;
		}

		views = response.data;

		if (views.length > 0) currentView.set(views[0]);
	});
</script>

<nav>
	<div>
		<div id="branding">
			<span id="title">Focus.</span>
			<span id="version">v0.0.1</span>
		</div>
		<div id="views">
			{#if views}
				<h2>{project.title}</h2>
				<ul>
					{#each views as view}
						<li>
							<ViewIcon />
							<!-- on:click={() => {
									currentView.set(view);
								}} -->
							<span>
								{view.title}
							</span>
						</li>
					{/each}
				</ul>
			{/if}
		</div>
	</div>
	<div>
		<div class="separator"></div>
		<div id="newView" on:click={() => {}}>+ New view</div>
	</div>
</nav>

<style lang="less">
	nav {
		min-width: 300px;
		background-color: #273049;
		height: 100vh;
		display: flex;
		flex-direction: column;
		justify-content: space-between;
		font-size: 22px;
	}

	#branding {
		display: flex;
		flex-direction: row;
		align-items: center;
		justify-content: space-between;
		padding: 20px;

		#title {
			font-size: 40px;
		}

		#version {
			font-size: 30px;
			color: #aaa;
		}
	}

	#views {
		h2 {
			font-weight: normal;
			font-size: 25px;
			padding: 20px 10px;
		}

		ul {
			font-size: 22px;
			padding: 10px;
		}

		span {
			padding-left: 10px;
		}
	}

	.separator {
		height: 2px;
		widows: 100%;
		background-color: #444;
	}

	#newView {
		text-align: center;
		padding: 20px 0;
	}

	#newView:hover {
		cursor: pointer;
		background-color: #fff1;
	}
</style>
