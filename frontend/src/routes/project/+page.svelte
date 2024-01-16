<script lang="ts">
	import { page } from '$app/stores';
	import projectsApi from '$lib/api/projectsApi';
	import ProjectComponent from '$lib/components/project/Project.svelte';
	import Sidebar from '$lib/components/project/Sidebar.svelte';
	import currentView from '$lib/stores/currentView';
	import type Project from '$lib/types/Project';
	import { views } from '$lib/types/View';
	import { checkTauriUrl } from '$lib/utils/api';
	import { SvelteToast } from '@zerodevx/svelte-toast';
	import { onMount } from 'svelte';
	import { get } from 'svelte/store';

	let project: Project;

	onMount(async () => {
		await checkTauriUrl(window);
		const projectId = parseInt($page.url.searchParams.get('id') || '0');
		const res = await projectsApi.get(projectId);

		if (!res) return;

		project = res;

		await projectsApi.getTags(project);
		await projectsApi.getViews(project);
		await projectsApi.getCards(project);

		if (get(views).length > 0) {
			currentView.set(get(views)[0]);
		}
	});
</script>

{#if project}
	<div>
		<Sidebar {project} />
		<ProjectComponent {project} />
	</div>
	<SvelteToast />
{/if}

<style lang="less">
	div {
		display: flex;
	}
</style>
