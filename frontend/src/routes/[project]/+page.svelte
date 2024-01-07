<script lang="ts">
	import { page } from '$app/stores';
	import projectsApi from '$lib/api/projectsApi';
	import Sidebar from '$lib/components/project/Sidebar.svelte';
	import type Project from '$lib/types/Project';
	import { SvelteToast } from '@zerodevx/svelte-toast';
	import { onMount } from 'svelte';

	let projectId: number = +$page.params.project;

	let project: Project;

	onMount(async () => {
		const res = await projectsApi.get(projectId);

		if (!res) return;

		project = res;

		await projectsApi.getTags(project);
		await projectsApi.getViews(project);
	});
</script>

{#if project}
	<div>
		<Sidebar {project} />
		<!-- <ProjectComponent {project} /> -->
	</div>
	<SvelteToast />
{/if}

<style>
	div {
		display: flex;
	}
</style>
