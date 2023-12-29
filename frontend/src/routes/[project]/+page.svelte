<script lang="ts">
	import Project from '../../components/project.svelte';
	import Sidebar from '../../components/sidebar.svelte';
	import { page } from '$app/stores';

	let projectId: number = +$page.params.project;
</script>

<div id="projectPage">
	<Sidebar />
	<Project {projectId} />
</div>

<style>
	#projectPage {
		display: flex;
	}
</style>
