<script lang="ts">
    export const projectId: number = 0;
</script>

{#if Number.isNaN(projectId)}
    <p>Invalid number</p>
{:else}
    <p>Hello on dashboard {projectId}!</p>
{/if}