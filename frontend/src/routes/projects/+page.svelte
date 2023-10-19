<script lang="ts">
    import { fetchProjects } from "$lib/fetchProjects";
    import type { Project } from "$lib/models/project";
    import { onMount } from "svelte";
    import ProjectCard from "./projectCard.svelte";
    
    let data : Project[];
    onMount(async () => {
        data = await fetchProjects()
    })
</script>

<div class="grid md:grid-cols-2">
    {#if data}
        {#each data as project, _}
            <ProjectCard project={project}/>
            <ProjectCard project={project}/>
            <ProjectCard project={project}/>
            <ProjectCard project={project}/>
            <ProjectCard project={project}/>
            <ProjectCard project={project}/>
            <ProjectCard project={project}/>
            <ProjectCard project={project}/>
        {/each}
    {/if}

</div>