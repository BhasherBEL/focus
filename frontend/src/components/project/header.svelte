<script lang="ts">
	import { get } from 'svelte/store';
	import type { Project, TagValue, View } from '../../stores/interfaces';
	import { cards, currentView } from '../../stores/smallStore';
	import projectTags from '../../stores/projectTags';
	import GroupMenu from './groupMenu.svelte';

	export let project: Project;
	export let currentTagId: number;
	let groupMenuOpen = false;

	function getEmptyTags(): TagValue[] {
		const tags: TagValue[] = [];
		for (let tag of Object.values(get(projectTags))) {
			tags.push({
				card_id: -1,
				tag_id: tag.id,
				option_id: -1,
				value: ''
			});
		}
		return tags;
	}

	async function setGroup(id: number): Promise<boolean> {
		if ($currentView == null) return false;

		return await currentView.update({
			...$currentView,
			primary_tag_id: id
		});
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
				currentChoice={currentTagId}
			/>
		</div>
		<div>
			<button class:disabled={true}>Sub-group</button>
		</div>
		<button class:disabled={true}>Filter</button>
		<button class:disabled={true}>Sort</button>
		<button id="newButton" on:click={async () => cards.add(project.id, getEmptyTags())}>New</button>
	</nav>
</header>

<style lang="less">
	header {
		padding: 20px 0;
		display: flex;
		align-items: center;
		justify-content: space-between;
		border-bottom: 2px solid #444;
	}

	h2 {
		font-size: 40px;
	}

	nav {
		display: flex;
		flex-direction: row;
		align-items: center;

		button {
			cursor: pointer;
			color: #aaa;
			padding: 5px 10px;
			margin-left: 10px;
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
