<script lang="ts">
	import currentView from '$lib/stores/currentView';
	import Card, { cards } from '$lib/types/Card';
	import type Filter from '$lib/types/Filter';
	import type Project from '$lib/types/Project';
	import type ProjectTag from '$lib/types/ProjectTag';
	import { projectTags } from '$lib/types/ProjectTag';
	import type TagOption from '$lib/types/TagOption';
	import type View from '$lib/types/View';
	import Column from './Column.svelte';
	import Header from './Header.svelte';

	export let project: Project;

	function cardComparator(a: Card, b: Card, sortTag: ProjectTag | null, sortDirection: number) {
		if (!sortTag) return 0;

		const aTag = a.cardTags.find((t) => t.projectTag === sortTag);
		const bTag = b.cardTags.find((t) => t.projectTag === sortTag);

		if (!aTag) return -sortDirection;
		if (!bTag) return sortDirection;

		const aValue = aTag.value || aTag.option?.value || '';
		const bValue = bTag.value || bTag.option?.value || '';

		return aValue < bValue ? sortDirection : -sortDirection;
	}

	function passFilters(card: Card, filters: Filter[]): boolean {
		for (const projectTag of $projectTags) {
			let is: TagOption[] = [];

			const cardTag = card.cardTags.find((t) => t.projectTag === projectTag);

			for (const filter of filters) {
				if (projectTag !== filter.projectTag) continue;
				if (!filter.tagOption) continue;

				if (filter.filterType === 0) {
					is.push(filter.tagOption);
				} else if (filter.filterType === 1) {
					if (cardTag?.option === filter.tagOption) return false;
				}
			}

			if (is.length > 0) {
				if (!cardTag) return false;
				if (!is.some((o) => o === cardTag?.option)) return false;
			}
		}
		return true;
	}

	function columnPassFilters(tagOption: TagOption | null, filters: Filter[]): boolean {
		let is: TagOption[] = [];

		for (const filter of filters) {
			if ($currentView?.sortTag !== filter.projectTag) continue;
			if (!filter.tagOption) continue;

			if (filter.filterType === 0) {
				is.push(filter.tagOption);
			} else if (filter.filterType === 1) {
				if (tagOption === filter.tagOption) return false;
			}
		}

		if (is.length > 0) {
			if (!is.some((o) => o === tagOption)) return false;
		}

		return true;
	}

	function extractColumnCards(view: View | null, cards: Card[], tagOption: TagOption | null) {
		if (!view) return cards;

		const filteredCards = cards.filter((c) => passFilters(c, view.filters));

		const primaryTag = view.primaryTag;

		if (!primaryTag) return filteredCards;

		if (!tagOption) {
			return filteredCards.filter((c) => !c.cardTags.map((t) => t.projectTag).includes(primaryTag));
		}

		const rightOptionCards = filteredCards.filter((c) => {
			const tag = c.cardTags.find((t) => t.projectTag === primaryTag);
			if (!tag) return false;
			return tag.option?.id === tagOption.id;
		});

		const sortedCards = rightOptionCards.sort((a, b) =>
			cardComparator(a, b, view.sortTag, $currentView?.sortDirection || 1)
		);

		return sortedCards;
	}

	function isFiltered(tagOption: TagOption) {
		if (!$currentView) return false;
		if (!$currentView.filters) return false;

		for (const filter of $currentView.filters) {
			if (filter.tagOption === tagOption) return true;
		}

		return false;
	}
</script>

{#if project}
	<section>
		{#if $currentView}
			<Header {project} />
			{#if $cards}
				<div class="grid">
					{#if $currentView.primaryTag}
						{#each $currentView.primaryTag.options as option (option.id)}
							{#if columnPassFilters(option, $currentView.filters)}
								<Column
									{option}
									primaryTag={$currentView.primaryTag}
									columnCards={extractColumnCards($currentView, $cards, option)}
									{project}
								/>
							{/if}
						{/each}
					{/if}
					{#if columnPassFilters(null, $currentView.filters)}
						<Column
							primaryTag={$currentView.primaryTag}
							columnCards={extractColumnCards($currentView, $cards, null)}
							{project}
						/>
					{/if}
				</div>
			{/if}
		{/if}
	</section>
{/if}

<style lang="less">
	@import '../../styles/breakpoints.less';

	section {
		display: flex;
		flex-direction: column;
		height: 100vh;
		flex-basis: 100%;
		overflow: hidden;
	}

	.grid {
		display: flex;
		flex-direction: row;
		flex: 1;
		overflow: scroll;
		padding: 0 2%;
		scroll-snap-type: x mandatory;

		.one-column({
			padding: 0 10%;
		});
	}
</style>
