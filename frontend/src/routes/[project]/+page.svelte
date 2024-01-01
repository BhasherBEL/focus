<script lang="ts">
	import { page } from '$app/stores';
	import { SvelteToast } from '@zerodevx/svelte-toast';
	import { onMount } from 'svelte';
	import { getProjectAPI } from '../../api/projects';
	import { type Project as P } from '../../stores/interfaces';
	import Sidebar from '../../components/sidebar.svelte';
	import Project from '../../components/project/project.svelte';

	let projectId: number = +$page.params.project;

	let project: P;

	onMount(() => {
		getProjectAPI(projectId).then((p) => {
			project = p;
		});
	});
</script>

{#if project}
	<div>
		<Sidebar {project} />
		<Project {project} />
	</div>
	<SvelteToast />
{/if}

<style>
	div {
		display: flex;
	}
</style>
