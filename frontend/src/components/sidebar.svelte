<script lang="ts">
	import { onMount } from 'svelte';
	import api, { processError } from '../utils/api';
	import type { Project } from '../stores/interfaces';

	let newProject = false;
	let editProject: number | undefined = undefined;
	let projects: Project[];

	onMount(async () => {
		const response = await api.get(`/v1/projects`);

		if (response.status !== 200) {
			processError(response, 'Failed to fetch projects');
			return;
		}

		projects = response.data;
	});

	async function createProject(project: Project) {
		const response = await api.post(`/v1/projects`, project);

		if (response.status !== 201) {
			processError(response, 'Failed to create project');
			return;
		}

		project.id = response.data.id;

		projects = [...projects, project];
	}

	async function updateProject(project: Project) {
		const response = await api.put(`/v1/projects/${project.id}`, project);

		if (response.status !== 204) {
			processError(response, 'Failed to update project');
			return;
		}

		projects = projects.map((p) => {
			if (p.id === project.id) {
				return project;
			}
			return p;
		});
	}

	function handleKeydown(event: KeyboardEvent, id: number | undefined = undefined) {
		if (event.key === 'Enter' && event.target) {
			if (id !== undefined) {
				updateProject({
					id: id,
					title: (event.target as HTMLInputElement).value
				});
				editProject = undefined;
			} else {
				createNewProject((event.target as HTMLInputElement).value);
				newProject = false;
			}
		}
	}

	function createNewProject(value: string) {
		createProject({
			title: value,
			id: undefined
		});
	}
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
		{#if projects}
			<ul>
				{#each projects as project}
					<li>
						{#if editProject === project.id && editProject !== undefined}
							<input
								type="text"
								on:keydown={(e) => handleKeydown(e, project.id)}
								value={project.title}
								class="edit-input"
							/>
						{:else}
							<a href="/{project.id}">{project.title}</a>
							<img src="img/edit-icon.svg" alt="" on:click={() => (editProject = project.id)} />
						{/if}
					</li>
				{/each}
				{#if newProject}
					<li>
						<input
							type="text"
							placeholder="Enter project title"
							on:keydown={handleKeydown}
							style="padding: 8px; border: 1px solid #ccc; border-radius: 4px; width: 100%;"
						/>
					</li>
				{/if}
			</ul>
		{/if}
	</div>
	<div class="bottom-links">
		<span on:click={() => (newProject = true)}>New project</span>
		<span>Settings</span>
	</div>
</div>
