<script lang="ts">
	import type Project from '$lib/types/Project';

	export let project: Project;

	let edit = false;
	let newTitle = project.title;

	function focus(el: HTMLElement) {
		el.focus();
	}

	async function updateProject() {
		if (newTitle === project.title) {
			edit = false;
			return;
		}

		if (!(await project.setTitle(newTitle))) return;

		edit = false;
	}
</script>

<!-- svelte-ignore a11y-no-noninteractive-element-interactions -->
<!-- svelte-ignore a11y-no-noninteractive-element-to-interactive-role -->
<li>
	{#if edit}
		<input
			type="text"
			bind:value={newTitle}
			on:blur={updateProject}
			on:keydown={(e) => {
				if (e.key === 'Enter') {
					updateProject();
				}
			}}
			use:focus
		/>
	{:else}
		<div
			class="title"
			on:click={() => (location.href = location.origin + `/project/?id=${project.id}`)}
			on:keydown={(e) => {
				if (e.key === 'Enter') {
					location.href = location.origin + `/project/?id=${project.id}`;
				}
			}}
			tabindex="0"
			role="button"
		>
			{project.title}
		</div>
	{/if}
	<!-- svelte-ignore a11y-no-static-element-interactions -->
	<div class="buttons" on:keydown|stopPropagation>
		{#if !edit}
			<img
				src="/img/edit-icon.svg"
				alt="Edit"
				class="button"
				on:click={() => (edit = !edit)}
				role="button"
				tabindex="0"
				on:keydown|stopPropagation={(e) => {
					if (e.key === 'Enter') {
						edit = !edit;
					}
				}}
			/>
		{/if}
		<img
			src="/img/delete-icon.svg"
			alt="Delete"
			class="button"
			on:click={() => project.delete()}
			role="button"
			tabindex="0"
			on:keydown={(e) => {
				if (e.key === 'Enter') {
					project.delete();
				}
			}}
		/>
	</div>
</li>

<style lang="less">
	li {
		cursor: pointer;
		margin: 10px 0;
		border: 1px solid #555;
		border-radius: 4px;
		display: flex;
		flex-direction: row;
		align-items: center;
		justify-content: space-between;
	}

	.title {
		font-weight: bold;
		padding: 20px;
		width: 100%;
	}

	.title:hover {
		background-color: #303030;
	}

	.buttons {
		display: flex;
		flex-direction: row;
		align-items: center;
		justify-content: space-between;
	}

	li img {
		padding: 20px;
	}

	li img:hover {
		background-color: #333;
	}

	input {
		padding: 20px;
		width: 100%;
		background-color: #333;
		color: inherit;
		font-weight: bold;
		font-size: inherit;
	}
</style>
