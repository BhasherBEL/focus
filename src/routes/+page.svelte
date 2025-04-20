<script lang="ts">
	import { SvelteToast } from '@zerodevx/svelte-toast';

	const { data } = $props();
	const projects = data.projects;
</script>

<section class="relative mx-2 my-10 max-w-3xl md:mx-auto">
	<h2 class="mb-8 text-center text-3xl font-bold">Projects</h2>
	<table class="w-full border-separate border-spacing-y-2">
		<thead>
			<tr>
				<th class="w-10 px-4 py-2 text-center font-semibold text-gray-400">#</th>
				<th class="px-4 py-2 text-left font-semibold text-gray-400">Name</th>
				<th class="w-32 px-4 py-2 text-center font-semibold text-gray-400"># Cards</th>
				<th class="w-24 px-4 py-2"></th>
			</tr>
		</thead>
		<tbody>
			{#each projects as project, i (project.id)}
				<tr class="rounded-lg bg-gray-800">
					<td class="rounded-l-lg px-4 py-4 text-center font-mono text-gray-300">{i + 1}</td>
					<td class="px-4 py-4 font-bold">
						<a
							href={`/project/?id=${project.id}`}
							class="hover:underline focus:outline-none"
							tabindex="0"
						>
							{project.title}
						</a>
					</td>
					<td class="px-4 py-4 text-center">14</td>
					<td class="rounded-r-lg px-2 py-2">
						<div class="flex items-center gap-2">
							<a
								href={`/edit/${project.id}`}
								class="rounded p-2 hover:bg-gray-700 focus:outline-none"
								aria-label="Edit"
								tabindex="0"
							>
								<svg
									xmlns="http://www.w3.org/2000/svg"
									class="h-5 w-5 fill-current"
									viewBox="0 0 20 20"
								>
									<path
										d="M17.414 2.586a2 2 0 0 0-2.828 0l-9.5 9.5A2 2 0 0 0 4 13.914V16a1 1 0 0 0 1 1h2.086a2 2 0 0 0 1.414-.586l9.5-9.5a2 2 0 0 0 0-2.828l-1.586-1.586zM6.5 15H5v-1.5l8.793-8.793 1.5 1.5L6.5 15zm9.207-9.207-1.5-1.5 1.207-1.207a1 1 0 0 1 1.414 1.414L15.707 5.793z"
									/>
								</svg>
							</a>
							<form
								method="POST"
								action="?/delete"
								class="m-0 p-0"
								onsubmit={(e) => {
									e.preventDefault();
									if (window.confirm('Are you sure you want to delete this project?')) {
										(e.target as HTMLFormElement).submit();
									}
								}}
							>
								<input type="hidden" name="id" value={project.id} />
								<button
									type="submit"
									class="cursor-pointer rounded p-2 hover:bg-gray-700 focus:outline-none"
									aria-label="Delete"
									tabindex="0"
								>
									<svg
										xmlns="http://www.w3.org/2000/svg"
										class="h-5 w-5 fill-current"
										viewBox="0 0 20 20"
									>
										<path
											d="M6 8a1 1 0 0 1 2 0v6a1 1 0 1 1-2 0V8zm4 0a1 1 0 0 1 2 0v6a1 1 0 1 1-2 0V8zm5-3h-3.5l-1-1h-3l-1 1H3a1 1 0 1 0 0 2h1v10a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V7h1a1 1 0 1 0 0-2zM5 7v10h10V7H5z"
										/>
									</svg>
								</button>
							</form>
						</div>
					</td>
				</tr>
			{/each}
		</tbody>
	</table>
	<a
		href="/create"
		class="btn btn-primary float-end items-center justify-center rounded text-2xl"
		aria-label="Create new project"
	>
		+
	</a>
</section>

<SvelteToast />

<style lang="less"></style>
