<script lang="ts">
    import { fetchProjects } from "$lib/fetchProjects";
    import type { ProjectCardDTO } from "$lib/models/project";
    import { onMount } from "svelte";
    import ProjectCard from "./projectCard.svelte";
    import Error from "../components/error.svelte";
    
    let data : ProjectCardDTO[];
    onMount(async () => {
        data = await fetchProjects()
    })
</script>
{#if data}
<div class="grid md:grid-cols-2">
    {#each data as project, _}
    <ProjectCard project={project}/>
    {/each}
</div>
{:else}
<Error/>
{/if}
