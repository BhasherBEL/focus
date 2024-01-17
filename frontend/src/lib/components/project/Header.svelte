<script lang="ts">
	import GroupMenu from '$lib/components/menu/GroupMenu.svelte';
	import SortMenu from '$lib/components/menu/SortMenu.svelte';
	import currentView from '$lib/stores/currentView';
	import Card, { cards } from '$lib/types/Card';
	import type Project from '$lib/types/Project';
	import type ProjectTag from '$lib/types/ProjectTag';
	import { projectTags } from '$lib/types/ProjectTag';
	import { get } from 'svelte/store';
	import FilterMenu from '../menu/FilterMenu.svelte';

	export let project: Project;
	let groupMenuOpen = false;
	let sortMenuOpen = false;
	let filterMenuOpen = false;

	async function setGroup(projectTag: ProjectTag): Promise<boolean> {
		const view = get(currentView);
		if (!view) return false;

		const res = await view.setPrimaryTag(projectTag);

		if (res) currentView.reload();

		return res;
	}

	async function setSort(projectTag: ProjectTag): Promise<boolean> {
		const view = get(currentView);
		if (!view) return false;

		const res = await view.setSortTag(projectTag, view.sortDirection ? -view.sortDirection : 1);

		if (res) currentView.reload();

		return res;
	}

	async function addCard() {
		const card = await Card.create(project);

		if (!card) return;

		if ($currentView?.filters && $currentView.filters.length > 0) {
			for (const projectTag of $projectTags) {
				for (const filter of $currentView.filters) {
					if (projectTag !== filter.projectTag) continue;
					if (!filter.tagOption) continue;
					if (filter.filterType !== 0) continue;

					if (await card.addTag(projectTag, filter.tagOption, null)) break;
				}
			}
		}

		cards.reload();
	}
</script>

<header>
	<h2>{project.title}</h2>
	<nav>
		<div>
			<button
				on:click={() => (groupMenuOpen = !groupMenuOpen)}
				class:defined={$currentView?.primaryTag}>Group</button
			>
			<GroupMenu
				bind:isOpen={groupMenuOpen}
				choices={$projectTags}
				onChoice={async (projectTag) => {
					if (!(await setGroup(projectTag))) return;
					groupMenuOpen = false;
				}}
				currentChoice={$currentView?.primaryTag || null}
			/>
		</div>
		<button class:disabled={true}>Sub-group</button>
		<div>
			<button
				on:click={() => (filterMenuOpen = !filterMenuOpen)}
				class:defined={$currentView?.filters && $currentView?.filters.length > 0}
			>
				Filter
			</button>
			<FilterMenu bind:isOpen={filterMenuOpen} />
		</div>
		<div>
			<button on:click={() => (sortMenuOpen = !sortMenuOpen)} class:defined={$currentView?.sortTag}>
				Sort
			</button>
			<SortMenu
				bind:isOpen={sortMenuOpen}
				choices={$projectTags}
				onChoice={async (projectTag) => {
					if (!(await setSort(projectTag))) return;
					sortMenuOpen = false;
				}}
				currentChoice={$currentView?.sortTag || null}
				currentDirection={$currentView?.sortDirection || null}
			/>
		</div>
		<button id="newButton" on:click={addCard}>New</button>
	</nav>
</header>

<style lang="less">
	@import '../../styles/breakpoints.less';

	header {
		margin: 0 2%;
		padding: 20px 0;
		display: flex;
		align-items: center;
		justify-content: space-between;
		flex-wrap: wrap;
		border-bottom: 2px solid #444;
	}

	h2 {
		flex-grow: 100000;
		font-size: 3rem;
		margin: 0;

		.nosidebar({
			padding-left: 50px;
		});
	}

	nav {
		display: flex;
		flex-grow: 1;
		flex-shrink: 0;
		flex-direction: row;
		align-items: center;
		justify-content: center;
		gap: 7px;

		button {
			cursor: pointer;
			color: #aaa;
			padding: 5px 10px;
			border-radius: 7px;
			border: none;
			background-color: transparent;
			font-size: 1.5rem;
			flex-shrink: 0;

			&.defined {
				color: #6481cc;
			}

			&.disabled {
				color: #555;
				cursor: not-allowed;
			}
		}

		#newButton {
			background: #324067;
			border-radius: 10px;
			padding: 10px 20px;
		}
	}
</style>
