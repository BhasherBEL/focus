<script lang="ts">
	import Project from '../../components/project.svelte';
	import Sidebar from '../../components/sidebar.svelte';
	import { page } from '$app/stores';
	import { SvelteToast } from '@zerodevx/svelte-toast';
	import type { View } from '../../stores/interfaces';

	let projectID: number = +$page.params.project;

	let currentView: View;
</script>

<div id="projectPage">
	<Sidebar {projectID} />
	<Project projectId={projectID} />
</div>
<SvelteToast />

<style>
	#projectPage {
		display: flex;
	}
</style>
