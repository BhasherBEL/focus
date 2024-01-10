<script lang="ts">
	import currentView from '$lib/stores/currentView';
	import Filter from '$lib/types/Filter';
	import ProjectTag, { projectTags } from '$lib/types/ProjectTag';
	import type TagOption from '$lib/types/TagOption';
	import Menu from './Menu.svelte';

	export let filter: Filter | null = null;

	let isProjectTagOpen = false;
	let isFilterTypeOpen = false;
	let isOptionOpen = false;

	async function selectProjectTag(projectTag: ProjectTag) {
		if (!$currentView) return;
		if (!filter) {
			filter = await $currentView?.addFilter(projectTag, 0, null);
			return;
		}

		if (filter.projectTag.id !== projectTag.id) {
			const res = await filter.setProjectTag(projectTag);
			if (!res) return;
			currentView.reload();
		}
		isProjectTagOpen = false;
	}

	async function selectFilterType(filterType: number) {
		if (!filter) return;

		if (filter.filterType !== filterType) {
			const res = await filter.setFilterType(filterType);
			if (!res) return;
			currentView.reload();
		}
		isFilterTypeOpen = false;
	}

	async function selectOption(option: TagOption) {
		if (!filter) return;

		if (filter.tagOption !== option) {
			const res = await filter.setTagOption(option);
			if (!res) return;
			currentView.reload();
		}
		isOptionOpen = false;
	}
</script>

<div class="item">
	<div>
		<div
			class="part"
			on:click={() => (isProjectTagOpen = !isProjectTagOpen)}
			on:keydown={(e) => {
				if (e.key === 'Enter') isProjectTagOpen = !isProjectTagOpen;
			}}
			tabindex="0"
			role="button"
		>
			{#if filter}
				{filter.projectTag.title}
			{/if}
		</div>
		<Menu bind:isOpen={isProjectTagOpen}>
			{#each $projectTags as projectTag}
				<button on:click={() => selectProjectTag(projectTag)}>
					{projectTag.title}
				</button>
			{/each}
		</Menu>
	</div>
	<div>
		<div
			class="part"
			on:click={() => (isFilterTypeOpen = !isFilterTypeOpen)}
			on:keydown={(e) => {
				if (e.key === 'Enter') isFilterTypeOpen = !isFilterTypeOpen;
			}}
			tabindex="0"
			role="button"
		>
			{#if filter}
				{filter.filterType}
			{/if}
		</div>
		{#if filter}
			<Menu bind:isOpen={isFilterTypeOpen}>
				<button on:click={() => selectFilterType(0)}> is </button>
				<button
					on:click={() => {
						selectFilterType(1);
						isFilterTypeOpen = false;
					}}
				>
					is not
				</button>
			</Menu>
		{/if}
	</div>
	<div>
		<div
			class="part"
			on:click={() => (isOptionOpen = !isOptionOpen)}
			on:keydown={(e) => {
				if (e.key === 'Enter') isOptionOpen = !isOptionOpen;
			}}
			tabindex="0"
			role="button"
		>
			{#if filter && filter.tagOption}
				{filter.tagOption.value}
			{/if}
		</div>
		{#if filter && filter.projectTag}
			<Menu bind:isOpen={isOptionOpen}>
				{#each filter.projectTag.options as option}
					<button on:click={() => selectOption(option)}>
						{option.value}
					</button>
				{/each}
			</Menu>
		{/if}
	</div>
</div>

<style lang="less">
	.item {
		display: flex;
		flex-direction: row;
	}

	.part {
		min-width: 50px;
		height: 30px;
		margin: 5px;
		text-align: center;
		line-height: 30px;
		padding: 0 5px;

		&:hover {
			background-color: #fff2;
			border-radius: 5px;
		}
	}

	button {
		min-width: 100px;
		text-align: left;
		padding: 5px;
		margin: 2px 5px;
	}
</style>
