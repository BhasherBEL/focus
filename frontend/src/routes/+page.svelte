<script lang="ts">
	import { SvelteToast } from '@zerodevx/svelte-toast';
	import type { Project } from '../stores/interfaces';
	import { onMount } from 'svelte';
	import api, { processError } from '../utils/api';
	import SelectProject from '../components/selectProject.svelte';

	let projects: Project[];

	onMount(async () => {
		const response = await api.get(`/v1/projects`);

		if (response.status !== 200) {
			processError(response, 'Failed to fetch projects');
			return;
		}

		projects = response.data || [];
	});

	async function deleteProject(project: Project) {
		if (!confirm(`Are you sure you want to delete ${project.title}?`)) {
			return;
		}
		const response = await api.delete(`/v1/projects/${project.id}`);

		if (response.status !== 204) {
			processError(response, 'Failed to delete project');
			return;
		}

		projects = projects.filter((p) => p.id !== project.id);
	}

	async function createProject() {
		const response = await api.post(`/v1/projects`, {
			title: 'New Project'
		});

		if (response.status !== 201) {
			processError(response, 'Failed to create project');
			return;
		}

		projects = [...projects, response.data];
	}
</script>

<svelte:head>
	<link rel="stylesheet" type="text/css" href="/css/projects.css" />
</svelte:head>

<section id="projects">
	<h2>Projects</h2>
	<ul>
		{#if projects}
			{#each projects as project}
				<SelectProject {project} {deleteProject} />
			{/each}
		{/if}
	</ul>
	<div
		id="add"
		tabindex="0"
		role="button"
		on:click={createProject}
		on:keydown={(e) => {
			if (e.key === 'Enter') {
				createProject();
			}
		}}
	>
		<img src="/img/add-icon.svg" alt="Add" width="30" />
	</div>
</section>

<SvelteToast />
