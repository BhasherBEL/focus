<script lang="ts">
	import { projects } from '../stores/projects';

	let newProject = false;
	let editProject: number | undefined = undefined;

	function handleKeydown(event: KeyboardEvent, id: number | undefined = undefined) {
		if (event.key === 'Enter' && event.target) {
			if (id !== undefined) {
				projects.edit({
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
		projects.add({
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
		{#await projects.init()}
			<p>Loading ...</p>
		{:then}
			<h2>Projects</h2>
			{#if $projects}
				<ul>
					{#each $projects as project}
						<li>
							{#if editProject === project.id}
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
		{:catch error}
			<p>Something went wrong: {error.message}</p>
		{/await}
	</div>
	<div class="bottom-links">
		<span on:click={() => (newProject = true)}>New project</span>
		<span>Settings</span>
	</div>
</div>
