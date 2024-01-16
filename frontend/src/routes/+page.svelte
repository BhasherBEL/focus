<script lang="ts">
	import projectsApi from '$lib/api/projectsApi';
	import Project, { projects } from '$lib/types/Project';
	import { checkTauriUrl } from '$lib/utils/api';
	import WebSocketManager from '$lib/utils/webSocketManager';
	import { SvelteToast } from '@zerodevx/svelte-toast';
	import { onMount } from 'svelte';
	import SelectProject from '../lib/components/projects/SelectProject.svelte';

	const wsManager = new WebSocketManager();

	onMount(async () => {
		await checkTauriUrl(window);
		await projectsApi.getAll();
		wsManager.connect();
	});
</script>

<section>
	<h2>Projects</h2>
	<ul>
		{#if $projects}
			{#each $projects as project (project.id)}
				<SelectProject {project} />
			{/each}
		{/if}
	</ul>
	<div
		id="add"
		tabindex="0"
		role="button"
		on:click={Project.create}
		on:keydown={(e) => {
			if (e.key === 'Enter') {
				Project.create();
			}
		}}
	>
		<img src="/img/add-icon.svg" alt="Add" width="30" />
	</div>
</section>

<SvelteToast />

<style lang="less">
	section {
		margin: 40px;
	}
	h2 {
		text-align: center;
		margin-bottom: 40px;
		font-size: 3rem;
	}

	ul {
		list-style: none;
		padding: 0;
		margin: 0;
		font-size: 1.5rem;
	}
	#add {
		width: 100%;
		display: flex;
		justify-content: center;
		padding: 20px 0;
		cursor: pointer;
	}

	#add:hover {
		background-color: #303030;
	}
</style>
