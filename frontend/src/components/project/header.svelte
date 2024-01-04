<script lang="ts">
	import type { Project, TagValue, View } from '../../stores/interfaces';
	import { cards, currentView, views } from '../../stores/smallStore';
	import projectTags from '../../stores/projectTags';
	import GroupMenu from './groupMenu.svelte';

	export let project: Project;
	export let view: View;
	let groupMenuOpen = false;

	async function setGroup(id: number): Promise<boolean> {
		if ($currentView == null) return false;

		const view = {
			...$currentView,
			primary_tag_id: id
		};

		const res = await views.edit(view);

		if (res) currentView.set(view);

		return res;
	}
</script>

<header>
	<h2>{project.title}</h2>
	<nav>
		<div>
			<button
				on:click={() => (groupMenuOpen = !groupMenuOpen)}
				class:defined={$currentView?.primary_tag_id}>Group</button
			>
			<GroupMenu
				isOpen={groupMenuOpen}
				choices={Object.values($projectTags).map((tag) => ({ id: tag.id, value: tag.title }))}
				onChoice={async (id) => {
					if (!(await setGroup(id))) return;
					groupMenuOpen = false;
				}}
				currentChoice={view?.primary_tag_id}
			/>
		</div>
		<button class:disabled={true}>Sub-group</button>
		<button class:disabled={true}>Filter</button>
		<button class:disabled={true}>Sort</button>
		<button id="newButton" on:click={async () => cards.add(project.id, [])}>New</button>
	</nav>
</header>

<style lang="less">
	header {
		margin: 0 50px;
		padding: 20px 0;
		display: flex;
		align-items: center;
		justify-content: space-between;
		flex-wrap: wrap;
		border-bottom: 2px solid #444;
	}

	h2 {
		font-size: 30px;
	}

	nav {
		display: flex;
		flex-direction: row;
		align-items: center;
		gap: 7px;

		button {
			cursor: pointer;
			color: #aaa;
			padding: 5px 10px;
			border-radius: 7px;
			border: none;
			background-color: transparent;
			font-size: inherit;

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
