<script lang="ts">
	import { onMount } from 'svelte';
	import Column from './column.svelte';
	import type { Project, View } from '../../stores/interfaces';
	import projectTags from '../../stores/projectTags';
	import { cards, currentView, views } from '../../stores/smallStore';
	import Header from './header.svelte';

	export let project: Project;

	let view: View | null = null;

	onMount(async () => {
		await cards.init(project.id);

		if (!(await projectTags.init(project.id))) {
			return;
		}

		currentView.subscribe((v) => {
			view = v;
		});
	});
</script>

{#if project}
	<section>
		{#if view}
			<Header {project} {view} />
			{#if cards}
				<div class="grid">
					{#if view.primary_tag_id}
						{#each $projectTags[view.primary_tag_id].options as option}
							<Column
								optionId={option.id}
								primary_tag_id={view.primary_tag_id}
								title={option.value}
								columnCards={$cards.filter((c) =>
									c.tags.map((t) => t.option_id).includes(option.id)
								)}
								projectId={project.id}
							/>
						{/each}
					{/if}
					<Column
						primary_tag_id={view.primary_tag_id}
						title={view.primary_tag_id
							? `No ${$projectTags[view.primary_tag_id].title}`
							: 'No groups'}
						columnCards={view.primary_tag_id
							? $cards.filter(
									(c) => !c.tags.map((t) => t.tag_id).includes(view?.primary_tag_id || -2)
								)
							: $cards}
						projectId={project.id}
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
