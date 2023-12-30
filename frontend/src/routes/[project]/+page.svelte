<script lang="ts">
	import Project from '../../components/project.svelte';
	import Sidebar from '../../components/sidebar.svelte';
	import { page } from '$app/stores';
	import { SvelteToast } from '@zerodevx/svelte-toast';

	let projectId: number = +$page.params.project;
</script>

<div id="projectPage">
	<Sidebar />
	<Project {projectId} />
</div>
<SvelteToast />

<style>
	#projectPage {
		display: flex;
	}
</style>
