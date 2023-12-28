<script lang="ts">
	import { projects } from "../stores/projects";

    let newProject = false;
    let editProject: number | undefined = undefined;

    function handleKeydown(event: KeyboardEvent, id: number | undefined = undefined) {
        if (event.key === 'Enter' && event.target) {
            if(id !== undefined){
                projects.edit({
                    id: id,
                    title: (event.target as HTMLInputElement).value
                })
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
            <img src="img/icon.svg" alt="">
            <span class="title">Focus</span>
            <span class="version">v0.0.1</span>
        </a>
    </div>
    <div class="boards">
        {#await projects.init()}
        <p>Loading ...</p>
        {:then} 
        <h2>Projects</h2>
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
                    <a href="/{project.id}">{project.title}</a><span class="edit-icon" on:click={() => editProject = project.id}>️
                        <svg xmlns="http://www.w3.org/2000/svg" height="16" width="18" viewBox="0 0 576 512"><path d="M402.6 83.2l90.2 90.2c3.8 3.8 3.8 10 0 13.8L274.4 405.6l-92.8 10.3c-12.4 1.4-22.9-9.1-21.5-21.5l10.3-92.8L388.8 83.2c3.8-3.8 10-3.8 13.8 0zm162-22.9l-48.8-48.8c-15.2-15.2-39.9-15.2-55.2 0l-35.4 35.4c-3.8 3.8-3.8 10 0 13.8l90.2 90.2c3.8 3.8 10 3.8 13.8 0l35.4-35.4c15.2-15.3 15.2-40 0-55.2zM384 346.2V448H64V128h229.8c3.2 0 6.2-1.3 8.5-3.5l40-40c7.6-7.6 2.2-20.5-8.5-20.5H48C21.5 64 0 85.5 0 112v352c0 26.5 21.5 48 48 48h352c26.5 0 48-21.5 48-48V306.2c0-10.7-12.9-16-20.5-8.5l-40 40c-2.2 2.3-3.5 5.3-3.5 8.5z" fill="#aaa"/></svg>
                    </span>
                {/if}
            </li>
        {/each}
            {#if newProject}
            <li>
                <input 
                type="text" 
                placeholder="Enter project title" 
                on:keydown={handleKeydown}
                style="padding: 8px; border: 1px solid #ccc; border-radius: 4px; width: 100%;" />
            </li>
            {/if}
        </ul>
        {:catch error}
        <p>Something went wrong: {error.message}</p>
        {/await}
    </div>
    <div class="bottom-links">
        <span on:click={() => newProject = true}>New project</span>
        <span>Settings</span>
    </div>
</div>