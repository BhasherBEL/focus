<script lang="ts">
	import projectsApi from '$lib/api/projectsApi';
	import Sidebar from '$lib/components/project/Sidebar.svelte';
	import type Project from '$lib/types/Project';
	import { SvelteToast } from '@zerodevx/svelte-toast';
	import { onMount } from 'svelte';
	import ProjectComponent from '$lib/components/project/Project.svelte';
	import { page } from '$app/stores';

	let project: Project;

	onMount(async () => {
		const projectId = parseInt($page.url.searchParams.get('id') || '0');
		const res = await projectsApi.get(projectId);

		if (!res) return;

		project = res;

		await projectsApi.getTags(project);
		await projectsApi.getViews(project);
		await projectsApi.getCards(project);
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
