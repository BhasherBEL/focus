<script>
	import { projects } from "../stores/projects";
</script>

<svelte:head>
    <link rel="stylesheet" type="text/css" href="/css/index.css" />
</svelte:head>

{#await projects.init()}
<p>Loading ...</p>
{:then} 
<div id="sidebar" class="sidebar">
    <div class="logo">
        <img src="img/icon.svg" alt="">
        <span class="title">Focus</span>
        <span class="version">v0.0.1</span>
    </div>
    <div class="boards">
        <h2>Projects</h2>
        <ul>
                {#each $projects as project}
                    <!-- <li><a href="/{{project.id}}">{{project.title}}</a></li> -->
                {/each}
            </ul>
        </div>
        <div class="bottom-links">
            <a href="#">Add board</a>
            <a href="#">Settings</a>
        </div>
    </div>
{:catch error}
    <p>Something went wrong: {error.message}</p>
{/await}