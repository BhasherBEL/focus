<script lang="ts">
	import { onMount, tick } from 'svelte';
	import api, { processError } from '../utils/api';
	import type { Project, View } from '../stores/interfaces';
	import { currentView, views } from '../stores/smallStore';
	import ViewIcon from './icons/viewIcon.svelte';
	import projectTags from '../stores/projectTags';
	import EditIcon from './icons/editIcon.svelte';
	import { get } from 'svelte/store';

	export let project: Project;

	let viewEditId: number;
	let viewEditValue: string;

	onMount(async () => {
		await views.init(project.id);

		if ($views && $views.length > 0) currentView.set($views[0]);
	});

	async function createView() {
		if (!$views) return;

		const primaryTagId =
			$currentView?.primary_tag_id || Object.values($projectTags).find((t) => true)?.id || -1;

		const newView = await views.add(project.id, 'New view', primaryTagId);

		if (!newView) return;

		currentView.set(newView);
		viewEditId = newView.id;
		viewEditValue = newView.title;
		document.getElementById(`viewTitle-${newView.id}`)?.focus();
	}

	async function saveView(view: View) {
		if (!view || !$views.includes(view)) return;
		if (viewEditId === view.id && viewEditValue !== view.title) {
			if (!(await views.edit(view))) return;
		}

		viewEditId = -1;
		viewEditValue = '';
	}
</script>

<nav>
	<div>
		<div id="branding">
			<span id="title">Focus.</span>
			<span id="version">v0.0.1</span>
		</div>
		<div id="views">
			<h2>{project.title}</h2>
			{#if views}
				<ul>
					{#each get(views) as view}
						<!-- svelte-ignore a11y-no-noninteractive-element-to-interactive-role -->
						<li
							on:click={() => currentView.set(view)}
							role="button"
							tabindex="0"
							on:keydown={(e) => {
								if (e.key === 'Enter') {
									currentView.set(view);
								}
							}}
							class:active={$currentView === view}
						>
							<ViewIcon />
							<input
								type="text"
								readonly={viewEditId !== view.id}
								bind:value={view.title}
								class:inEdit={viewEditId === view.id}
								on:blur={() => saveView(view)}
								id="viewTitle-{view.id}"
								on:keydown={(e) => {
									if (e.key === 'Enter') {
										e.currentTarget.blur();
									}
								}}
							/>
							<button
								on:click={() => {
									if (viewEditId === view.id) {
										saveView(view);
									} else {
										viewEditId = view.id;
										viewEditValue = view.title;
										document.getElementById(`viewTitle-${view.id}`)?.focus();
									}
								}}
							>
								<EditIcon />
							</button>
						</li>
					{/each}
				</ul>
			{/if}
		</div>
	</div>
	<div>
		<div class="separator"></div>
		<div
			id="newView"
			on:click={createView}
			role="button"
			tabindex="0"
			on:keydown={(e) => {
				if (e.key === 'Enter') {
					createView();
				}
			}}
		>
			+ New view
		</div>
	</div>
</nav>

<style lang="less">
	nav {
		min-width: 300px;
		background-color: #273049;
		height: 100vh;
		display: flex;
		flex-direction: column;
		justify-content: space-between;
		font-size: 22px;
	}

	#branding {
		display: flex;
		flex-direction: row;
		align-items: center;
		justify-content: space-between;
		padding: 20px;

		#title {
			font-size: 40px;
		}

		#version {
			font-size: 30px;
			color: #aaa;
		}
	}

	#views {
		h2 {
			font-weight: normal;
			font-size: 25px;
			padding: 20px 10px;
		}

		ul {
			font-size: 22px;
			padding: 10px;
		}

		li {
			cursor: pointer;
			padding: 0 10px;
		}

		input {
			cursor: pointer;
			padding: 10px;
			border-radius: 5px;
			background-color: transparent;
			border: none;
			color: inherit;
			font-size: 17px;
			width: 60%;

			&:focus {
				outline: 0;
			}

			&.inEdit {
				background-color: #fff5;
			}
		}

		span {
			padding-left: 10px;
		}

		button {
			background-color: transparent;
			border: none;
			color: inherit;
			font-size: 17px;
			padding: 5px 0;
			float: right;
			cursor: pointer;
		}

		.active {
			background-color: #fff1;
		}
	}

	.separator {
		height: 2px;
		widows: 100%;
		background-color: #444;
	}

	#newView {
		text-align: center;
		padding: 20px 0;
	}

	#newView:hover {
		cursor: pointer;
		background-color: #fff1;
	}
</style>
