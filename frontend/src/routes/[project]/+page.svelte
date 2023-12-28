<script lang="ts">
	import Project from "../../components/project.svelte";
	import Sidebar from "../../components/sidebar.svelte";
	import { page } from '$app/stores';

	let projectId: number = +$page.params.project;
</script>

<Sidebar />
<Project {projectId} />