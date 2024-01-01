<script lang="ts">
	import { get } from 'svelte/store';
	import type { Project, TagValue } from '../../stores/interfaces';
	import { cards } from '../../stores/smallStore';
	import projectTags from '../../stores/projectTags';

	export let project: Project;

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
</script>

<header>
	<h2>{project.title}</h2>
	<nav>
		<span>Group</span>
		<span>Sub-group</span>
		<span>Filter</span>
		<span>Sort</span>
		<button on:click={async () => cards.add(project.id, getEmptyTags())}>New</button>
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
		* {
			cursor: pointer;
		}

		span {
			margin-right: 10px;
			color: #aaa;
			padding: 5px 10px;
			border-radius: 7px;

			&:hover {
				// background-color: #fff2;
			}
		}

		button {
			background: #324067;
			color: inherit;
			border: none;
			border-radius: 10px;
			padding: 10px 20px;
			font-size: inherit;

			&:hover {
				// background-color: #3a4a77;
			}
		}
	}
</style>
