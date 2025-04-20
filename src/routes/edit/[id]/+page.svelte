<script lang="ts">
	import { enhance } from '$app/forms';
	import AlertIcon from '$lib/components/utils/alertIcon.svelte';

	const { form, data } = $props();
	const project = data.project;
</script>

<section class="mx-auto my-10 w-full max-w-3xl px-4 sm:px-8">
	<h2 class="mb-8 text-center text-3xl font-bold">Edit Project</h2>

	{#if form?.error}
		<div role="alert" class="alert alert-error mb-2">
			<AlertIcon />
			<span>{form.error}</span>
		</div>
	{/if}

	<form method="POST" use:enhance class="mx-auto flex w-full flex-col gap-4">
		<div class="form-control">
			<label for="title" class="label mb-1 font-semibold">Project Title</label>
			<input
				id="title"
				name="title"
				type="text"
				class="input input-bordered w-full p-2"
				required
				autocomplete="off"
				value={project.title}
			/>
		</div>
		<div class="mt-2 flex gap-4">
			<a
				href="/"
				class="btn btn-neutral flex flex-1 items-center justify-center"
				aria-label="Back to projects"
			>
				<span class="mr-2">&larr;</span> Back
			</a>
			<button type="submit" class="btn btn-primary flex-1 p-2">Save Changes</button>
		</div>
	</form>
</section>
