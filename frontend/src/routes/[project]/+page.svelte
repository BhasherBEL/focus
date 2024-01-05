<script lang="ts">
	import { page } from '$app/stores';
	import { getProjectAPI } from '$lib/api/projects';
	import ProjectComponent from '$lib/components/project/project.svelte';
	import Sidebar from '$lib/components/sidebar.svelte';
	import type Project from '$lib/types/Project';
	import { SvelteToast } from '@zerodevx/svelte-toast';
	import { onMount } from 'svelte';

	let projectId: number = +$page.params.project;

	let project: Project;

	onMount(() => {
		getProjectAPI(projectId).then((p) => {
			project = p;
		});
	});
</script>

{#if project}
	<div>
		<Sidebar {project} />
		<ProjectComponent {project} />
	</div>
	<SvelteToast />
{/if}

<style>
	div {
		display: flex;
	}
</style>
