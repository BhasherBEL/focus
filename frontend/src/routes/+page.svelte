<script lang="ts">
	import { SvelteToast } from '@zerodevx/svelte-toast';
	import type { Project } from '../stores/interfaces';
	import { onMount } from 'svelte';
	import api, { processError } from '../utils/api';
	import SelectProject from '../components/projects/selectProject.svelte';
	import { toastAlert } from '../utils/toasts';

	let projects: Project[];

	onMount(async () => {
		try {
			const response = await api.get(`/v1/projects`);
			if (response.status !== 200) {
				processError(response, 'Failed to fetch projects');
				return;
			}

			projects = response.data || [];
		} catch (e: any) {
			toastAlert('Failed to fetch projects', e);
			setTimeout(() => {
				window.location.reload();
			}, 11000);
		}
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

<section>
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

<style lang="less">
	section {
		margin: 40px;
	}
	h2 {
		text-align: center;
		margin-bottom: 40px;
	}

	ul {
		list-style: none;
		padding: 0;
		margin: 0;
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
