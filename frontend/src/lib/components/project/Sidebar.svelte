<script lang="ts">
	import currentView from '$lib/stores/currentView';
	import Project from '$lib/types/Project';
	import View, { views } from '$lib/types/View';
	import { onMount } from 'svelte';
	import EditIcon from '../icons/EditIcon.svelte';
	import MenuOpener from '../icons/MenuOpener.svelte';
	import ViewIcon from '../icons/ViewIcon.svelte';

	export let project: Project;

	let viewEdit: View | null;
	let newTitle: string;

	let isVisible = false;

	onMount(() => {
		if (views && $views.length > 0) currentView.set($views[0]);
	});

	async function createView() {
		if (!$views) return;

		const newView = await View.create(project);

		if (!newView) return;

		currentView.set(newView);
		viewEdit = newView;
		document.getElementById(`viewTitle-${newView.id}`)?.focus();
	}

	async function saveView() {
		if (!viewEdit) return;
		if (!newTitle) return;
		if (newTitle != viewEdit.title) {
			await viewEdit.setTitle(newTitle);
		}

		viewEdit = null;
	}
</script>

<nav class:hidden={!isVisible}>
	<div>
		<div id="branding">
			<a href="/">
				<span id="title">Focus.</span>
				<span id="version">v0.3.2</span>
			</a>
		</div>
		<div id="views">
			<h2>{project.title}</h2>
			{#if views}
				<ul>
					{#each $views as view}
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
							{#if viewEdit && viewEdit === view}
								<input
									type="text"
									bind:value={newTitle}
									on:blur={saveView}
									id="viewTitle-{view.id}"
									class="inEdit"
									on:keydown={(e) => {
										if (e.key === 'Enter') {
											e.currentTarget.blur();
										}
									}}
								/>
							{:else}
								<span class="title">{view.title}</span>
							{/if}
							<button
								on:click={() => {
									if (viewEdit && viewEdit.id === view.id) {
										saveView();
									} else {
										viewEdit = view;
										newTitle = view.title;
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
<button class="toggle" class:open={isVisible} on:click={() => (isVisible = !isVisible)}>
	<MenuOpener />
</button>

<style lang="less">
	nav {
		position: fixed;
		top: 0;
		left: 0;
		bottom: 0;
		width: 250px;
		background-color: #273049;
		display: flex;
		flex-direction: column;
		justify-content: space-between;
		font-size: 22px;
		transition: transform 0.3s ease-in-out;
	}

	.toggle {
		visibility: hidden;
		position: fixed;
		top: 5px;
		left: 5px;
		border-radius: 50%;
		width: 50px;
		height: 50px;
		transform: scale(1);
		transition-property: left, transform;
		transition-duration: 0.3s;
		transition-timing-function: ease-in-out;

		&:hover {
			cursor: pointer;
			background-color: #fff2;
		}

		&:active {
			background-color: #fff2;
		}

		&.open {
			left: 195px;
			transform: scaleX(-1);
		}
	}

	@media (max-width: 800px) {
		nav.hidden {
			transform: translateX(-100%);
		}

		.toggle {
			visibility: visible;
		}
	}

	#branding {
		display: flex;
		flex-direction: row;
		align-items: end;
		gap: 10px;
		padding: 20px;

		#title {
			font-size: 40px;
		}

		#version {
			font-size: 25px;
			color: #aaa;
			margin-bottom: 3px;
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
			display: flex;
			flex-direction: row;
			align-items: center;
			justify-content: space-between;
			gap: 5px;
		}

		span.title,
		input {
			display: inline-block;
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
		}

		input.inEdit {
			background-color: #fff5;
		}

		button {
			background-color: transparent;
			border: none;
			color: inherit;
			font-size: 17px;
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
