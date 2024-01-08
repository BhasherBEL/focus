<script lang="ts">
	import currentView from '$lib/stores/currentView';
	import { cards } from '$lib/types/Card';
	import type Project from '$lib/types/Project';
	import Column from './Column.svelte';
	import Header from './Header.svelte';

	export let project: Project;
</script>

{#if project}
	<section>
		{#if $currentView}
			<Header {project} />
			{#if $cards}
				<div class="grid">
					{#if $currentView.primaryTag}
						{#each $currentView.primaryTag.options as option}
							<Column
								{option}
								primaryTag={$currentView.primaryTag}
								title={option.value}
								columnCards={$cards
									.filter((c) => c.cardTags.map((t) => t.option).includes(option))
									.sort((a, b) => {
										if (!$currentView?.sortTag) return 0;
										const aTag = a.cardTags.find((t) => t.projectTag === $currentView?.sortTag);
										const bTag = b.cardTags.find((t) => t.projectTag === $currentView?.sortTag);

										if (!aTag) return -($currentView?.sortDirection || 1);
										if (!bTag) return $currentView?.sortDirection || 1;

										const aValue = aTag.value || aTag.option?.value || '';
										const bValue = bTag.value || bTag.option?.value || '';

										return aValue < bValue
											? $currentView?.sortDirection || 1
											: -($currentView?.sortDirection || 1);
									})}
								{project}
							/>
						{/each}
					{/if}
					<Column
						primaryTag={$currentView.primaryTag}
						title={$currentView.primaryTag ? `No ${$currentView.title}` : 'No groups'}
						columnCards={$currentView.primaryTag != null
							? (() => {
									const primaryTag = $currentView.primaryTag;
									return $cards.filter(
										(c) => !c.cardTags.map((t) => t.projectTag).includes(primaryTag)
									);
								})()
							: $cards}
						{project}
					/>
				</div>
			{/if}
		{/if}
	</section>
{/if}

<style>
	section {
		display: flex;
		flex-direction: column;
		height: 100vh;
		transition: all 0.3s ease-in-out;
		width: 100vw;

		@media (min-width: 800px) {
			margin-left: 250px;
			width: calc(100vw - 250px);
		}
	}

	.grid {
		display: flex;
		flex-direction: row;
		flex: 1;
		overflow: scroll;
		padding: 0 40px;
	}
</style>
