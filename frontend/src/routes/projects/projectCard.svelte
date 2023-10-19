<script lang="ts">
    import { PUBLIC_API_URL } from "$env/static/public";
    import type { ProjectCardDTO } from "$lib/models/project";
    import { onMount } from "svelte";

    export let project:ProjectCardDTO

    let dateTime:string
    let imgURL:string

    onMount(() => {
        let date:Date = new Date(project.Timestamp * 1000)
        dateTime = date.toLocaleDateString("nl-NL")

        imgURL = project.BannerURL ? `${PUBLIC_API_URL}/img/${project.BannerURL}` : "404-error.png"
    })
</script>

<a href="//{project.ProjectUrl}" class="card p-4 mx-auto my-5 min-w-[80%]">
    <img src={imgURL} class="mx-auto max-h-96" alt={project.BannerAlt}/>
    <h2 class="h2 mb-5">{project.Title}</h2>
    <p>{project.Description}</p>
    <p class="text-end">{dateTime}</p>
</a>