<script lang="ts">
	import axios from "axios";
	import { onMount } from "svelte";
	import Card from "./card.svelte";

    export let projectId: number = 0;

    const backend = 'http://127.0.0.1:3000';

    interface Project {
        id: number | undefined,
        title: string,
    }

    
    interface Card {
        id: number,
        title: string,
        content: string,
    }
    
    let project: Project;
    let cards: Card[];

    onMount(async () => {
        let response = await axios.get(`${backend}/api/project/${projectId}`);

        project = response.data; 

        response = await axios.get(`${backend}/api/cards/${projectId}`)

        if(response.data.status === "ok") {
            cards = response.data.data;
        } else {
            console.error(response.data)
        }
    })
</script>

<svelte:head>
    <link rel="stylesheet" type="text/css" href="/css/project.css" />
</svelte:head>

{#if project}
<div id="project">
    <h2>{project.title}</h2>

    <ul>
        {#if cards}
        {#each cards as card}
            <Card {card} />
        {/each}
        {/if}
    </ul>
</div>
{/if}